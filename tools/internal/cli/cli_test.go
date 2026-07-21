package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dapi/memory-bank/tools/internal/doctor"
	"github.com/dapi/memory-bank/tools/internal/lint"
)

func testRepository(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", "lint", "testdata", "repository"))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func commitCLISource(t *testing.T, root, message string) string {
	t.Helper()
	if _, err := os.Stat(filepath.Join(root, ".git")); os.IsNotExist(err) {
		runCLIGit(t, root, "init", "--quiet")
	}
	runCLIGit(t, root, "add", "--all")
	runCLIGit(t, root, "-c", "user.name=Memory Bank Tests", "-c", "user.email=tests@example.invalid", "commit", "--quiet", "-m", message)
	return runCLIGit(t, root, "rev-parse", "HEAD")
}

func runCLIGit(t *testing.T, root string, arguments ...string) string {
	t.Helper()
	command := exec.Command("git", append([]string{"-C", root}, arguments...)...)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v failed: %v\n%s", arguments, err, output)
	}
	return strings.TrimSpace(string(output))
}

func TestPrimaryAndCompatibilityEntrypointsHaveLintParity(t *testing.T) {
	arguments := []string{"--repo-root", testRepository(t), "--max-depth", "1", "--json"}
	var primaryStdout, primaryStderr bytes.Buffer
	primaryExit := Run(append([]string{"lint"}, arguments...), "test", &primaryStdout, &primaryStderr)

	var compatibilityStdout, compatibilityStderr bytes.Buffer
	compatibilityExit := RunLint(arguments, "memory-bank-lint", "test", &compatibilityStdout, &compatibilityStderr)

	if primaryExit != compatibilityExit || primaryStdout.String() != compatibilityStdout.String() {
		t.Fatalf("entrypoints differ:\nprimary exit=%d stderr=%q\ncompatibility exit=%d stderr=%q", primaryExit, primaryStderr.String(), compatibilityExit, compatibilityStderr.String())
	}
	var report lint.Report
	if err := json.Unmarshal(primaryStdout.Bytes(), &report); err != nil {
		t.Fatalf("invalid JSON report: %v", err)
	}
	if report.FormatVersion != 1 || report.Stats.BrokenLinkCount != 1 {
		t.Fatalf("unexpected report: %#v", report)
	}
}

func TestRootHelpAndVersion(t *testing.T) {
	for _, test := range []struct {
		arguments []string
		want      string
	}{
		{arguments: []string{"--help"}, want: "Usage: memory-bank <command>"},
		{arguments: []string{"--version"}, want: "memory-bank v1.2.3\n"},
	} {
		var stdout, stderr bytes.Buffer
		if exitCode := Run(test.arguments, "v1.2.3", &stdout, &stderr); exitCode != 0 {
			t.Fatalf("unexpected exit code %d for %v: %s", exitCode, test.arguments, stderr.String())
		}
		if !strings.Contains(stdout.String(), test.want) {
			t.Fatalf("unexpected stdout for %v: %q", test.arguments, stdout.String())
		}
	}
}

func TestRootRejectsMissingAndUnknownCommands(t *testing.T) {
	for _, arguments := range [][]string{nil, {"unknown"}} {
		var stdout, stderr bytes.Buffer
		if exitCode := Run(arguments, "test", &stdout, &stderr); exitCode != 2 {
			t.Fatalf("unexpected exit code %d for %v", exitCode, arguments)
		}
		if !strings.Contains(stderr.String(), "Usage: memory-bank <command>") {
			t.Fatalf("unexpected stderr for %v: %q", arguments, stderr.String())
		}
	}
}

func TestLintRejectsNegativeDepth(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if exitCode := Run([]string{"lint", "--max-depth", "-1"}, "test", &stdout, &stderr); exitCode != 2 {
		t.Fatalf("unexpected exit code: %d", exitCode)
	}
	if !strings.Contains(stderr.String(), "greater than or equal to 0") {
		t.Fatalf("unexpected stderr: %s", stderr.String())
	}
}

func TestCompatibilityHelpAndVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if exitCode := RunLint([]string{"--help"}, "memory-bank-lint", "v1.2.3", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("unexpected help exit code: %d", exitCode)
	}
	if !strings.Contains(stderr.String(), "Usage: memory-bank-lint") {
		t.Fatalf("unexpected help: %q", stderr.String())
	}

	stdout.Reset()
	stderr.Reset()
	if exitCode := RunLint([]string{"--version"}, "memory-bank-lint", "v1.2.3", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("unexpected version exit code: %d", exitCode)
	}
	if stdout.String() != "memory-bank-lint v1.2.3\n" {
		t.Fatalf("unexpected version: %q", stdout.String())
	}
}

