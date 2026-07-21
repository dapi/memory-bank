package ownership

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/dapi/memory-bank/tools/internal/agentinstructions"
)

type payload struct {
	data   []byte
	digest string
	mode   string
}

type mutation struct {
	decision         Decision
	data             []byte
	mode             fs.FileMode
	modeSet          bool
	expectedExists   bool
	expectedDigest   string
	expectedMode     string
	preconditions    []destinationPrecondition
	topology         *topologySnapshot
	topologyReplaced bool
}

type destinationPrecondition struct {
	path   string
	digest string
	mode   string
}

var immutableRefPattern = regexp.MustCompile(`^[0-9a-fA-F]{40}([0-9a-fA-F]{24})?$`)

// Init creates a lock and safely installs missing template files.
func Init(options Options) (Report, error) {
	repo, err := pinRepoRoot(options.RepoRoot)
	if err != nil {
		return Report{}, err
	}
	options.RepoRoot = repo.root
	if _, exists, _, err := readLockSnapshot(repo); err != nil {
		return Report{}, err
	} else if exists {
		return Report{}, fmt.Errorf("%s already exists; use memory-bank update", LockFileName)
	}
	return run(options, Lock{}, false, repo, "")
}

// Update applies a source template against an existing lock.
func Update(options Options) (Report, error) {
	repo, err := pinRepoRoot(options.RepoRoot)
	if err != nil {
		return Report{}, err
	}
	options.RepoRoot = repo.root
	lock, exists, lockDigest, err := readLockSnapshot(repo)
	if err != nil {
		return Report{}, err
	}
	if !exists {
		return Report{}, ErrLockNotFound
	}
	return run(options, lock, true, repo, lockDigest)
}

func run(options Options, old Lock, hasLock bool, repo pinnedRepo, lockDigest string) (Report, error) {
	if options.RepoRoot == "" || options.SourceRoot == "" || options.TemplateVersion == "" || options.SourceRef == "" {
		return Report{}, errors.New("repo root, source root, template version, and immutable source ref are required")
	}
	if !immutableRefPattern.MatchString(options.SourceRef) {
		return Report{}, errors.New("source ref must be a full 40- or 64-character hexadecimal commit ID")
	}
	pinnedSource, err := pinSourceRoot(options.SourceRoot)
	if err != nil {
		return Report{}, err
	}
	if err := rejectOverlappingRoots(repo, pinnedSource); err != nil {
		return Report{}, err
	}
	verifySource := verifySourceCheckout
	if options.verifySource != nil {
		verifySource = options.verifySource
	}
	if err := verifySource(pinnedSource.root, options.SourceRef); err != nil {
		return Report{}, err
	}
	options.SourceRoot = pinnedSource.root
	var source map[string]payload
	if options.verifySource == nil {
		source, err = readGitSource(pinnedSource, options.SourceRef)
	} else {
		source, err = readSource(pinnedSource)
	}
	if err != nil {
		return Report{}, err
	}
	if err := verifySource(pinnedSource.root, options.SourceRef); err != nil {
		return Report{}, fmt.Errorf("source checkout changed while reading template: %w", err)
	}
	mutations, decisions, next, err := buildPlan(repo, source, old, hasLock)
	if err != nil {
		return Report{}, err
	}
	templateMutationCount := len(mutations)
	agentMutation, agentDecision, err := buildAgentPlan(repo, options.AgentFile)
	if err != nil {
		return Report{}, err
	}
	decisions = append(decisions, agentDecision)
	if agentMutation != nil {
		mutations = append(mutations, *agentMutation)
	}
	report := Report{FormatVersion: ReportFormatVersion, DryRun: options.DryRun, Decisions: decisions}
	for _, decision := range decisions {
		if decision.Action == Conflict {
			report.ConflictCount++
		}
	}
	if report.ConflictCount > 0 || options.DryRun {
		return report, nil
	}
	template := Template{Version: options.TemplateVersion, SourceRef: options.SourceRef}
	needsLockWrite := !hasLock || templateMutationCount > 0 || old.SchemaVersion != CurrentSchemaVersion || old.Template != template
	if !needsLockWrite {
		if len(mutations) == 0 {
			return report, nil
		}
		if err := applyAtomicallyPinned(options, mutations, repo); err != nil {
			var committed *committedError
			if errors.As(err, &committed) {
				report.Applied = true
				return report, err
			}
			return Report{}, err
		}
		report.Applied = true
		return report, nil
	}
	now := time.Now
	if options.Now != nil {
		now = options.Now
	}
	next.SchemaVersion = CurrentSchemaVersion
	next.Template = template
	next.LastUpdate = UpdateRecord{Version: options.TemplateVersion, At: now().UTC()}
	lockData, err := marshalLock(next)
	if err != nil {
		return Report{}, err
	}
	mutations = append(mutations, mutation{
		decision:       Decision{Path: LockFileName, Action: UpdateFile, Reason: "record successful update"},
		data:           lockData,
		expectedExists: hasLock,
		expectedDigest: lockDigest,
		preconditions:  lockPreconditions(next),
	})
	if err := applyAtomicallyPinned(options, mutations, repo); err != nil {
		var committed *committedError
		if errors.As(err, &committed) {
			report.Applied = true
			return report, err
		}
		return Report{}, err
	}
	report.Applied = true
	return report, nil
}

