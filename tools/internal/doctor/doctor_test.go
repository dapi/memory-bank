package doctor

import (
	"encoding/json"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/dapi/memory-bank/tools/internal/ownership"
)

type snapshotEntry struct {
	Mode fs.FileMode
	Data string
}

func snapshot(t *testing.T, root string) map[string]snapshotEntry {
	t.Helper()
	result := map[string]snapshotEntry{}
	if err := filepath.WalkDir(root, func(fullPath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(root, fullPath)
		if err != nil {
			return err
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		item := snapshotEntry{Mode: info.Mode()}
		if info.Mode().IsRegular() {
			data, err := os.ReadFile(fullPath)
			if err != nil {
				return err
			}
			item.Data = string(data)
		}
		result[filepath.ToSlash(relative)] = item
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return result
}

func fixture(t *testing.T, name string) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("testdata", name))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func copyFixture(t *testing.T, name string) string {
	t.Helper()
	source, destination := fixture(t, name), t.TempDir()
	if err := filepath.WalkDir(source, func(sourcePath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(source, sourcePath)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, relative)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		data, err := os.ReadFile(sourcePath)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	}); err != nil {
		t.Fatal(err)
	}
	return destination
}

func runGit(t *testing.T, root string, arguments ...string) string {
	t.Helper()
	command := exec.Command("git", append([]string{"-C", root}, arguments...)...)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v\n%s", arguments, err, output)
	}
	return string(output)
}

func TestProfilesUseSeparateFixturesAndProduceCleanReports(t *testing.T) {
	for _, test := range []struct {
		name string
		want Profile
	}{
		{name: "template", want: ProfileTemplate},
		{name: "downstream", want: ProfileDownstream},
	} {
		t.Run(test.name, func(t *testing.T) {
			report, err := Run(Options{RepoRoot: fixture(t, test.name), ScopeRoot: "memory-bank", AgentFile: "AGENTS.md", Profile: ProfileAuto, MaxDepth: 3})
			if err != nil {
				t.Fatal(err)
			}
			if report.FormatVersion != ReportFormatVersion || report.Profile != test.want || report.Summary.Errors != 0 {
				t.Fatalf("unexpected report: %#v", report)
			}
		})
	}
}