func TestOwnershipDryRunJSONReportsPlanWithoutMutation(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(source, "memory-bank", "dna"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "memory-bank", "dna", "rule.md"), []byte("rule\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	sourceRef := commitCLISource(t, source, "initial source")
	arguments := []string{"init", "--repo-root", repo, "--source", source, "--template-version", "v1", "--source-ref", sourceRef, "--dry-run", "--json"}
	var stdout, stderr bytes.Buffer
	if exitCode := Run(arguments, "test", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("unexpected exit %d: %s", exitCode, stderr.String())
	}
	var report struct {
		FormatVersion int  `json:"format_version"`
		DryRun        bool `json:"dry_run"`
		Decisions     []struct {
			Action string `json:"action"`
		} `json:"decisions"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &report); err != nil {
		t.Fatalf("invalid report: %v\n%s", err, stdout.String())
	}
	if report.FormatVersion != 1 || !report.DryRun || len(report.Decisions) != 2 || report.Decisions[0].Action != "create" || report.Decisions[1].Action != "create" {
		t.Fatalf("unexpected report: %#v", report)
	}
	if _, err := os.Stat(filepath.Join(repo, "memory-bank", "dna", "rule.md")); !os.IsNotExist(err) {
		t.Fatalf("dry-run mutated repository: %v", err)
	}
}

func TestDoctorAndAlternativeAgentTarget(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(source, "memory-bank"), 0o755); err != nil {
		t.Fatal(err)
	}
	readme := "---\ndoc_function: index\npurpose: Test index for doctor.\nstatus: active\n---\n# Memory Bank\n"
	if err := os.WriteFile(filepath.Join(source, "memory-bank", "README.md"), []byte(readme), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repo, "CLAUDE.md"), []byte("project rules\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	sourceRef := commitCLISource(t, source, "source")
	args := []string{"init", "--repo-root", repo, "--source", source, "--template-version", "v1", "--source-ref", sourceRef, "--agent-file", "CLAUDE.md"}
	var stdout, stderr bytes.Buffer
	if exitCode := Run(args, "test", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("init failed with %d: %s", exitCode, stderr.String())
	}
	claude, err := os.ReadFile(filepath.Join(repo, "CLAUDE.md"))
	if err != nil || !strings.HasPrefix(string(claude), "project rules\n\n<!-- MEMORY BANK START -->") {
		t.Fatalf("alternative target did not preserve content: %q, %v", claude, err)
	}
	if _, err := os.Stat(filepath.Join(repo, "AGENTS.md")); !os.IsNotExist(err) {
		t.Fatalf("default target was also created: %v", err)
	}
	stdout.Reset()
	stderr.Reset()
	if exitCode := Run([]string{"doctor", "--repo-root", repo, "--agent-file", "CLAUDE.md", "--json"}, "test", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("doctor failed with %d: %s", exitCode, stderr.String())
	}
	var report struct {
		FormatVersion int `json:"format_version"`
		Summary       struct {
			Errors int `json:"errors"`
		} `json:"summary"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &report); err != nil || report.FormatVersion != doctor.ReportFormatVersion || report.Summary.Errors != 0 {
		t.Fatalf("unexpected doctor report: %s, %v", stdout.String(), err)
	}
	lockBefore, err := os.ReadFile(filepath.Join(repo, "memory-bank", ".lock"))
	if err != nil {
		t.Fatal(err)
	}
	outdated := []byte("project rules\n\n<!-- MEMORY BANK START -->\nold\n<!-- MEMORY BANK END -->\n")
	if err := os.WriteFile(filepath.Join(repo, "CLAUDE.md"), outdated, 0o644); err != nil {
		t.Fatal(err)
	}
	stdout.Reset()
	stderr.Reset()
	updateArgs := append([]string{"update"}, args[1:]...)
	if exitCode := Run(updateArgs, "test", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("managed-block update failed with %d: %s", exitCode, stderr.String())
	}
	after, _ := os.ReadFile(filepath.Join(repo, "CLAUDE.md"))
	if !bytes.Equal(after, claude) {
		t.Fatalf("managed-block update did not restore current payload: %q", after)
	}
	lockAfter, err := os.ReadFile(filepath.Join(repo, "memory-bank", ".lock"))
	if err != nil || !bytes.Equal(lockBefore, lockAfter) {
		t.Fatalf("agent-only update changed template lock: %v", err)
	}
	before := append([]byte(nil), after...)
	stdout.Reset()
	stderr.Reset()
	if exitCode := Run(updateArgs, "test", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("idempotent update failed with %d: %s", exitCode, stderr.String())
	}
	after, _ = os.ReadFile(filepath.Join(repo, "CLAUDE.md"))
	if !bytes.Equal(before, after) {
		t.Fatalf("idempotent update changed target")
	}
}

func TestDoctorRejectsUnknownProfile(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if exitCode := Run([]string{"doctor", "--profile", "mystery"}, "test", &stdout, &stderr); exitCode != 2 {
		t.Fatalf("unexpected exit %d", exitCode)
	}
	if !strings.Contains(stderr.String(), "auto, template, or downstream") {
		t.Fatalf("unexpected stderr: %s", stderr.String())
	}
}

func TestDoctorRejectsScopeRootOutsideRepository(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if exitCode := Run([]string{"doctor", "--scope-root", "../sibling"}, "test", &stdout, &stderr); exitCode != 1 {
		t.Fatalf("unexpected exit %d: %s", exitCode, stderr.String())
	}
	if !strings.Contains(stderr.String(), "must not contain parent-directory traversal") {
		t.Fatalf("unexpected stderr: %s", stderr.String())
	}
}

func TestInitRejectsAmbiguousMarkersWithoutMutation(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(source, "memory-bank"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "memory-bank", "README.md"), []byte("template\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	original := []byte("rules\n<!-- MEMORY BANK START -->\n<!-- MEMORY BANK START -->\n<!-- MEMORY BANK END -->\n")
	if err := os.WriteFile(filepath.Join(repo, "AGENTS.md"), original, 0o644); err != nil {
		t.Fatal(err)
	}
	sourceRef := commitCLISource(t, source, "source")
	arguments := []string{"init", "--repo-root", repo, "--source", source, "--template-version", "v1", "--source-ref", sourceRef}
	var stdout, stderr bytes.Buffer
	if exitCode := Run(arguments, "test", &stdout, &stderr); exitCode != 1 {
		t.Fatalf("unexpected exit %d: %s", exitCode, stderr.String())
	}
	current, err := os.ReadFile(filepath.Join(repo, "AGENTS.md"))
	if err != nil || !bytes.Equal(current, original) {
		t.Fatalf("conflict changed user content: %q, %v", current, err)
	}
	if _, err := os.Stat(filepath.Join(repo, "memory-bank", "README.md")); !os.IsNotExist(err) {
		t.Fatalf("conflict partially installed template: %v", err)
	}
	if _, err := os.Stat(filepath.Join(repo, "memory-bank", ".lock")); !os.IsNotExist(err) {
		t.Fatalf("conflict created lock: %v", err)
	}
}

func TestOwnershipDryRunJSONReportsCollisionAsUserOwned(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	dna := filepath.Join(source, "memory-bank", "dna")
	if err := os.MkdirAll(dna, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dna, "seed.md"), []byte("seed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	initialRef := commitCLISource(t, source, "initial source")
	baseArguments := []string{"--repo-root", repo, "--source", source}
	initArguments := append([]string{"init"}, baseArguments...)
	initArguments = append(initArguments, "--template-version", "v1", "--source-ref", initialRef)
	var stdout, stderr bytes.Buffer
	if exitCode := Run(initArguments, "test", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("unexpected init exit %d: %s", exitCode, stderr.String())
	}

	collision := filepath.Join("memory-bank", "dna", "collision.md")
	if err := os.WriteFile(filepath.Join(repo, collision), []byte("downstream\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, collision), []byte("upstream\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	updatedRef := commitCLISource(t, source, "add collision")
	stdout.Reset()
	stderr.Reset()
	updateArguments := append([]string{"update"}, baseArguments...)
	updateArguments = append(updateArguments, "--template-version", "v2", "--source-ref", updatedRef, "--dry-run", "--json")
	if exitCode := Run(updateArguments, "test", &stdout, &stderr); exitCode != 0 {
		t.Fatalf("unexpected update exit %d: %s", exitCode, stderr.String())
	}
	var report struct {
		Decisions []struct {
			Path      string `json:"path"`
			Ownership string `json:"ownership"`
			Action    string `json:"action"`
		} `json:"decisions"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &report); err != nil {
		t.Fatalf("invalid report: %v\n%s", err, stdout.String())
	}
	for _, decision := range report.Decisions {
		if decision.Path == filepath.ToSlash(collision) {
			if decision.Action != "preserve" || decision.Ownership != "user-owned" {
				t.Fatalf("collision decision disagrees with persisted ownership: %#v", decision)
			}
			return
		}
	}
	t.Fatalf("collision decision missing from report: %#v", report.Decisions)
}
