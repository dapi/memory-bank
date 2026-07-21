package doctor

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/dapi/memory-bank/tools/internal/lint"
	"github.com/dapi/memory-bank/tools/internal/ownership"
	"gopkg.in/yaml.v3"
)

// doctorCommandPattern accepts a memory-bank invocation at the start of a shell
// command, including after common shell command separators. It intentionally
// does not match arbitrary text in a run block (for example, echo commands).
var doctorCommandPattern = regexp.MustCompile(`(?m)(?:^|[;&|]\s*|\bthen\s+|\bdo\s+)(?:[A-Za-z_][A-Za-z0-9_]*=[^\s]+\s+)*(?:\S*/)?memory-bank["']?\s+doctor(?:\s|$)`)

func NormalizeProfile(value string) (Profile, error) {
	profile := Profile(strings.ToLower(strings.TrimSpace(value)))
	switch profile {
	case ProfileAuto, ProfileTemplate, ProfileDownstream:
		return profile, nil
	default:
		return "", fmt.Errorf("--profile must be auto, template, or downstream")
	}
}

// Run performs only reads. In particular it never stages files or invokes Git.
func Run(options Options) (Report, error) {
	repoRoot, err := filepath.Abs(options.RepoRoot)
	if err != nil {
		return Report{}, err
	}
	scopeRoot, err := lint.NormalizeScopeRoot(options.ScopeRoot)
	if err != nil {
		return Report{}, err
	}
	profile := options.Profile
	if profile == ProfileAuto {
		profile = detectProfile(repoRoot)
	}
	navigation, err := lint.Run(lint.Options{RepoRoot: repoRoot, ScopeRoot: scopeRoot, MaxDepth: options.MaxDepth})
	if err != nil {
		return Report{}, err
	}
	report := Report{FormatVersion: ReportFormatVersion, Profile: profile, RepoRoot: repoRoot, Navigation: navigation, Findings: []Finding{}}
	report.addNavigationFindings()
	report.checkIdentityAndDrift(options.AgentFile)
	report.checkGovernance(scopeRoot)
	report.checkCI()
	sort.SliceStable(report.Findings, func(i, j int) bool {
		left, right := report.Findings[i], report.Findings[j]
		if left.Group != right.Group {
			return left.Group < right.Group
		}
		if left.Code != right.Code {
			return left.Code < right.Code
		}
		if left.Path != right.Path {
			return left.Path < right.Path
		}
		return left.Subject < right.Subject
	})
	for _, finding := range report.Findings {
		switch finding.Severity {
		case Error:
			report.Summary.Errors++
		case Warning:
			report.Summary.Warnings++
		case Info:
			report.Summary.Info++
		}
	}
	return report, nil
}

func detectProfile(repoRoot string) Profile {
	if _, err := os.Lstat(filepath.Join(repoRoot, ownership.LockFileName)); err == nil {
		return ProfileDownstream
	}
	module, err := os.ReadFile(filepath.Join(repoRoot, "tools", "go.mod"))
	if err == nil && strings.Contains(string(module), "module github.com/dapi/memory-bank/tools") {
		return ProfileTemplate
	}
	return ProfileDownstream
}

func (report *Report) add(finding Finding) { report.Findings = append(report.Findings, finding) }

func (report *Report) checkIdentityAndDrift(agentFile string) {
	lock, exists, lockErr := ownership.ReadLock(report.RepoRoot)
	if lockErr != nil {
		report.add(Finding{Code: "manifest.invalid", Severity: Error, Group: "manifest", Path: ownership.LockFileName, Message: lockErr.Error(), Remediation: "Repair or recreate the ownership lock with memory-bank init from a trusted template checkout."})
	} else if exists {
		report.TemplateIdentity = TemplateIdentity{SchemaVersion: lock.SchemaVersion, Version: lock.Template.Version, SourceRef: lock.Template.SourceRef}
		report.add(Finding{Code: "template.identity", Severity: Info, Group: "template_identity", Path: ownership.LockFileName, Subject: lock.Template.Version, Message: "Installed template identity is recorded in the ownership lock.", Remediation: "Use memory-bank update --dry-run to compare with a newer pinned template source."})
		report.checkManagedDrift(lock)
	} else if report.Profile == ProfileDownstream {
		report.add(Finding{Code: "template.identity_missing", Severity: Error, Group: "template_identity", Path: ownership.LockFileName, Message: "The downstream repository has no ownership lock, so its installed template version is unknown.", Remediation: "Adopt the template with memory-bank init and commit memory-bank/.lock."})
	} else {
		report.add(Finding{Code: "template.source_repository", Severity: Info, Group: "template_identity", Subject: "template", Message: "Template source profile detected; an installed-template lock is not expected.", Remediation: "Create locks only in downstream repositories through memory-bank init."})
	}

	contents, _, err := readRegularWithinRoot(report.RepoRoot, agentFile)
	if err != nil {
		severity := Error
		message := "Agent instruction entrypoint is missing or unreadable."
		if !os.IsNotExist(err) {
			message = fmt.Sprintf("Cannot read agent instruction entrypoint: %v", err)
		}
		report.add(Finding{Code: "agent.entrypoint_missing", Severity: severity, Group: "agent_integration", Path: agentFile, Message: message, Remediation: "Create the agent instruction file and link it to memory-bank/README.md."})
	} else if !strings.Contains(string(contents), "memory-bank/README.md") {
		report.add(Finding{Code: "agent.memory_bank_link_missing", Severity: Error, Group: "agent_integration", Path: agentFile, Message: "Agent instructions do not route readers to memory-bank/README.md.", Remediation: "Add a repository-relative link to memory-bank/README.md or run memory-bank update with the same --agent-file."})
	}
	if exists && lockErr == nil {
		agentReport, err := ownership.InspectAgentInstructions(report.RepoRoot, agentFile)
		if err != nil {
			report.add(Finding{Code: "agent.managed_block_invalid", Severity: Error, Group: "agent_integration", Path: agentFile, Message: err.Error(), Remediation: "Resolve unsafe paths or damaged markers, then run memory-bank update."})
		} else if agentReport.DriftCount > 0 {
			decision := agentReport.Decisions[0]
			report.add(Finding{Code: "agent.managed_block_drift", Severity: Error, Group: "agent_integration", Path: decision.Path, Subject: string(decision.Action), Message: decision.Reason, Remediation: "Run memory-bank update using the pinned template checkout and the same --agent-file."})
		}
	}
}