func TestJSONReportContractIsVersionedAndFindingsAreActionable(t *testing.T) {
	report, err := Run(Options{RepoRoot: fixture(t, "template"), ScopeRoot: "missing-bank", AgentFile: "AGENTS.md", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"format_version", "profile", "repo_root", "template_identity", "summary", "findings", "navigation"} {
		if _, exists := payload[field]; !exists {
			t.Fatalf("public JSON field %q is missing: %s", field, data)
		}
	}
	findings, ok := payload["findings"].([]any)
	if !ok || len(findings) == 0 {
		t.Fatalf("expected findings in contract sample: %s", data)
	}
	for _, raw := range findings {
		finding := raw.(map[string]any)
		for _, field := range []string{"code", "severity", "group", "message", "remediation"} {
			if finding[field] == nil || finding[field] == "" {
				t.Fatalf("finding field %q is missing: %#v", field, finding)
			}
		}
	}
}

func TestDoctorDoesNotMutateWorktree(t *testing.T) {
	root := copyFixture(t, "downstream")
	runGit(t, root, "init", "--quiet")
	runGit(t, root, "add", "--all")
	runGit(t, root, "-c", "user.name=Doctor Test", "-c", "user.email=doctor@example.invalid", "commit", "--quiet", "-m", "fixture")
	indexBefore := runGit(t, root, "write-tree")
	statusBefore := runGit(t, root, "status", "--porcelain=v1")
	before := snapshot(t, root)
	if _, err := Run(Options{RepoRoot: root, ScopeRoot: "memory-bank", AgentFile: "AGENTS.md", Profile: ProfileAuto, MaxDepth: 3}); err != nil {
		t.Fatal(err)
	}
	after := snapshot(t, root)
	indexAfter := runGit(t, root, "write-tree")
	statusAfter := runGit(t, root, "status", "--porcelain=v1")
	if !reflect.DeepEqual(before, after) || indexBefore != indexAfter || statusBefore != statusAfter {
		t.Fatalf("doctor mutated fixture\nbefore=%#v\nafter=%#v", before, after)
	}
}

func TestManagedDriftProducesStableFinding(t *testing.T) {
	repo := copyFixture(t, "downstream")
	rulePath := filepath.Join(repo, "memory-bank", "dna", "rule.md")
	if err := os.MkdirAll(filepath.Dir(rulePath), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(rulePath, []byte("---\nstatus: active\n---\n# Rule\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	lockPath := filepath.Join(repo, filepath.FromSlash(ownership.LockFileName))
	data, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatal(err)
	}
	var lock ownership.Lock
	if err := json.Unmarshal(data, &lock); err != nil {
		t.Fatal(err)
	}
	lock.Files["memory-bank/dna/rule.md"] = ownership.File{
		Ownership: ownership.Managed, BaseDigest: "sha256:" + strings.Repeat("0", 64), PayloadDigest: "sha256:" + strings.Repeat("0", 64), BaseMode: "100644", PayloadMode: "100644",
	}
	data, err = json.Marshal(lock)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(lockPath, data, 0o644); err != nil {
		t.Fatal(err)
	}
	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", AgentFile: "AGENTS.md", Profile: ProfileAuto, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range report.Findings {
		if finding.Code == "manifest.managed_content_drift" && finding.Path == "memory-bank/dna/rule.md" {
			return
		}
	}
	t.Fatalf("managed drift finding missing: %#v", report.Findings)
}

func TestGovernanceCycleAndLifecycleFindingsHaveStableCodes(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("AGENTS.md", "Read memory-bank/README.md.\n")
	write("tools/go.mod", "module github.com/dapi/memory-bank/tools\n")
	write(".github/workflows/ci.yml", "run: memory-bank doctor\n")
	write("memory-bank/README.md", "---\ndoc_function: index\npurpose: Root fixture index.\nstatus: active\n---\n# Root\n\n- [Features](features/README.md) — Feature packages and their lifecycle documents.\n")
	write("memory-bank/features/README.md", "---\ndoc_function: index\npurpose: Feature fixture index.\nderived_from:\n  - ../README.md\nstatus: active\n---\n# Features\n\n- [Feature](FT-001/README.md) — Feature package used by doctor tests.\n")
	write("memory-bank/features/FT-001/README.md", "---\ndoc_function: index\npurpose: Feature package fixture.\nderived_from:\n  - brief.md\nstatus: active\n---\n# Feature\n\n- [Brief](brief.md) — Canonical problem and lifecycle owner.\n")
	write("memory-bank/features/FT-001/brief.md", "---\nderived_from:\n  - README.md\nstatus: active\ndelivery_status: in_progress\n---\n# Brief\n")
	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", AgentFile: "AGENTS.md", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	codes := map[string]bool{}
	for _, finding := range report.Findings {
		codes[finding.Code] = true
	}
	for _, code := range []string{"governance.derived_from_cycle", "lifecycle.execution_plan_not_active"} {
		if !codes[code] {
			t.Fatalf("missing %s in %#v", code, report.Findings)
		}
	}
}

func TestGovernanceRejectsInvalidClassificationEnums(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/dna/principles.md", "---\nstatus: active\ndoc_kind: typo\ndoc_function: invalid\n---\n# Principles\n")
	write("memory-bank/flows/README.md", "---\nstatus: active\ndoc_kind: governance\ndoc_function: index\nderived_from:\n  - ../dna/principles.md\n---\n# Flows\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, code := range []string{"governance.doc_kind_invalid", "governance.doc_function_invalid"} {
		if !hasFinding(report, code) {
			t.Fatalf("missing %s in %#v", code, report.Findings)
		}
	}
}

func TestGovernanceLayerMatchingUsesScopeTopLevel(t *testing.T) {
	tests := []struct {
		documentPath string
		want         bool
	}{
		{documentPath: "memory-bank/dna/principles.md", want: true},
		{documentPath: "memory-bank/flows/README.md", want: true},
		{documentPath: "memory-bank/domain/flows/state.md", want: false},
		{documentPath: "memory-bank/domain/dna/state.md", want: false},
	}
	for _, test := range tests {
		t.Run(test.documentPath, func(t *testing.T) {
			if got := isGovernanceLayerDocument(test.documentPath, "memory-bank"); got != test.want {
				t.Fatalf("isGovernanceLayerDocument() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestGovernanceRootUsesExactScopeDNAPath(t *testing.T) {
	if !isGovernanceRoot("memory-bank/dna/principles.md", "memory-bank", true) {
		t.Fatal("scope DNA principles should be a governance root")
	}
	if isGovernanceRoot("memory-bank/domain/dna/principles.md", "memory-bank", true) {
		t.Fatal("nested DNA principles should require derived_from")
	}
}

func TestFeatureBriefOwnsDeliveryStatusWithoutOptionalClassification(t *testing.T) {
	if !isCanonicalFeatureBrief(governedDocument{
		path:        "memory-bank/features/FT-001/brief.md",
		frontmatter: map[string]any{"status": "active", "delivery_status": "planned"},
	}) {
		t.Fatal("canonical feature brief path should own delivery_status without optional classification")
	}
}

func TestWorkflowRunsDoctorOnlyForExecutableRunCommands(t *testing.T) {
	for _, test := range []struct {
		name     string
		workflow string
		want     bool
	}{
		{name: "comment", workflow: "# memory-bank doctor\nname: CI\n", want: false},
		{name: "workflow metadata", workflow: "name: memory-bank doctor\n", want: false},
		{name: "action input", workflow: "jobs:\n  check:\n    steps:\n      - uses: example/action@v1\n        with:\n          run: memory-bank doctor\n", want: false},
		{name: "echo", workflow: "jobs:\n  check:\n    steps:\n      - run: echo memory-bank doctor\n", want: false},
		{name: "printf", workflow: "jobs:\n  check:\n    steps:\n      - run: printf 'memory-bank doctor'\n", want: false},
		{name: "shell suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor || true\n", want: false},
		{name: "arbitrary shell suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor || notify-team\n", want: false},
		{name: "echo shell suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor || echo 'doctor failed'\n", want: false},
		{name: "printf shell suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor || printf '%s\\n' 'doctor failed'\n", want: false},
		{name: "comment does not suppress", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor # never use || true\n", want: true},
		{name: "and-or list suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor && upload-results || true\n", want: false},
		{name: "errexit prevents semicolon suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor; exit 0\n", want: true},
		{name: "explicitly disabled errexit permits semicolon suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: set +e; memory-bank doctor; exit 0\n", want: false},
		{name: "errexit re-enabled before doctor blocks", workflow: "jobs:\n  check:\n    steps:\n      - run: set +e; setup; set -e; memory-bank doctor; cleanup\n", want: true},
		{name: "later suppression does not suppress doctor", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor; cleanup || true\n", want: true},
		{name: "multiline shell suppression", workflow: "jobs:\n  check:\n    steps:\n      - run: |\n          memory-bank doctor ||\n            true\n", want: false},
		{name: "multiline shell suppression with separate commands", workflow: "jobs:\n  check:\n    steps:\n      - run: |\n          set +e\n          memory-bank doctor\n          exit 0\n", want: false},
		{name: "delayed exit after another command suppresses doctor", workflow: "jobs:\n  check:\n    steps:\n      - run: |\n          set +e\n          memory-bank doctor\n          log_result\n          exit 0\n", want: false},
		{name: "custom shell does not provide errexit", workflow: "jobs:\n  check:\n    steps:\n      - shell: bash {0}\n        run: |\n          memory-bank doctor\n          cleanup\n", want: false},
		{name: "custom shell with explicit errexit blocks", workflow: "jobs:\n  check:\n    steps:\n      - shell: bash -e {0}\n        run: |\n          memory-bank doctor\n          cleanup\n", want: true},
		{name: "custom shell without pipefail does not block", workflow: "jobs:\n  check:\n    steps:\n      - shell: bash -e {0}\n        run: memory-bank doctor | tee report.txt\n", want: false},
		{name: "custom shell checks pipeline status", workflow: "jobs:\n  check:\n    steps:\n      - shell: bash -e {0}\n        run: |\n          memory-bank doctor | tee report.txt\n          test ${PIPESTATUS[0]} -eq 0\n", want: true},
		{name: "custom shell with pipefail blocks pipeline", workflow: "jobs:\n  check:\n    steps:\n      - shell: bash -e -o pipefail {0}\n        run: memory-bank doctor | tee report.txt\n", want: true},
		{name: "builtin bash provides errexit and pipefail", workflow: "jobs:\n  check:\n    steps:\n      - shell: bash\n        run: memory-bank doctor | tee report.txt\n", want: true},
		{name: "builtin sh provides errexit", workflow: "jobs:\n  check:\n    steps:\n      - shell: sh\n        run: memory-bank doctor; cleanup\n", want: true},
		{name: "workflow shell default is inherited", workflow: "defaults:\n  run:\n    shell: bash {0}\njobs:\n  check:\n    steps:\n      - run: memory-bank doctor; cleanup\n", want: false},
		{name: "job shell default is inherited", workflow: "jobs:\n  check:\n    defaults:\n      run:\n        shell: bash {0}\n    steps:\n      - run: memory-bank doctor; cleanup\n", want: false},
		{name: "failing recovery preserves gate", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor || false\n", want: true},
		{name: "exit failure recovery preserves gate", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor || exit 1\n", want: true},
		{name: "continue on error", workflow: "jobs:\n  check:\n    steps:\n      - continue-on-error: true\n        run: memory-bank doctor\n", want: false},
		{name: "job continue on error", workflow: "jobs:\n  check:\n    continue-on-error: true\n    steps:\n      - run: memory-bank doctor\n", want: false},
		{name: "job explicit blocking policy", workflow: "jobs:\n  check:\n    continue-on-error: false\n    steps:\n      - run: memory-bank doctor\n", want: true},
		{name: "explicit blocking policy", workflow: "jobs:\n  check:\n    steps:\n      - continue-on-error: false\n        run: memory-bank doctor\n", want: true},
		{name: "expression blocking policy", workflow: "jobs:\n  check:\n    steps:\n      - continue-on-error: ${{ false }}\n        run: memory-bank doctor\n", want: true},
		{name: "direct command", workflow: "jobs:\n  check:\n    steps:\n      - run: memory-bank doctor --profile downstream\n", want: true},
		{name: "multiline after separator", workflow: "jobs:\n  check:\n    steps:\n      - run: |\n          make prepare && memory-bank doctor\n", want: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := workflowRunsDoctor([]byte(test.workflow)); got != test.want {
				t.Fatalf("workflowRunsDoctor() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestNestedScopeReadmeIsGovernanceRoot(t *testing.T) {
	repo := t.TempDir()
	root := filepath.Join(repo, "docs", "memory-bank")
	if err := os.MkdirAll(root, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("---\nstatus: active\n---\n# Root\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "docs/memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	if hasFinding(report, "governance.derived_from_missing") {
		t.Fatalf("nested scope README produced derived_from finding: %#v", report.Findings)
	}
}

func TestScopeReadmeRequiresDependencyWhenDNARootExists(t *testing.T) {
	repo := t.TempDir()
	root := filepath.Join(repo, "docs", "memory-bank")
	if err := os.MkdirAll(filepath.Join(root, "dna"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("---\nstatus: active\n---\n# Root\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "dna", "principles.md"), []byte("---\nstatus: active\ndoc_kind: governance\ndoc_function: canonical\n---\n# Principles\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "docs/memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	if !hasFinding(report, "governance.derived_from_missing") {
		t.Fatalf("scope README without derived_from passed despite DNA root: %#v", report.Findings)
	}
}

func TestParseFrontmatterAcceptsCRLFDelimiters(t *testing.T) {
	frontmatter, found, err := parseFrontmatter([]byte("---\r\nstatus: active\r\n---\r\n# Document\r\n"))
	if err != nil || !found {
		t.Fatalf("parseFrontmatter() found=%t, err=%v; want valid frontmatter", found, err)
	}
	if status := frontmatter["status"]; status != "active" {
		t.Fatalf("status=%#v, want active", status)
	}
}

func TestParseFrontmatterRejectsMalformedClosingDelimiter(t *testing.T) {
	_, found, err := parseFrontmatter([]byte("---\nstatus: active\n---not-a-delimiter\n# Document\n"))
	if !found || err == nil {
		t.Fatalf("parseFrontmatter() found=%t, err=%v; want malformed delimiter error", found, err)
	}
}

func TestDoctorAcceptsCRLFFrontmatter(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/README.md", "---\r\ndoc_function: index\r\npurpose: Root fixture index.\r\nstatus: active\r\n---\r\n# Root\r\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range report.Findings {
		if finding.Code == "governance.frontmatter_missing" || finding.Code == "navigation.index_contract" {
			t.Fatalf("CRLF frontmatter produced unexpected finding: %#v", finding)
		}
	}
}

func TestCancelledFeatureRequiresAbsentOrArchivedPlan(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/features/FT-001/brief.md", "---\nstatus: active\ndelivery_status: cancelled\n---\n# Brief\n")
	write("memory-bank/features/FT-001/implementation-plan.md", "---\nstatus: active\n---\n# Plan\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range report.Findings {
		if finding.Code == "lifecycle.cancelled_plan_not_archived" {
			return
		}
	}
	t.Fatalf("cancelled-plan lifecycle finding missing: %#v", report.Findings)
}

func TestPlanWithoutDesignAllowsExplicitNoDesignDecision(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/features/FT-001/brief.md", "---\nstatus: active\ndelivery_status: in_progress\n---\n# Brief\n\n## Design Requirement Decision\n\n| Decision | Reason |\n| --- | --- |\n| Design required: no | Local change follows an existing pattern. |\n")
	write("memory-bank/features/FT-001/implementation-plan.md", "---\nstatus: active\n---\n# Plan\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range report.Findings {
		if finding.Code == "lifecycle.plan_without_design" {
			t.Fatalf("explicit no-design decision produced unexpected finding: %#v", finding)
		}
	}
}

func TestPlanRequiresActiveUpstreamDocuments(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/features/FT-001/brief.md", "---\nstatus: draft\ndelivery_status: planned\n---\n# Brief\n\n## Design Requirement Decision\n\nDesign required: yes\n")
	write("memory-bank/features/FT-001/design.md", "---\nstatus: draft\n---\n# Design\n")
	write("memory-bank/features/FT-001/implementation-plan.md", "---\nstatus: active\n---\n# Plan\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	codes := map[string]bool{}
	for _, finding := range report.Findings {
		codes[finding.Code] = true
	}
	for _, code := range []string{"lifecycle.plan_brief_not_active", "lifecycle.plan_design_not_active"} {
		if !codes[code] {
			t.Fatalf("missing %s in %#v", code, report.Findings)
		}
	}
}

func TestDesignRequiresActiveBrief(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/features/FT-001/brief.md", "---\nstatus: draft\ndelivery_status: planned\n---\n# Brief\n\n## Design Requirement Decision\n\nDesign required: yes\n")
	write("memory-bank/features/FT-001/design.md", "---\nstatus: draft\n---\n# Design\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	if !hasFinding(report, "lifecycle.plan_brief_not_active") {
		t.Fatalf("design with draft brief did not produce lifecycle finding: %#v", report.Findings)
	}
}

func TestLaterStageArtifactsRequireExplicitDesignDecision(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/features/FT-001/brief.md", "---\nstatus: active\ndelivery_status: planned\n---\n# Brief\n")
	write("memory-bank/features/FT-001/design.md", "---\nstatus: active\n---\n# Design\n")
	write("memory-bank/features/FT-001/implementation-plan.md", "---\nstatus: active\n---\n# Plan\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	if !hasFinding(report, "lifecycle.design_requirement_decision_invalid") {
		t.Fatalf("missing explicit design decision did not produce lifecycle finding: %#v", report.Findings)
	}
}

func TestNoDesignDecisionRejectsDesignArtifact(t *testing.T) {
	repo := t.TempDir()
	write := func(relative, contents string) {
		t.Helper()
		fullPath := filepath.Join(repo, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("memory-bank/features/FT-001/brief.md", "---\nstatus: active\ndelivery_status: planned\n---\n# Brief\n\n## Design Requirement Decision\n\nDesign required: no\n")
	write("memory-bank/features/FT-001/design.md", "---\nstatus: active\n---\n# Design\n")

	report, err := Run(Options{RepoRoot: repo, ScopeRoot: "memory-bank", Profile: ProfileTemplate, MaxDepth: 3})
	if err != nil {
		t.Fatal(err)
	}
	if !hasFinding(report, "lifecycle.design_present_when_not_required") {
		t.Fatalf("design artifact with no-design decision did not produce lifecycle finding: %#v", report.Findings)
	}
}

func TestDerivedFromCycleUsesLintTargetNormalization(t *testing.T) {
	documents := map[string]governedDocument{
		"memory-bank/a.md":        {path: "memory-bank/a.md", frontmatter: map[string]any{"derived_from": "b"}},
		"memory-bank/b/README.md": {path: "memory-bank/b/README.md", frontmatter: map[string]any{"derived_from": "../a.md#contract"}},
	}
	report := Report{}
	report.checkDerivedFromCycles(documents)
	if !hasFinding(report, "governance.derived_from_cycle") {
		t.Fatalf("normalized derived_from cycle was not detected: %#v", report.Findings)
	}
}

func TestDesignDecisionMustBeInDesignRequirementSection(t *testing.T) {
	if decision, valid := featureDesignDecision("# Brief\n\nDesign required: no.\n"); valid || decision != "" {
		t.Fatalf("incidental design decision was accepted: decision=%q valid=%t", decision, valid)
	}
	if decision, valid := featureDesignDecision("# Brief\n\n## Design Requirement Decision\n\n```text\nDesign required: no.\n```\n\nDesign required: yes\n"); !valid || decision != "yes" {
		t.Fatalf("section design decision was not parsed: decision=%q valid=%t", decision, valid)
	}
	if decision, valid := featureDesignDecision("# Brief\n\n## Design Requirement Decision\n\n| `Design required: yes` | Rationale |\n"); !valid || decision != "yes" {
		t.Fatalf("inline-coded full design decision was not parsed: decision=%q valid=%t", decision, valid)
	}
	if decision, valid := featureDesignDecision("# Brief\n\n## Design Requirement Decision\n\n- Design required: no\n"); !valid || decision != "no" {
		t.Fatalf("list-form design decision was not parsed: decision=%q valid=%t", decision, valid)
	}
	if decision, valid := featureDesignDecision("# Brief\n\n## Design Requirement Decision\n\n1. Design required: yes\n"); !valid || decision != "yes" {
		t.Fatalf("numbered-list design decision was not parsed: decision=%q valid=%t", decision, valid)
	}
	if decision, valid := featureDesignDecision("# Brief\n\n```markdown\n## Design Requirement Decision\nDesign required: no\n```\n\n## Design Requirement Decision\n\nDesign required: yes\n"); !valid || decision != "yes" {
		t.Fatalf("fenced example was parsed as the real design decision: decision=%q valid=%t", decision, valid)
	}
	if decision, valid := featureDesignDecision("# Brief\n\n## Design Requirement Decision\n\n### Context\n\nDesign required: yes\n\n## Delivery\n"); !valid || decision != "yes" {
		t.Fatalf("nested decision subsection was not parsed: decision=%q valid=%t", decision, valid)
	}
}

func hasFinding(report Report, code string) bool {
	for _, finding := range report.Findings {
		if finding.Code == code {
			return true
		}
	}
	return false
}