func buildAgentPlan(repo pinnedRepo, target string) (*mutation, Decision, error) {
	return buildAgentPlanWithReader(repo, target, secureReadDestination)
}

func buildAgentPlanWithReader(repo pinnedRepo, target string, readDestination func(pinnedRepo, string) (os.FileInfo, []byte, error)) (*mutation, Decision, error) {
	if target == "" {
		target = agentinstructions.DefaultTarget
	}
	firstComponent := strings.SplitN(target, "/", 2)[0]
	if strings.EqualFold(firstComponent, "memory-bank") {
		return nil, Decision{}, fmt.Errorf("agent instruction file must be outside memory-bank/: %q", target)
	}
	_, info, exists, err := inspectDestination(repo, target)
	if err != nil {
		return nil, Decision{}, err
	}
	var original []byte
	currentDigest, currentMode := "", ""
	if exists {
		var readInfo os.FileInfo
		readInfo, original, err = readDestination(repo, target)
		if err != nil {
			return nil, Decision{}, fmt.Errorf("read agent instruction file %q: %w", target, err)
		}
		if !os.SameFile(info, readInfo) {
			return nil, Decision{}, fmt.Errorf("read agent instruction file %q: destination changed while it was being planned", target)
		}
		currentDigest = digest(original)
		currentMode = observedMode(info.Mode().Perm())
	}
	plan := agentinstructions.BuildPlan(original)
	decision := Decision{Path: target, Ownership: Managed, Diff: plan.Diff}
	switch plan.Status {
	case agentinstructions.Current:
		decision.Action, decision.Reason = Preserve, "managed Memory Bank block is current"
		return nil, decision, nil
	case agentinstructions.Missing:
		if exists {
			decision.Action, decision.Reason = UpdateFile, "add missing managed Memory Bank block"
		} else {
			decision.Action, decision.Reason = Create, "create agent instructions with managed Memory Bank block"
		}
	case agentinstructions.Outdated:
		decision.Action, decision.Reason = UpdateFile, "replace outdated managed Memory Bank block"
	case agentinstructions.Ambiguous:
		decision.Action, decision.Reason = Conflict, "managed Memory Bank markers are damaged or ambiguous"
		return nil, decision, nil
	}
	mode := fs.FileMode(0o644)
	if exists {
		mode = info.Mode().Perm()
	}
	return &mutation{decision: decision, data: plan.Data, mode: mode, modeSet: true, expectedExists: exists, expectedDigest: currentDigest, expectedMode: currentMode}, decision, nil
}

// Doctor checks only the managed agent-instruction contract without mutation.
func Doctor(repoRoot, target string) (Report, error) {
	repo, err := pinRepoRoot(repoRoot)
	if err != nil {
		return Report{}, err
	}
	_, decision, err := buildAgentPlan(repo, target)
	if err != nil {
		return Report{}, err
	}
	report := Report{FormatVersion: ReportFormatVersion, DryRun: true, Decisions: []Decision{decision}}
	if decision.Action == Conflict {
		report.ConflictCount = 1
	}
	if decision.Action != Preserve {
		report.DriftCount = 1
	}
	return report, nil
}

func readSource(source pinnedSource) (map[string]payload, error) {
	if err := inspectSourceRoot(source); err != nil {
		return nil, err
	}
	root := source.root
	memoryBankRoot := filepath.Join(root, "memory-bank")
	result := make(map[string]payload)
	err := filepath.WalkDir(memoryBankRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("template source contains unsupported symlink: %s", path)
		}
		if entry.IsDir() {
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("template source contains unsupported file: %s", path)
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		result[filepath.ToSlash(relative)] = payload{data: data, digest: digest(data), mode: gitMode(info.Mode().Perm())}
		return nil
	})
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("template source has no memory-bank directory: %s", root)
	}
	if err != nil {
		return nil, err
	}
	if err := inspectSourceRoot(source); err != nil {
		return nil, err
	}
	return result, nil
}

func readGitSource(source pinnedSource, ref string) (map[string]payload, error) {
	if err := inspectSourceRoot(source); err != nil {
		return nil, err
	}
	tree, err := gitBytes(source.root, "ls-tree", "-rz", "--full-tree", ref, "--", "memory-bank")
	if err != nil {
		return nil, fmt.Errorf("read pinned source tree: %w", err)
	}
	result := make(map[string]payload)
	for _, record := range strings.Split(string(tree), "\x00") {
		if record == "" {
			continue
		}
		header, filePath, found := strings.Cut(record, "\t")
		fields := strings.Fields(header)
		if !found || len(fields) != 3 || fields[1] != "blob" || fields[0] != "100644" && fields[0] != "100755" {
			return nil, fmt.Errorf("read pinned source tree: unsupported entry %q", filePath)
		}
		data, err := gitBytes(source.root, "cat-file", "blob", fields[2])
		if err != nil {
			return nil, fmt.Errorf("read pinned source file %q: %w", filePath, err)
		}
		result[filePath] = payload{data: data, digest: digest(data), mode: fields[0]}
	}
	if len(result) == 0 {
		return nil, errors.New("template source has no memory-bank payload")
	}
	if err := inspectSourceRoot(source); err != nil {
		return nil, err
	}
	return result, nil
}

func gitMode(mode fs.FileMode) string {
	if mode&0o111 != 0 {
		return "100755"
	}
	return "100644"
}

