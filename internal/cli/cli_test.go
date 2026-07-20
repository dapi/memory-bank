package cli

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dapi/memory-bank/internal/lint"
)

func testRepository(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", "lint", "testdata", "repository"))
	if err != nil {
		t.Fatal(err)
	}
	return root
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
	for _, arguments := range [][]string{nil, {"doctor"}} {
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
