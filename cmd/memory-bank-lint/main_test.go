package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRunJSONReportAndExitCode(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := run([]string{"--repo-root", "testdata/repository", "--max-depth", "1", "--json"}, &stdout, &stderr)
	if exitCode != 1 {
		t.Fatalf("unexpected exit code: %d; stderr: %s", exitCode, stderr.String())
	}
	var report Report
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

func TestRunPrintsVersion(t *testing.T) {
	previousVersion := version
	version = "v1.2.3"
	defer func() { version = previousVersion }()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if exitCode := run([]string{"--version"}, &stdout, &stderr); exitCode != 0 {
		t.Fatalf("unexpected exit code: %d; stderr: %s", exitCode, stderr.String())
	}
	if stdout.String() != "memory-bank-lint v1.2.3\n" {
		t.Fatalf("unexpected stdout: %q", stdout.String())
	}
}

func TestRunHelpSucceeds(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if exitCode := run([]string{"--help"}, &stdout, &stderr); exitCode != 0 {
		t.Fatalf("unexpected exit code: %d", exitCode)
	}
	if !strings.Contains(stderr.String(), "Usage: memory-bank-lint") {
		t.Fatalf("unexpected stderr: %s", stderr.String())
	}
}