func fileMode(mode string) fs.FileMode {
	if mode == "100755" {
		return 0o755
	}
	return 0o644
}

func observedMode(mode fs.FileMode) string {
	if runtime.GOOS == "windows" {
		return ""
	}
	return gitMode(mode)
}

func modeMatches(observed, expected string) bool {
	return observed == "" || observed == expected
}

func buildPlan(repo pinnedRepo, source map[string]payload, old Lock, hasLock bool) ([]mutation, []Decision, Lock, error) {
	if _, err := inspectRepoRoot(repo.root, repo.info); err != nil {
		return nil, nil, Lock{}, err
	}
	if _, exists := source[LockFileName]; exists {
		return nil, nil, Lock{}, fmt.Errorf("template source contains reserved metadata path: %s", LockFileName)
	}
	next := Lock{Files: make(map[string]File)}
	var removalMutations []mutation
	var removalDecisions []Decision
	cleanRemovals := make(map[string]string)
	removalMutationIndex := make(map[string]int)
	removed := make([]string, 0)
	for path := range old.Files {
		if _, exists := source[path]; !exists {
			removed = append(removed, path)
		}
	}
	sort.Strings(removed)
	for _, path := range removed {
		prior := old.Files[path]
		currentDigest, exists, err := digestDestinationFile(repo, path)
		if err != nil {
			return nil, nil, Lock{}, err
		}
		decision := Decision{Path: path, Ownership: prior.Ownership}
		currentMode := ""
		if exists {
			_, info, _, inspectErr := inspectDestination(repo, path)
			if inspectErr != nil {
				return nil, nil, Lock{}, inspectErr
			}
			currentMode = observedMode(info.Mode().Perm())
		}
		switch {
		case !exists:
			decision.Action, decision.Reason = Preserve, "file already absent"
		case prior.Ownership == Managed && currentDigest == prior.PayloadDigest && destinationModeMatches(repo, path, prior.PayloadMode):
			decision.Action, decision.Reason = Delete, "unmodified managed file was removed upstream"
			cleanRemovals[path] = currentDigest
			removalMutationIndex[path] = len(removalMutations)
			removalMutations = append(removalMutations, mutation{decision: decision, expectedExists: true, expectedDigest: currentDigest, expectedMode: currentMode})
		case prior.Ownership == Generated:
			decision.Action, decision.Reason = Delete, "generated file was removed upstream"
			cleanRemovals[path] = currentDigest
			removalMutationIndex[path] = len(removalMutations)
			removalMutations = append(removalMutations, mutation{decision: decision, expectedExists: true, expectedDigest: currentDigest, expectedMode: currentMode})
		case prior.Ownership == Managed:
			decision.Action, decision.Reason = Conflict, "removed managed file has downstream drift"
			next.Files[path] = prior
		default:
			decision.Action, decision.Reason = Preserve, "downstream-owned file is never deleted"
			next.Files[path] = prior
		}
		removalDecisions = append(removalDecisions, decision)
	}
	sort.SliceStable(removalMutations, func(i, j int) bool {
		return strings.Count(removalMutations[i].decision.Path, "/") > strings.Count(removalMutations[j].decision.Path, "/")
	})
	for index := range removalMutations {
		removalMutationIndex[removalMutations[index].decision.Path] = index
	}

	var sourceMutations []mutation
	var sourceDecisions []Decision
	paths := make([]string, 0, len(source))
	for path := range source {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		incoming := source[path]
		class := Classify(path)
		prior, tracked := old.Files[path]
		currentDigest, exists, topology, err := inspectDestinationForPlan(repo, path, cleanRemovals)
		if err != nil {
			return nil, nil, Lock{}, err
		}
		currentMode := ""
		if exists {
			_, info, _, inspectErr := inspectDestination(repo, path)
			if inspectErr != nil {
				return nil, nil, Lock{}, inspectErr
			}
			currentMode = observedMode(info.Mode().Perm())
		}
		priorBaseMode, priorPayloadMode := prior.BaseMode, prior.PayloadMode
		if priorBaseMode == "" {
			priorBaseMode = currentMode
			if priorBaseMode == "" {
				priorBaseMode = incoming.mode
			}
		}
		if priorPayloadMode == "" {
			priorPayloadMode = currentMode
			if priorPayloadMode == "" {
				priorPayloadMode = incoming.mode
			}
		}
		decision := Decision{Path: path, Ownership: class}
		file := File{Ownership: class, BaseDigest: incoming.digest, BaseMode: incoming.mode}
		if class == Managed || class == Generated {
			file.PayloadDigest = incoming.digest
			file.PayloadMode = incoming.mode
		}
		switch {
		case !exists && !tracked:
			decision.Action, decision.Reason = Create, "template file is missing"
		case !exists && prior.Ownership == Generated:
			decision.Action, decision.Reason = UpdateFile, "regenerate deterministic payload"
		case !exists && prior.Ownership == Managed:
			decision.Action, decision.Reason = Conflict, "managed file has downstream drift"
			file = prior
		case !exists:
			decision.Action, decision.Reason = Conflict, "downstream-owned file was deleted; explicit resolution required"
			file = prior
		case !hasLock && class == Generated:
			if currentDigest == incoming.digest && modeMatches(currentMode, incoming.mode) {
				decision.Action, decision.Reason = Preserve, "adopt existing generated payload"
			} else {
				decision.Action, decision.Reason = UpdateFile, "regenerate deterministic payload"
			}
		case !hasLock:
			if class == Managed && (currentDigest != incoming.digest || !modeMatches(currentMode, incoming.mode)) {
				decision.Action, decision.Reason = Conflict, "existing managed file does not match initialization source"
			} else {
				decision.Action, decision.Reason = Preserve, "adopt existing file without overwriting it"
			}
		case !tracked:
			decision.Action, decision.Reason = Preserve, "untracked existing file is downstream-owned"
			file = File{Ownership: UserOwned}
		case class == UserOwned || prior.Ownership == UserOwned:
			decision.Action, decision.Reason = Preserve, "user-owned files are never overwritten"
			file = prior
		case class == Generated:
			if currentDigest == incoming.digest && modeMatches(currentMode, incoming.mode) {
				decision.Action, decision.Reason = Preserve, "generated payload is current"
			} else {
				decision.Action, decision.Reason = UpdateFile, "regenerate deterministic payload"
			}
		case class == Managed:
			if currentDigest == incoming.digest && modeMatches(currentMode, incoming.mode) {
				decision.Action, decision.Reason = Preserve, "managed payload matches incoming template"
			} else if currentDigest != prior.PayloadDigest || !modeMatches(currentMode, priorPayloadMode) {
				decision.Action, decision.Reason = Conflict, "managed file has downstream drift"
			} else if incoming.digest != prior.BaseDigest || incoming.mode != priorBaseMode {
				decision.Action, decision.Reason = UpdateFile, "managed template payload changed"
			} else {
				decision.Action, decision.Reason = Preserve, "managed payload is current"
			}
		case class == Adapted:
			downstreamChanged := currentDigest != prior.BaseDigest || !modeMatches(currentMode, priorBaseMode)
			upstreamChanged := incoming.digest != prior.BaseDigest || incoming.mode != priorBaseMode
			if currentDigest == incoming.digest && modeMatches(currentMode, incoming.mode) {
				decision.Action, decision.Reason = Preserve, "adapted payload matches incoming template base"
			} else if downstreamChanged && upstreamChanged {
				decision.Action, decision.Reason = Conflict, "adapted file changed both upstream and downstream"
				file = prior
			} else if upstreamChanged {
				decision.Action, decision.Reason = UpdateFile, "unmodified adapted file follows new template base"
			} else {
				decision.Action, decision.Reason = Preserve, "preserve downstream adaptation"
			}
		}
		decision.Ownership = file.Ownership
		if decision.Action == Create || decision.Action == UpdateFile {
			sourceMutations = append(sourceMutations, mutation{
				decision: decision, data: incoming.data, mode: fileMode(incoming.mode), modeSet: true, expectedExists: exists, expectedDigest: currentDigest, expectedMode: currentMode, topology: topology,
			})
			if topology != nil {
				for _, prerequisite := range topology.files {
					if index, ok := removalMutationIndex[prerequisite.path]; ok {
						removalMutations[index].topologyReplaced = true
					}
				}
			}
		}
		next.Files[path] = file
		sourceDecisions = append(sourceDecisions, decision)
	}
	mutations := append(removalMutations, sourceMutations...)
	decisions := append(sourceDecisions, removalDecisions...)
	return mutations, decisions, next, nil
}

