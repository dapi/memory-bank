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
var doctorCommandPattern = regexp.MustCompile(`(?m)(?:^|[;&|]\s*|\bthen\s+|\bdo\s+)(?:[A-Za-z_][A-Za-z0-9_]*=[^\s]+\s+)*(?:\S*/)?memory-bank["']?\s+doctor(?:[ \t]|$|[;&|])`)

var shellErrexitStatePattern = regexp.MustCompile(`(?m)(?:^|[;\n])\s*set\s+([+-])(?:e|o\s+(?:errexit|e))(?:\s|;|$)`)
var shellSuccessfulExitPattern = regexp.MustCompile(`(?m)(?:^|[;\n])\s*exit\s+0(?:\s*#.*)?\s*$`)
var shellGuaranteedFailurePattern = regexp.MustCompile(`^(?:false|exit\s+[1-9][0-9]*)(?:\s*#.*)?$`)
var shellStatusPropagationPattern = regexp.MustCompile(`^exit\s+\$\?(?:\s*#.*)?$`)
var shellErrexitPattern = regexp.MustCompile(`(?:^|\s)-[A-Za-z]*e[A-Za-z]*(?:\s|$)|(?:^|\s)-o\s+errexit(?:\s|$)`)
var shellPipefailPattern = regexp.MustCompile(`(?:^|\s)-o\s+pipefail(?:\s|$)|(?:^|\s)pipefail(?:\s|$)`)
var workflowFalseExpressionPattern = regexp.MustCompile(`^\$\{\{\s*false\s*\}\}$`)

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
	if len(document.Content) != 1 {
		return false
	}
	jobs := yamlMappingValue(document.Content[0], "jobs")
	if jobs == nil || jobs.Kind != yaml.MappingNode {
		return false
	}
	for index := 1; index < len(jobs.Content); index += 2 {
		job := jobs.Content[index]
		if workflowConditionStaticallyFalse(yamlMappingValue(job, "if")) || workflowJobAllowsFailure(job) {
			continue
		}
		steps := yamlMappingValue(job, "steps")
		if steps == nil || steps.Kind != yaml.SequenceNode {
			continue
		}
		for _, step := range steps.Content {
			if workflowConditionStaticallyFalse(yamlMappingValue(step, "if")) || workflowStepAllowsFailure(step) {
				continue
			}
			run := yamlMappingValue(step, "run")
			shell := effectiveWorkflowShell(document.Content[0], job, step)
			if run != nil && run.Kind == yaml.ScalarNode && runHasBlockingDoctor(run.Value, workflowShellHasErrexit(shell), workflowShellHasPipefail(shell, job)) {
				return true
			}
		}
	}
	return false
}

func effectiveWorkflowShell(workflow, job, step *yaml.Node) *yaml.Node {
	if shell := yamlMappingValue(step, "shell"); shell != nil {
		return shell
	}
	if defaults := yamlMappingValue(job, "defaults"); defaults != nil {
		if runDefaults := yamlMappingValue(defaults, "run"); runDefaults != nil {
			if shell := yamlMappingValue(runDefaults, "shell"); shell != nil {
				return shell
			}
		}
	}
	if defaults := yamlMappingValue(workflow, "defaults"); defaults != nil {
		if runDefaults := yamlMappingValue(defaults, "run"); runDefaults != nil {
			if shell := yamlMappingValue(runDefaults, "shell"); shell != nil {
				return shell
			}
		}
	}
	return nil
}

func workflowJobAllowsFailure(job *yaml.Node) bool {
	return workflowContinueOnErrorAllowsFailure(yamlMappingValue(job, "continue-on-error"))
}

func workflowStepAllowsFailure(step *yaml.Node) bool {
	return workflowContinueOnErrorAllowsFailure(yamlMappingValue(step, "continue-on-error"))
}

func workflowConditionStaticallyFalse(condition *yaml.Node) bool {
	if condition == nil || condition.Kind != yaml.ScalarNode {
		return false
	}
	if condition.Tag == "!!bool" {
		var value bool
		return condition.Decode(&value) == nil && !value
	}
	value := strings.TrimSpace(condition.Value)
	return value == "false" || workflowFalseExpressionPattern.MatchString(value)
}

