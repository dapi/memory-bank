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

func TestWorkflowRunsDoctorOnlyForExecutableRunCommands(t *testing.T) {
	for _, test := range []struct {
		name     string
		workflow string
		want     bool
	}{
		{name: "comment", workflow: "# memory-bank doctor\nname: CI\n", want: false},
		{name: "workflow metadata", workflow: "name: memory-bank doctor\n", want: false},
		{name: "echo", workflow: "jobs:\n  check:\n    steps:\n      - run: echo memory-bank doctor\n", want: false},
		{name: "printf", workflow: "jobs:\n  check:\n    steps:\n      - run: printf 'memory-bank doctor'\n", want: false},
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
	write("memory-bank/features/FT-001/brief.md", "---\nstatus: draft\ndelivery_status: planned\n---\n# Brief\n\nDesign required: yes\n")
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