func destinationModeMatches(repo pinnedRepo, relative, expected string) bool {
	if expected == "" {
		return true
	}
	_, info, exists, err := inspectDestination(repo, relative)
	return err == nil && exists && modeMatches(observedMode(info.Mode().Perm()), expected)
}

func lockPreconditions(lock Lock) []destinationPrecondition {
	paths := make([]string, 0, len(lock.Files))
	for path, file := range lock.Files {
		if file.Ownership == Managed || file.Ownership == Generated {
			paths = append(paths, path)
		}
	}
	sort.Strings(paths)
	result := make([]destinationPrecondition, 0, len(paths))
	for _, path := range paths {
		result = append(result, destinationPrecondition{path: path, digest: lock.Files[path].PayloadDigest, mode: lock.Files[path].PayloadMode})
	}
	return result
}

type transactionOps struct {
	writeFile             func(string, []byte, fs.FileMode) error
	rename                func(string, string) error
	link                  func(string, string) error
	removeAll             func(string) error
	renameToDestination   func(pinnedRepo, string, string) error
	renameFromDestination func(pinnedRepo, string, string) error
	linkToDestination     func(pinnedRepo, string, string) error
}

type committedError struct {
	err error
}

func (err *committedError) Error() string {
	return "update committed but " + err.err.Error()
}

func (err *committedError) Unwrap() error {
	return err.err
}

var osTransactionOps = transactionOps{
	writeFile:             os.WriteFile,
	rename:                os.Rename,
	link:                  os.Link,
	removeAll:             os.RemoveAll,
	renameToDestination:   secureRenameToDestination,
	renameFromDestination: secureRenameFromDestination,
	linkToDestination:     secureLinkToDestination,
}

type stagedMutation struct {
	mutation
	target               string
	replacement          string
	backup               string
	applied              string
	originalInfo         fs.FileInfo
	originalDigest       string
	replacementInfo      fs.FileInfo
	replacementDigest    string
	originalMoved        bool
	replacementInstalled bool
}

func applyAtomically(options Options, mutations []mutation) error {
	repo, err := pinRepoRoot(options.RepoRoot)
	if err != nil {
		return err
	}
	return applyAtomicallyPinnedWithOps(options, mutations, repo, osTransactionOps)
}