func workflowContinueOnErrorAllowsFailure(policy *yaml.Node) bool {
	if policy == nil {
		return false
	}
	// GitHub parses literal boolean expressions as strings in YAML, but the
	// expression is still statically guaranteed to disable continue-on-error.
	if policy.Kind == yaml.ScalarNode && workflowFalseExpressionPattern.MatchString(strings.TrimSpace(policy.Value)) {
		return false
	}
	// Only an explicit YAML boolean false guarantees that doctor can fail the
	// job. Expressions and other scalar forms may evaluate to a non-blocking
	// policy at runtime.
	if policy.Kind != yaml.ScalarNode || policy.Tag != "!!bool" {
		return true
	}
	var allowsFailure bool
	if err := policy.Decode(&allowsFailure); err != nil {
		return true
	}
	return allowsFailure
}

func runHasBlockingDoctor(run string, shellHasErrexit, shellHasPipefail bool) bool {
	for _, match := range doctorCommandPattern.FindAllStringIndex(run, -1) {
		if !doctorFailureSuppressionAt(run, match[1], shellHasErrexit, shellHasPipefail) {
			return true
		}
	}
	return false
}

func doctorFailureSuppressionAt(run string, start int, shellHasErrexit, shellHasPipefail bool) bool {
	if start > 0 && strings.ContainsRune(";&|", rune(run[start-1])) {
		start--
	}
	operator := shellOperatorIndex(run, start)
	if operator < 0 {
		return false
	}
	if run[operator] == ';' || run[operator] == '\n' {
		if !shellErrexitEnabledAt(run[:start], shellHasErrexit) {
			remainder := strings.TrimSpace(run[operator+1:])
			if shellImmediatelyExitsSuccessfully(run, operator) || shellSuccessfulExitPattern.MatchString(remainder) || shellStatusPropagationPattern.MatchString(remainder) {
				return true
			}
			// Without errexit, any subsequent command can successfully replace
			// doctor's exit status (for example, `doctor; cleanup`).
			return shellFollowingCommandSuppresses(remainder)
		}
		return false
	}

	// A doctor command can be followed by several AND/OR operators.  In
	// `doctor && upload || true`, the final true masks doctor failures; looking
	// only at the first && incorrectly treats this as a blocking gate.
	listEnd := shellListEnd(run, operator)
	if run[operator] == '|' && (operator+1 >= len(run) || run[operator+1] != '|') && !shellHasPipefail {
		// A custom shell without pipefail can still make the pipeline blocking
		// by checking the doctor's PIPESTATUS entry afterward.
		return !pipelineStatusCheckBlocks(run, operator, pipelineDoctorIndex(run, start))
	}
	lastOperator := operator
	for next := shellOperatorIndex(run, operator+shellOperatorLength(run, operator)); next >= 0 && next < listEnd; next = shellOperatorIndex(run, next+shellOperatorLength(run, next)) {
		if run[next] == ';' {
			break
		}
		lastOperator = next
	}
	if !strings.HasPrefix(run[lastOperator:], "||") {
		return !shellErrexitEnabledAt(run[:start], shellHasErrexit) && shellFollowingCommandSuppresses(strings.TrimSpace(run[listEnd+1:]))
	}
	// Any recovery command after `||` can turn a failed doctor invocation into
	// a successful step; restricting this to a small allowlist misses commands
	// such as `notify-team` or `pwd`.
	recovery := strings.TrimSpace(run[lastOperator+2 : listEnd])
	if recovery != "" && !strings.HasPrefix(recovery, "#") && !shellGuaranteedFailurePattern.MatchString(recovery) {
		return true
	}
	return !shellErrexitEnabledAt(run[:start], shellHasErrexit) && shellFollowingCommandSuppresses(strings.TrimSpace(run[listEnd+1:]))
}

