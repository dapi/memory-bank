//go:build darwin || linux

package ownership

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/sys/unix"
)

func TestConcurrentParentCreationIsNotRecordedAsTransactionOwned(t *testing.T) {
	repo, err := pinRepoRoot(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	created := []string{}
	mkdirat := func(fd int, path string, mode uint32) error {
		if err := unix.Mkdirat(fd, path, mode); err != nil {
			return err
		}
		// Model another process winning the race after Openat reported ENOENT.
		return unix.EEXIST
	}

	fd, _, err := openDestinationParentWithMkdir(repo, "memory-bank/dna/rule.md", true, &created, mkdirat)
	if err != nil {
		t.Fatal(err)
	}
	if closeErr := unix.Close(fd); closeErr != nil && !errors.Is(closeErr, unix.EBADF) {
		t.Fatal(closeErr)
	}
	if len(created) != 0 {
		t.Fatalf("concurrently created parent was recorded for rollback: %v", created)
	}
}

func TestSecureReadDestinationKeepsPinnedParentAfterAncestorReplacement(t *testing.T) {
	root := t.TempDir()
	insideParent := filepath.Join(root, "memory-bank", "dna")
	if err := os.MkdirAll(insideParent, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(insideParent, "rule.md"), []byte("inside"), 0o644); err != nil {
		t.Fatal(err)
	}
	outside := t.TempDir()
	if err := os.MkdirAll(filepath.Join(outside, "dna"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(outside, "dna", "rule.md"), []byte("outside"), 0o644); err != nil {
		t.Fatal(err)
	}
	repo, err := pinRepoRoot(root)
	if err != nil {
		t.Fatal(err)
	}

	var replaceErr error
	_, data, err := secureReadDestinationWithParentOpened(repo, "memory-bank/dna/rule.md", func() {
		memoryBank := filepath.Join(root, "memory-bank")
		replaceErr = os.Rename(memoryBank, filepath.Join(root, "original-memory-bank"))
		if replaceErr == nil {
			replaceErr = os.Symlink(outside, memoryBank)
		}
	})
	if replaceErr != nil {
		t.Fatal(replaceErr)
	}
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "inside" {
		t.Fatalf("read was redirected after ancestor replacement: %q", data)
	}
}