func applyAtomicallyPinned(options Options, mutations []mutation, repo pinnedRepo) error {
	return applyAtomicallyPinnedWithOps(options, mutations, repo, osTransactionOps)
}

// applyAtomicallyWithOps prepares every payload before it mutates a target.
// Existing targets are moved, rather than copied, into same-filesystem staging
// so rollback does not need to allocate space to restore them.
func applyAtomicallyWithOps(options Options, mutations []mutation, ops transactionOps) error {
	repo, err := pinRepoRoot(options.RepoRoot)
	if err != nil {
		return err
	}
	return applyAtomicallyPinnedWithOps(options, mutations, repo, ops)
}

func applyAtomicallyPinnedWithOps(options Options, mutations []mutation, repo pinnedRepo, ops transactionOps) (resultErr error) {
	if _, err := inspectRepoRoot(repo.root, repo.info); err != nil {
		return err
	}
	if ops.writeFile == nil {
		ops.writeFile = os.WriteFile
	}
	if ops.rename == nil {
		ops.rename = os.Rename
	}
	if ops.link == nil {
		ops.link = os.Link
	}
	repoRoot := repo.root
	options.RepoRoot = repoRoot
	staging, err := os.MkdirTemp(repoRoot, ".memory-bank-update-")
	if err != nil {
		return fmt.Errorf("create update staging: %w", err)
	}
	cleanupStaging := true
	commitComplete := false
	defer func() {
		if !cleanupStaging {
			return
		}
		if _, err := inspectRepoRoot(repo.root, repo.info); err != nil {
			cleanupErr := fmt.Errorf("update staging retained at %s because repo root changed: %w", staging, err)
			if commitComplete {
				resultErr = errors.Join(resultErr, &committedError{err: cleanupErr})
			} else {
				resultErr = errors.Join(resultErr, cleanupErr)
			}
			return
		}
		removeAll := ops.removeAll
		if removeAll == nil {
			removeAll = os.RemoveAll
		}
		if err := removeAll(staging); err != nil {
			cleanupErr := fmt.Errorf("update staging retained at %s: %w", staging, err)
			if commitComplete {
				resultErr = errors.Join(resultErr, &committedError{err: cleanupErr})
			} else {
				resultErr = errors.Join(resultErr, cleanupErr)
			}
		}
	}()

	newDirectory := filepath.Join(staging, "new")
	oldDirectory := filepath.Join(staging, "old")
	appliedDirectory := filepath.Join(staging, "applied")
	if err := os.Mkdir(newDirectory, 0o700); err != nil {
		return fmt.Errorf("prepare update staging: %w", err)
	}
	if err := os.Mkdir(oldDirectory, 0o700); err != nil {
		return fmt.Errorf("prepare update staging: %w", err)
	}
	if err := os.Mkdir(appliedDirectory, 0o700); err != nil {
		return fmt.Errorf("prepare update staging: %w", err)
	}
	linkProbeSource := filepath.Join(staging, "link-probe-source")
	linkProbeTarget := filepath.Join(staging, "link-probe-target")
	if err := os.WriteFile(linkProbeSource, nil, 0o600); err != nil {
		return fmt.Errorf("prepare rollback-link probe: %w", err)
	}
	if err := os.Link(linkProbeSource, linkProbeTarget); err != nil {
		return fmt.Errorf("filesystem does not support safe no-clobber updates: %w", err)
	}
	if err := os.Remove(linkProbeSource); err != nil {
		return fmt.Errorf("remove rollback-link probe source: %w", err)
	}
	if err := os.Remove(linkProbeTarget); err != nil {
		return fmt.Errorf("remove rollback-link probe target: %w", err)
	}

	staged := make([]stagedMutation, len(mutations))
	for index, item := range mutations {
		var target string
		var info fs.FileInfo
		var exists bool
		var err error
		if item.topology != nil {
			target, _, err = destinationPathLexicalPinned(repo, item.decision.Path)
			if err == nil {
				err = verifyTopologySnapshot(repo, item.decision.Path, item.topology)
			}
		} else {
			target, info, exists, err = inspectDestination(repo, item.decision.Path)
		}
		if err != nil {
			return fmt.Errorf("prepare %s: %w", item.decision.Path, err)
		}
		if exists != item.expectedExists {
			return fmt.Errorf("prepare %s: destination changed while update was being planned", item.decision.Path)
		}
		currentDigest := ""
		if exists {
			var digestExists bool
			currentDigest, digestExists, err = digestDestinationFile(repo, item.decision.Path)
			if err != nil {
				return fmt.Errorf("prepare %s: %w", item.decision.Path, err)
			}
			if !digestExists || (item.expectedDigest != "" && currentDigest != item.expectedDigest) {
				return fmt.Errorf("prepare %s: destination content changed while update was being planned", item.decision.Path)
			}
			if item.expectedMode != "" && !modeMatches(observedMode(info.Mode().Perm()), item.expectedMode) {
				return fmt.Errorf("prepare %s: destination mode changed while update was being planned", item.decision.Path)
			}
		}
		staged[index] = stagedMutation{
			mutation:       item,
			target:         target,
			backup:         filepath.Join(oldDirectory, fmt.Sprintf("%06d", index)),
			applied:        filepath.Join(appliedDirectory, fmt.Sprintf("%06d", index)),
			originalInfo:   info,
			originalDigest: currentDigest,
		}
		if item.decision.Action == Delete {
			continue
		}
		mode := item.mode
		if mode == 0 && !item.modeSet {
			mode = 0o644
		}
		staged[index].replacement = filepath.Join(newDirectory, fmt.Sprintf("%06d", index))
		if err := ops.writeFile(staged[index].replacement, item.data, mode); err != nil {
			return fmt.Errorf("stage %s: %w", item.decision.Path, err)
		}
		if err := os.Chmod(staged[index].replacement, mode); err != nil {
			return fmt.Errorf("stage %s mode: %w", item.decision.Path, err)
		}
		replacementInfo, replacementDigest, err := inspectRegularFile(staged[index].replacement)
		if err != nil {
			return fmt.Errorf("stage %s: %w", item.decision.Path, err)
		}
		if replacementDigest != digest(item.data) {
			return fmt.Errorf("stage %s: staged payload content mismatch", item.decision.Path)
		}
		staged[index].replacementInfo = replacementInfo
		staged[index].replacementDigest = replacementDigest
	}

	createdDirectories := make([]string, 0)
	removedDirectories := make([]removedDirectory, 0)
	rollback := func() error {
		var rollbackErrors []error
		if _, err := inspectRepoRoot(repo.root, repo.info); err != nil {
			return fmt.Errorf("restore pinned repo root: %w", err)
		}
		for index := len(staged) - 1; index >= 0; index-- {
			item := &staged[index]
			if item.replacementInstalled {
				if err := moveInstalledToRecovery(repo, item, ops); err != nil {
					rollbackErrors = append(rollbackErrors, fmt.Errorf("recover current %s: %w", item.decision.Path, err))
				}
			}
		}
		for index := len(createdDirectories) - 1; index >= 0; index-- {
			relative := createdDirectories[index]
			if err := secureRemoveDestination(repo, relative, true); err != nil && !errors.Is(err, os.ErrNotExist) {
				rollbackErrors = append(rollbackErrors, fmt.Errorf("remove new directory %s: %w", relative, err))
			}
		}
		for index := len(removedDirectories) - 1; index >= 0; index-- {
			directory := removedDirectories[index]
			if err := secureMkdirDestination(repo, directory.path, directory.mode); err != nil && !errors.Is(err, os.ErrExist) {
				rollbackErrors = append(rollbackErrors, fmt.Errorf("recreate replaced directory %s: %w", directory.path, err))
			}
		}
		for index := len(staged) - 1; index >= 0; index-- {
			item := &staged[index]
			if item.originalMoved {
				if err := restoreOriginalFromBackup(repo, item, ops); err != nil {
					rollbackErrors = append(rollbackErrors, fmt.Errorf("restore %s: %w", item.decision.Path, err))
				}
			}
		}
		return errors.Join(rollbackErrors...)
	}

	fail := func(cause error) error {
		rollbackErr := rollback()
		if rollbackErr == nil {
			cleanupStaging = true
			return cause
		}
		return errors.Join(
			cause,
			fmt.Errorf("rollback incomplete; recovery data retained at %s: %w", staging, rollbackErr),
		)
	}

	// From the first target mutation onward, staging may contain the only copy
	// of an original. Unexpected unwinding must retain it for recovery.
	cleanupStaging = false
	for index := range staged {
		item := &staged[index]
		if options.BeforeMutation != nil {
			if err := options.BeforeMutation(item.decision); err != nil {
				return fail(fmt.Errorf("update interrupted before %s: %w", item.decision.Path, err))
			}
		}
		if _, err := inspectRepoRoot(repo.root, repo.info); err != nil {
			return fail(fmt.Errorf("apply %s: %w", item.decision.Path, err))
		}
		if item.decision.Path == LockFileName {
			for priorIndex := 0; priorIndex < index; priorIndex++ {
				prior := &staged[priorIndex]
				if !prior.originalMoved && !prior.replacementInstalled {
					continue
				}
				if prior.originalMoved {
					if err := verifyOriginalBackup(prior); err != nil {
						return fail(fmt.Errorf("apply %s: backup for %s changed: %w", item.decision.Path, prior.decision.Path, err))
					}
				}
				if !prior.topologyReplaced {
					if err := verifyRollbackTarget(repo, prior); err != nil {
						return fail(fmt.Errorf("apply %s: previously mutated %s changed: %w", item.decision.Path, prior.decision.Path, err))
					}
				}
			}
		}
		for _, precondition := range item.preconditions {
			if err := verifyDestinationPrecondition(repo, precondition); err != nil {
				return fail(fmt.Errorf("apply %s: verify %s: %w", item.decision.Path, precondition.path, err))
			}
		}
		if item.topology == nil {
			if err := verifyOriginalTarget(repo, item); err != nil {
				return fail(fmt.Errorf("apply %s: %w", item.decision.Path, err))
			}
		}
		if item.expectedExists {
			if ops.renameFromDestination != nil && sameOperation(ops.rename, os.Rename) {
				if err := ops.renameFromDestination(repo, item.decision.Path, item.backup); err != nil {
					return fail(fmt.Errorf("apply %s: move original to staging: %w", item.decision.Path, err))
				}
			} else if err := ops.rename(item.target, item.backup); err != nil {
				return fail(fmt.Errorf("apply %s: move original to staging: %w", item.decision.Path, err))
			}
			item.originalMoved = true
			backupInfo, backupDigest, err := inspectRegularFile(item.backup)
			if err != nil {
				return fail(fmt.Errorf("apply %s: inspect staged original: %w", item.decision.Path, err))
			}
			if !os.SameFile(item.originalInfo, backupInfo) {
				return fail(fmt.Errorf("apply %s: original identity changed while moving to staging", item.decision.Path))
			}
			if backupDigest != item.originalDigest {
				item.originalDigest = backupDigest
				return fail(fmt.Errorf("apply %s: destination content changed while moving to staging", item.decision.Path))
			}
		}
		if item.decision.Action == Delete {
			continue
		}
		if item.topology != nil {
			if err := prepareTopologyDestination(repo, item.decision.Path, item.topology, &removedDirectories); err != nil {
				return fail(fmt.Errorf("apply %s: %w", item.decision.Path, err))
			}
		}
		if err := ensureDestinationParents(repo, item.decision.Path, &createdDirectories); err != nil {
			return fail(fmt.Errorf("apply %s: %w", item.decision.Path, err))
		}
		if _, _, nowExists, err := inspectDestination(repo, item.decision.Path); err != nil {
			return fail(fmt.Errorf("apply %s: %w", item.decision.Path, err))
		} else if nowExists {
			return fail(fmt.Errorf("apply %s: destination appeared during update", item.decision.Path))
		}
		if err := verifyStagedReplacement(item); err != nil {
			return fail(fmt.Errorf("apply %s: %w", item.decision.Path, err))
		}
		if ops.linkToDestination != nil && sameOperation(ops.link, os.Link) {
			if err := ops.linkToDestination(repo, item.decision.Path, item.replacement); err != nil {
				return fail(fmt.Errorf("apply %s: install staged payload without replacing destination: %w", item.decision.Path, err))
			}
		} else if err := ops.link(item.replacement, item.target); err != nil {
			return fail(fmt.Errorf("apply %s: install staged payload without replacing destination: %w", item.decision.Path, err))
		}
		item.replacementInstalled = true
		if err := os.Remove(item.replacement); err != nil {
			return fail(fmt.Errorf("apply %s: detach installed payload from staging: %w", item.decision.Path, err))
		}
	}
	commitComplete = true
	cleanupStaging = true
	return nil
}