func shellFollowingCommandSuppresses(remainder string) bool {
	if remainder == "" || strings.HasPrefix(remainder, "#") {
		return false
	}
	return !shellStatusPropagationPattern.MatchString(remainder) && !shellGuaranteedFailurePattern.MatchString(remainder)
}

// pipelineStatusCheckBlocks recognizes the common explicit Bash check used
// after piping doctor's output through tee, for example:
// `memory-bank doctor | tee report.txt; test ${PIPESTATUS[0]} -eq 0`.
func pipelineStatusCheckBlocks(run string, pipelineOperator, doctorIndex int) bool {
	if pipelineOperator < 0 || pipelineOperator >= len(run) {
		return false
	}
	pattern := regexp.MustCompile(fmt.Sprintf(`(?m)\b(?:test|\[)\s+\$\{?PIPESTATUS\[%d\]\}?\s+(?:-eq|==|=)\s+0\b`, doctorIndex))
	return pattern.MatchString(run[pipelineOperator+1:])
}

func pipelineDoctorIndex(run string, doctorStart int) int {
	prefix := run[:doctorStart]
	boundary := -1
	for _, marker := range []string{";", "\n", "&&", "||"} {
		if index := strings.LastIndex(prefix, marker); index+len(marker)-1 > boundary {
			boundary = index + len(marker) - 1
		}
	}
	prefix = prefix[boundary+1:]
	index := 0
	for position := 0; position < len(prefix); position++ {
		if prefix[position] == '|' && (position+1 >= len(prefix) || prefix[position+1] != '|') && (position == 0 || prefix[position-1] != '|') {
			index++
		}
	}
	return index
}

func shellErrexitEnabledAt(run string, defaultEnabled bool) bool {
	enabled := defaultEnabled
	for _, match := range shellErrexitStatePattern.FindAllStringSubmatch(run, -1) {
		enabled = match[1] == "-"
	}
	return enabled
}

// workflowShellHasErrexit reports whether the shell used by a step is known
// to stop on a failed command. GitHub adds -e only for its default shell;
// custom shells must opt into errexit themselves.
func workflowShellHasErrexit(shell *yaml.Node) bool {
	if shell == nil {
		return true
	}
	if shell.Kind != yaml.ScalarNode {
		return false
	}
	value := shell.Value
	if value == "bash" || value == "sh" {
		return true
	}
	return shellErrexitPattern.MatchString(value)
}

// workflowShellHasPipefail reports whether the shell propagates a failed
// command from inside a pipeline. An omitted shell uses GitHub's implicit
// `bash -e {0}` on Linux and PowerShell on Windows. The latter propagates a
// native command's LASTEXITCODE through the runner wrapper.
func workflowShellHasPipefail(shell, job *yaml.Node) bool {
	if shell == nil {
		return workflowJobRunsOnWindows(job)
	}
	if shell.Kind != yaml.ScalarNode {
		return false
	}
	if shell.Value == "bash" {
		return true
	}
	return shellPipefailPattern.MatchString(shell.Value)
}

func workflowJobRunsOnWindows(job *yaml.Node) bool {
	runsOn := yamlMappingValue(job, "runs-on")
	if runsOn == nil {
		return false
	}
	if runsOn.Kind == yaml.MappingNode {
		return workflowRunsOnLabelsIncludeWindows(yamlMappingValue(runsOn, "labels"))
	}
	return workflowRunsOnLabelsIncludeWindows(runsOn)
}

func workflowRunsOnLabelsIncludeWindows(runsOn *yaml.Node) bool {
	if runsOn == nil {
		return false
	}
	if runsOn.Kind == yaml.ScalarNode {
		return isWindowsRunnerLabel(runsOn.Value)
	}
	if runsOn.Kind == yaml.SequenceNode {
		for _, label := range runsOn.Content {
			if label.Kind == yaml.ScalarNode && isWindowsRunnerLabel(label.Value) {
				return true
			}
		}
	}
	return false
}

