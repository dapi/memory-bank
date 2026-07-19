package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dapi/memory-bank/internal/audit"
)

func TestRunJSONReportAndExitCode(t *testing.T) {
	repositoryRoot, err := filepath.Abs(filepath.Join("..", "..", "internal", "audit", "testdata", "repository"))
	if err != nil {
		t.Fatal(err)
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := run([]string{"--repo-root", repositoryRoot, "--max-depth", "1", "--json"}, &stdout, &stderr)
	if exitCode != 1 {
		t.Fatalf("unexpected exit code: %d; stderr: %s", exitCode, stderr.String())
	}
	var report audit.Report
	if err := json.Unmarshal(stdout.Bytes(), &report); err != nil {
		t.Fatalf("invalid JSON report: %v\n%s", err, stdout.String())
	}
	if report.FormatVersion != 1 || report.Stats.BrokenLinkCount != 1 {
		t.Fatalf("unexpected report: %#v", report)
	}
}

func TestRunRejectsNegativeDepth(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if exitCode := run([]string{"--max-depth", "-1"}, &stdout, &stderr); exitCode != 2 {
		t.Fatalf("unexpected exit code: %d", exitCode)
	}
	if !strings.Contains(stderr.String(), "greater than or equal to 0") {
		t.Fatalf("unexpected stderr: %s", stderr.String())
	}
}