func inspectDestination(repo pinnedRepo, relative string) (string, fs.FileInfo, bool, error) {
	target, err := destinationPathPinned(repo, relative)
	if err != nil {
		return "", nil, false, err
	}
	info, err := os.Lstat(target)
	if errors.Is(err, os.ErrNotExist) {
		return target, nil, false, nil
	}
	if err != nil {
		return "", nil, false, fmt.Errorf("inspect destination path %q: %w", relative, err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return "", nil, false, fmt.Errorf("unsafe destination path %q: file is a symlink", relative)
	}
	if !info.Mode().IsRegular() {
		return "", nil, false, fmt.Errorf("unsupported destination file %q", relative)
	}
	return target, info, true, nil
}

func ensureDestinationParents(repo pinnedRepo, relative string, created *[]string) error {
	return secureEnsureDestinationParents(repo, relative, created)
}

func verifyOriginalTarget(repo pinnedRepo, item *stagedMutation) error {
	_, currentInfo, exists, err := inspectDestination(repo, item.decision.Path)
	if err != nil {
		return err
	}
	if exists != item.expectedExists {
		return errors.New("destination existence changed after staging")
	}
	if !exists {
		return nil
	}
	if !os.SameFile(item.originalInfo, currentInfo) {
		return errors.New("destination identity changed after staging")
	}
	if item.expectedMode != "" && !modeMatches(observedMode(currentInfo.Mode().Perm()), item.expectedMode) {
		return errors.New("destination mode changed after staging")
	}
	currentDigest, digestExists, err := digestDestinationFile(repo, item.decision.Path)
	if err != nil {
		return err
	}
	if !digestExists || currentDigest != item.originalDigest {
		return errors.New("destination content changed after staging")
	}
	return nil
}

func verifyDestinationPrecondition(repo pinnedRepo, precondition destinationPrecondition) error {
	_, info, exists, err := inspectDestination(repo, precondition.path)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("managed payload is missing")
	}
	readInfo, data, err := secureReadDestination(repo, precondition.path)
	if err != nil {
		return err
	}
	if !os.SameFile(info, readInfo) {
		return errors.New("managed payload identity changed before lock commit")
	}
	currentDigest := digest(data)
	if currentDigest != precondition.digest {
		return errors.New("managed payload content changed before lock commit")
	}
	if precondition.mode != "" && !modeMatches(observedMode(info.Mode().Perm()), precondition.mode) {
		return errors.New("managed payload mode changed before lock commit")
	}
	return nil
}

