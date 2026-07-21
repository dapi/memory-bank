package repository

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveRootUsesExplicitPath(t *testing.T) {
	root := t.TempDir()
	resolved, err := ResolveRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if resolved != root {
		t.Fatalf("unexpected root: got %q, want %q", resolved, root)
	}
}

func TestFindNearestGitRoot(t *testing.T) {
	repositoryRoot := t.TempDir()
	if err := os.Mkdir(filepath.Join(repositoryRoot, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	nestedDirectory := filepath.Join(repositoryRoot, "one", "two")
	if err := os.MkdirAll(nestedDirectory, 0o755); err != nil {
		t.Fatal(err)
	}

	got, ok := findNearestGitRoot(nestedDirectory)
	if !ok {
		t.Fatal("expected to find nearest .git root")
	}
	if got != repositoryRoot {
		t.Fatalf("unexpected git root: got %q, want %q", got, repositoryRoot)
	}
}