func isWindowsRunnerLabel(label string) bool {
	value := strings.ToLower(strings.TrimSpace(label))
	return value == "windows" || strings.HasPrefix(value, "windows-")
}

// shellImmediatelyExitsSuccessfully recognizes an exit 0 command that
// directly follows the doctor command. Delayed terminal exits are handled by
// shellSuccessfulExitPattern in doctorFailureSuppressionAt.
func shellImmediatelyExitsSuccessfully(run string, separator int) bool {
	position := separator + 1
	for position < len(run) && (run[position] == ' ' || run[position] == '\t' || run[position] == '\r' || run[position] == '\n') {
		position++
	}
	if !strings.HasPrefix(run[position:], "exit") {
		return false
	}
	position += len("exit")
	if position >= len(run) || (run[position] != ' ' && run[position] != '\t') {
		return false
	}
	for position < len(run) && (run[position] == ' ' || run[position] == '\t') {
		position++
	}
	if !strings.HasPrefix(run[position:], "0") {
		return false
	}
	position++
	return position == len(run) || run[position] == ' ' || run[position] == '\t' || run[position] == '\r' || run[position] == '\n' || run[position] == '#'
}

// shellListEnd returns the end of the current AND/OR list. A semicolon starts
// a separate shell list, so a later `exit 0` cannot mask a failed command when
// GitHub Actions runs the script with errexit enabled.
func shellListEnd(run string, start int) int {
	var quote byte
	escaped := false
	inComment := false
	for index := start; index < len(run); index++ {
		character := run[index]
		if inComment {
			if character == '\n' {
				inComment = false
				if shellNewlineSeparates(run, index) {
					return index
				}
			}
			continue
		}
		if escaped {
			escaped = false
			continue
		}
		if character == '\\' && quote != '\'' {
			escaped = true
			continue
		}
		if quote != 0 {
			if character == quote {
				quote = 0
			}
			continue
		}
		if character == '\'' || character == '"' {
			quote = character
			continue
		}
		if character == '#' && (index == 0 || run[index-1] == ' ' || run[index-1] == '\t') {
			inComment = true
			continue
		}
		if character == ';' {
			return index
		}
		if character == '\n' {
			linePrefix := strings.TrimSpace(run[start:index])
			if !strings.HasSuffix(linePrefix, "&&") && !strings.HasSuffix(linePrefix, "||") {
				return index
			}
		}
	}
	return len(run)
}

func shellOperatorLength(run string, index int) int {
	if index+1 < len(run) && run[index] == run[index+1] && (run[index] == '&' || run[index] == '|') {
		return 2
	}
	return 1
}

// shellOperatorIndex finds the first command separator after the doctor
// invocation. Quoted separators are arguments, not shell control operators.
func shellOperatorIndex(run string, start int) int {
	var quote byte
	escaped := false
	inComment := false
	for index := start; index < len(run); index++ {
		character := run[index]
		if inComment {
			if character == '\n' {
				inComment = false
				if shellNewlineSeparates(run, index) {
					return index
				}
			}
			continue
		}
		if escaped {
			escaped = false
			continue
		}
		if character == '\\' && quote != '\'' {
			escaped = true
			continue
		}
		if quote != 0 {
			if character == quote {
				quote = 0
			}
			continue
		}
		if character == '\'' || character == '"' {
			quote = character
			continue
		}
		if character == '#' && (index == 0 || run[index-1] == ' ' || run[index-1] == '\t') {
			inComment = true
			continue
		}
		if character == ';' || character == '|' || character == '&' {
			return index
		}
		if character == '\n' && shellNewlineSeparates(run, index) {
			return index
		}
	}
	return -1
}

func shellNewlineSeparates(run string, index int) bool {
	linePrefix := strings.TrimSpace(run[:index])
	return !strings.HasSuffix(linePrefix, "&&") && !strings.HasSuffix(linePrefix, "||")
}

func yamlMappingValue(node *yaml.Node, key string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	for index := 0; index+1 < len(node.Content); index += 2 {
		if node.Content[index].Value == key {
			return node.Content[index+1]
		}
	}
	return nil
}