func verifyStagedReplacement(item *stagedMutation) error {
	info, payloadDigest, err := inspectRegularFile(item.replacement)
	if err != nil {
		return fmt.Errorf("inspect staged payload: %w", err)
	}
	if !os.SameFile(item.replacementInfo, info) || payloadDigest != item.replacementDigest {
		return errors.New("staged payload changed before installation")
	}
	if !modeMatches(observedMode(info.Mode().Perm()), gitMode(item.mode)) {
		return errors.New("staged payload mode changed before installation")
	}
	return nil
}

func verifyInstalledTarget(repo pinnedRepo, item *stagedMutation) error {
	_, currentInfo, exists, err := inspectDestination(repo, item.decision.Path)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("installed payload is missing")
	}
	if !os.SameFile(item.replacementInfo, currentInfo) {
		return errors.New("installed payload identity changed")
	}
	if !modeMatches(observedMode(currentInfo.Mode().Perm()), gitMode(item.mode)) {
		return errors.New("installed payload mode changed")
	}
	currentDigest, digestExists, err := digestDestinationFile(repo, item.decision.Path)
	if err != nil {
		return err
	}
	if !digestExists || currentDigest != item.replacementDigest {
		return errors.New("installed payload content changed")
	}
	return nil
}

func verifyRollbackTarget(repo pinnedRepo, item *stagedMutation) error {
	if item.replacementInstalled {
		return verifyInstalledTarget(repo, item)
	}
	_, _, exists, err := inspectDestination(repo, item.decision.Path)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("destination appeared after original was staged")
	}
	return nil
}