func (report *Report) checkManagedDrift(lock ownership.Lock) {
	paths := make([]string, 0, len(lock.Files))
	for filePath := range lock.Files {
		paths = append(paths, filePath)
	}
	sort.Strings(paths)
	for _, filePath := range paths {
		contract := lock.Files[filePath]
		if contract.Ownership != ownership.Managed && contract.Ownership != ownership.Generated {
			continue
		}
		data, info, err := readRegularWithinRoot(report.RepoRoot, filePath)
		if err != nil {
			code, message := "manifest.managed_unreadable", fmt.Sprintf("Managed file cannot be inspected: %v", err)
			if os.IsNotExist(err) {
				code, message = "manifest.managed_missing", "Managed file recorded by the lock is missing."
			}
			report.add(Finding{Code: code, Severity: Error, Group: "manifest", Path: filePath, Message: message, Remediation: "Restore the file from the pinned template with memory-bank update."})
			continue
		}
		digest := fmt.Sprintf("sha256:%x", sha256.Sum256(data))
		if digest != contract.PayloadDigest {
			report.add(Finding{Code: "manifest.managed_content_drift", Severity: Error, Group: "manifest", Path: filePath, Message: "Managed file content differs from the lock payload digest.", Remediation: "Review the local change, then restore/update it through memory-bank update."})
		}
		mode := "100644"
		if info.Mode().Perm()&0o111 != 0 {
			mode = "100755"
		}
		if contract.PayloadMode != "" && mode != contract.PayloadMode {
			report.add(Finding{Code: "manifest.managed_mode_drift", Severity: Error, Group: "manifest", Path: filePath, Message: "Managed file executable mode differs from the lock.", Remediation: "Restore the mode recorded by the ownership lock."})
		}
	}
}

func readRegularWithinRoot(repoRoot, relativePath string) ([]byte, fs.FileInfo, error) {
	current := repoRoot
	for _, component := range strings.Split(filepath.FromSlash(relativePath), string(filepath.Separator)) {
		if component == "" || component == "." || component == ".." {
			return nil, nil, fmt.Errorf("unsafe repository-relative path %q", relativePath)
		}
		current = filepath.Join(current, component)
		info, err := os.Lstat(current)
		if err != nil {
			return nil, nil, err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil, nil, fmt.Errorf("unsafe symlink in path %q", relativePath)
		}
	}
	info, err := os.Stat(current)
	if err != nil {
		return nil, nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, info, fmt.Errorf("path %q is not a regular file", relativePath)
	}
	data, err := os.ReadFile(current)
	return data, info, err
}

func (report *Report) checkCI() {
	workflowRoot := filepath.Join(report.RepoRoot, ".github", "workflows")
	doctorGatePresent := false
	cliInstalledAtLatest := false
	_ = filepath.WalkDir(workflowRoot, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		if ext := strings.ToLower(filepath.Ext(path)); ext != ".yml" && ext != ".yaml" {
			return nil
		}
		data, readErr := os.ReadFile(path)
		if readErr == nil {
			doctorGatePresent = doctorGatePresent || workflowRunsDoctor(data)
			cliInstalledAtLatest = cliInstalledAtLatest || strings.Contains(string(data), "cmd/memory-bank@latest")
		}
		return nil
	})
	if !doctorGatePresent {
		report.add(Finding{Code: "ci.doctor_gate_missing", Severity: Warning, Group: "downstream_ci", Path: ".github/workflows", Message: "No GitHub Actions workflow runs memory-bank doctor.", Remediation: "Add a read-only memory-bank doctor gate; downstream CI should install a pinned CLI version."})
	}
	if report.Profile == ProfileDownstream && cliInstalledAtLatest {
		report.add(Finding{Code: "ci.cli_version_unpinned", Severity: Warning, Group: "downstream_ci", Path: ".github/workflows", Message: "CI installs memory-bank from the moving @latest reference.", Remediation: "Pin the install to a release tag or full commit SHA."})
	}
}

// workflowRunsDoctor reports whether a valid GitHub Actions workflow has a run
// field that executes memory-bank doctor. Parsing YAML ensures comments and
// unrelated workflow metadata cannot satisfy the CI gate diagnostic.
func workflowRunsDoctor(data []byte) bool {
	var document yaml.Node
	if err := yaml.Unmarshal(data, &document); err != nil {
		return false
	}
	return yamlNodeRunsDoctor(&document)
}

func yamlNodeRunsDoctor(node *yaml.Node) bool {
	if node.Kind == yaml.MappingNode {
		for index := 0; index+1 < len(node.Content); index += 2 {
			key, value := node.Content[index], node.Content[index+1]
			if key.Value == "run" && value.Kind == yaml.ScalarNode && doctorCommandPattern.MatchString(value.Value) {
				return true
			}
			if yamlNodeRunsDoctor(value) {
				return true
			}
		}
		return false
	}
	for _, child := range node.Content {
		if yamlNodeRunsDoctor(child) {
			return true
		}
	}
	return false
}