func moveInstalledToRecovery(repo pinnedRepo, item *stagedMutation, ops transactionOps) error {
	_, _, exists, err := inspectDestination(repo, item.decision.Path)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("installed payload is missing")
	}
	var moveErr error
	if ops.renameFromDestination != nil && sameOperation(ops.rename, os.Rename) {
		moveErr = ops.renameFromDestination(repo, item.decision.Path, item.applied)
	} else {
		moveErr = ops.rename(item.target, item.applied)
	}
	if moveErr != nil {
		return fmt.Errorf("move installed payload to recovery: %w", moveErr)
	}
	movedInfo, movedDigest, err := inspectRegularFile(item.applied)
	if err != nil {
		return fmt.Errorf("inspect recovered payload: %w", err)
	}
	if os.SameFile(item.replacementInfo, movedInfo) && movedDigest == item.replacementDigest {
		return nil
	}
	if err := linkToDestination(repo, item, item.applied, ops); err != nil {
		return errors.Join(
			errors.New("installed payload changed before rollback"),
			fmt.Errorf("preserve changed payload: %w", err),
		)
	}
	return errors.New("installed payload changed before rollback; current content was preserved")
}

func restoreOriginalFromBackup(repo pinnedRepo, item *stagedMutation, ops transactionOps) error {
	backupInfo, backupDigest, err := inspectRegularFile(item.backup)
	if err != nil {
		return fmt.Errorf("inspect backup: %w", err)
	}
	if !os.SameFile(item.originalInfo, backupInfo) || backupDigest != item.originalDigest {
		return errors.New("backup changed during update")
	}
	_, _, exists, err := inspectDestination(repo, item.decision.Path)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("destination appeared before original could be restored")
	}
	if err := linkToDestination(repo, item, item.backup, ops); err != nil {
		return fmt.Errorf("restore original without replacing destination: %w", err)
	}
	currentInfo, data, err := secureReadDestination(repo, item.decision.Path)
	if err != nil {
		return fmt.Errorf("verify restored original: %w", err)
	}
	currentDigest := digest(data)
	if !os.SameFile(item.originalInfo, currentInfo) || currentDigest != item.originalDigest {
		return errors.New("restored original failed integrity verification")
	}
	return nil
}

func linkToDestination(repo pinnedRepo, item *stagedMutation, source string, ops transactionOps) error {
	if ops.linkToDestination != nil && sameOperation(ops.link, os.Link) {
		return ops.linkToDestination(repo, item.decision.Path, source)
	}
	return ops.link(source, item.target)
}

func sameOperation(left, right any) bool {
	return reflect.ValueOf(left).Pointer() == reflect.ValueOf(right).Pointer()
}

func verifyOriginalBackup(item *stagedMutation) error {
	backupInfo, backupDigest, err := inspectRegularFile(item.backup)
	if err != nil {
		return fmt.Errorf("inspect backup: %w", err)
	}
	if !os.SameFile(item.originalInfo, backupInfo) {
		return errors.New("backup identity changed during update")
	}
	if backupDigest != item.originalDigest {
		item.originalDigest = backupDigest
		return errors.New("backup content changed during update")
	}
	return nil
}

func inspectRegularFile(path string) (fs.FileInfo, string, error) {
	info, data, err := readRegularFile(path)
	if err != nil {
		return nil, "", err
	}
	return info, digest(data), nil
}

func readRegularFile(path string) (fs.FileInfo, []byte, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return nil, nil, fmt.Errorf("unsupported file: %s", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	after, err := os.Lstat(path)
	if err != nil {
		return nil, nil, err
	}
	if after.Mode()&os.ModeSymlink != 0 || !after.Mode().IsRegular() || !os.SameFile(info, after) {
		return nil, nil, fmt.Errorf("file changed while reading: %s", path)
	}
	return after, data, nil
}

func digestDestinationFile(repo pinnedRepo, relative string) (string, bool, error) {
	_, _, exists, err := inspectDestination(repo, relative)
	if err != nil {
		return "", false, err
	}
	if !exists {
		return "", false, nil
	}
	_, data, err := secureReadDestination(repo, relative)
	if err != nil {
		return "", false, err
	}
	return digest(data), true, nil
}

func digest(data []byte) string {
	sum := sha256.Sum256(data)
	return "sha256:" + strings.ToLower(hex.EncodeToString(sum[:]))
}
