package ownership

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestTransactionStagesAllPayloadsBeforeFirstMutation(t *testing.T) {
	repo := t.TempDir()
	first := "memory-bank/dna/a.md"
	second := "memory-bank/dna/b.md"
	write(t, repo, first, "a1\n")
	write(t, repo, second, "b1\n")

	stageErr := errors.New("simulated staging failure")
	writeCalls := 0
	renameCalls := 0
	mutationCalls := 0
	options := Options{
		RepoRoot: repo,
		BeforeMutation: func(Decision) error {
			mutationCalls++
			return nil
		},
	}
	mutations := []mutation{
		{
			decision:       Decision{Path: first, Action: UpdateFile},
			data:           []byte("a2\n"),
			expectedExists: true,
		},
		{
			decision:       Decision{Path: second, Action: UpdateFile},
			data:           []byte("b2\n"),
			expectedExists: true,
		},
	}
	ops := transactionOps{
		writeFile: func(path string, data []byte, mode os.FileMode) error {
			writeCalls++
			if writeCalls == 2 {
				return stageErr
			}
			return os.WriteFile(path, data, mode)
		},
		rename: func(oldPath, newPath string) error {
			renameCalls++
			return os.Rename(oldPath, newPath)
		},
	}

	err := applyAtomicallyWithOps(options, mutations, ops)
	if !errors.Is(err, stageErr) {
		t.Fatalf("expected staging error, got %v", err)
	}
	if mutationCalls != 0 || renameCalls != 0 {
		t.Fatalf("staging failure reached mutation phase: hooks=%d renames=%d", mutationCalls, renameCalls)
	}
	if got := read(t, repo, first); got != "a1\n" {
		t.Fatalf("first target changed during staging: %q", got)
	}
	if got := read(t, repo, second); got != "b1\n" {
		t.Fatalf("second target changed during staging: %q", got)
	}
	assertNoTransactionStaging(t, repo)
}

func TestTransactionDoesNotDefaultExplicitZeroMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not expose Unix permission bits")
	}
	repo := t.TempDir()
	stageErr := errors.New("stop after observing staged mode")
	var stagedMode os.FileMode
	mutations := []mutation{{
		decision: Decision{Path: "AGENTS.md", Action: Create},
		data:     []byte("instructions\n"),
		mode:     0,
		modeSet:  true,
	}}
	ops := transactionOps{
		writeFile: func(_ string, _ []byte, mode os.FileMode) error {
			stagedMode = mode
			return stageErr
		},
	}

	err := applyAtomicallyWithOps(Options{RepoRoot: repo}, mutations, ops)
	if !errors.Is(err, stageErr) {
		t.Fatalf("expected staging error, got %v", err)
	}
	if stagedMode != 0 {
		t.Fatalf("explicit zero mode defaulted to %04o", stagedMode)
	}
	assertNoTransactionStaging(t, repo)
}

func TestTransactionSurfacesRollbackFailure(t *testing.T) {
	repo := t.TempDir()
	first := "memory-bank/dna/a.md"
	second := "memory-bank/dna/b.md"
	write(t, repo, first, "a1\n")
	write(t, repo, second, "b1\n")

	applyErr := errors.New("simulated interruption")
	rollbackErr := errors.New("simulated rollback failure")
	mutationCalls := 0
	options := Options{
		RepoRoot: repo,
		BeforeMutation: func(Decision) error {
			mutationCalls++
			if mutationCalls == 2 {
				return applyErr
			}
			return nil
		},
	}
	mutations := []mutation{
		{
			decision:       Decision{Path: first, Action: UpdateFile},
			data:           []byte("a2\n"),
			expectedExists: true,
		},
		{
			decision:       Decision{Path: second, Action: UpdateFile},
			data:           []byte("b2\n"),
			expectedExists: true,
		},
	}
	pinnedRoot, err := pinRepoRoot(repo)
	if err != nil {
		t.Fatal(err)
	}
	firstTarget, err := destinationPathPinned(pinnedRoot, first)
	if err != nil {
		t.Fatal(err)
	}
	ops := transactionOps{
		writeFile: os.WriteFile,
		rename:    os.Rename,
		link: func(oldPath, newPath string) error {
			if newPath == firstTarget && filepath.Base(filepath.Dir(oldPath)) == "old" {
				return rollbackErr
			}
			return os.Link(oldPath, newPath)
		},
	}

	err = applyAtomicallyWithOps(options, mutations, ops)
	if !errors.Is(err, applyErr) {
		t.Fatalf("apply error was lost: %v", err)
	}
	if !errors.Is(err, rollbackErr) {
		t.Fatalf("rollback error was lost: %v", err)
	}
	if !strings.Contains(err.Error(), "rollback incomplete") {
		t.Fatalf("rollback failure lacks context: %v", err)
	}
	if got := read(t, repo, second); got != "b1\n" {
		t.Fatalf("unreached target changed: %q", got)
	}
	staging, err := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
	if err != nil {
		t.Fatal(err)
	}
	if len(staging) != 1 {
		t.Fatalf("expected retained recovery staging, got %v", staging)
	}
	if got := read(t, staging[0], "old/000000"); got != "a1\n" {
		t.Fatalf("recovery staging lost the original payload: %q", got)
	}
}

func TestLockMutationFailureRestoresAllPayloads(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	first := "memory-bank/dna/a.md"
	second := "memory-bank/dna/b.md"
	write(t, source, first, "a1\n")
	write(t, source, second, "b1\n")
	initialize(t, repo, source)
	write(t, source, first, "a2\n")
	write(t, source, second, "b2\n")
	lockBefore := read(t, repo, LockFileName)

	lockErr := errors.New("simulated lock interruption")
	sawLock := false
	options := opts(repo, source, "b")
	options.BeforeMutation = func(decision Decision) error {
		if decision.Path == LockFileName {
			sawLock = true
			return lockErr
		}
		return nil
	}

	_, err := Update(options)
	if !errors.Is(err, lockErr) {
		t.Fatalf("expected lock mutation error, got %v", err)
	}
	if !sawLock {
		t.Fatal("update never reached the lock mutation")
	}
	if got := read(t, repo, first); got != "a1\n" {
		t.Fatalf("first payload was not restored: %q", got)
	}
	if got := read(t, repo, second); got != "b1\n" {
		t.Fatalf("second payload was not restored: %q", got)
	}
	if got := read(t, repo, LockFileName); got != lockBefore {
		t.Fatal("failed lock mutation changed the lock")
	}
	assertNoTransactionStaging(t, repo)
}

func TestTopologyTransitionRollsBackOnLockFailure(t *testing.T) {
	for _, test := range []struct {
		name        string
		initialPath string
		initialData string
		updatedPath string
		updatedData string
	}{
		{name: "file to directory", initialPath: "memory-bank/dna/topic", initialData: "old file\n", updatedPath: "memory-bank/dna/topic/page.md", updatedData: "new child\n"},
		{name: "directory to file", initialPath: "memory-bank/dna/topic/page.md", initialData: "old child\n", updatedPath: "memory-bank/dna/topic", updatedData: "new file\n"},
	} {
		t.Run(test.name, func(t *testing.T) {
			repo, source := t.TempDir(), t.TempDir()
			write(t, source, test.initialPath, test.initialData)
			initialize(t, repo, source)
			lockBefore := read(t, repo, LockFileName)
			if err := os.RemoveAll(filepath.Join(source, "memory-bank", "dna", "topic")); err != nil {
				t.Fatal(err)
			}
			write(t, source, test.updatedPath, test.updatedData)

			interruption := errors.New("simulated lock failure")
			options := opts(repo, source, "b")
			options.BeforeMutation = func(decision Decision) error {
				if decision.Path == LockFileName {
					return interruption
				}
				return nil
			}
			if _, err := Update(options); !errors.Is(err, interruption) {
				t.Fatalf("expected lock failure, got %v", err)
			}
			if got := read(t, repo, test.initialPath); got != test.initialData {
				t.Fatalf("rollback did not restore original topology: %q", got)
			}
			if got := read(t, repo, LockFileName); got != lockBefore {
				t.Fatal("failed topology transition changed the lock")
			}
			assertNoTransactionStaging(t, repo)
		})
	}
}

func TestTransactionRollbackRemovesCreatedDirectories(t *testing.T) {
	repo := t.TempDir()
	first := "memory-bank/new/nested/a.md"
	second := "memory-bank/trigger.md"
	interruption := errors.New("simulated interruption")
	mutationCalls := 0
	options := Options{
		RepoRoot: repo,
		BeforeMutation: func(Decision) error {
			mutationCalls++
			if mutationCalls == 2 {
				return interruption
			}
			return nil
		},
	}
	mutations := []mutation{
		{decision: Decision{Path: first, Action: Create}, data: []byte("created\n")},
		{decision: Decision{Path: second, Action: Create}, data: []byte("never reached\n")},
	}

	err := applyAtomically(options, mutations)
	if !errors.Is(err, interruption) {
		t.Fatalf("expected interruption, got %v", err)
	}
	if _, err := os.Lstat(filepath.Join(repo, "memory-bank")); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("rollback left transaction-created directories behind: %v", err)
	}
	assertNoTransactionStaging(t, repo)
}

func TestTransactionRestoresOriginalWhenInstallRenameFails(t *testing.T) {
	repo := t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, repo, path, "one\n")
	installErr := errors.New("simulated install failure")
	mutations := []mutation{{
		decision:       Decision{Path: path, Action: UpdateFile},
		data:           []byte("two\n"),
		expectedExists: true,
		expectedDigest: digest([]byte("one\n")),
	}}
	ops := transactionOps{
		writeFile: os.WriteFile,
		rename:    os.Rename,
		link: func(oldPath, newPath string) error {
			if filepath.Base(filepath.Dir(oldPath)) == "new" {
				return installErr
			}
			return os.Link(oldPath, newPath)
		},
	}

	err := applyAtomicallyWithOps(Options{RepoRoot: repo}, mutations, ops)
	if !errors.Is(err, installErr) {
		t.Fatalf("expected install error, got %v", err)
	}
	if got := read(t, repo, path); got != "one\n" {
		t.Fatalf("failed install did not restore original: %q", got)
	}
	assertNoTransactionStaging(t, repo)
}

func TestTransactionPreservesEditMadeToAppliedTargetBeforeRollback(t *testing.T) {
	repo := t.TempDir()
	first := "memory-bank/dna/a.md"
	second := "memory-bank/dna/b.md"
	write(t, repo, first, "a1\n")
	write(t, repo, second, "b1\n")
	interruption := errors.New("simulated interruption")
	options := Options{
		RepoRoot: repo,
		BeforeMutation: func(decision Decision) error {
			if decision.Path != second {
				return nil
			}
			write(t, repo, first, "concurrent edit\n")
			return interruption
		},
	}
	mutations := []mutation{
		{
			decision:       Decision{Path: first, Action: UpdateFile},
			data:           []byte("a2\n"),
			expectedExists: true,
			expectedDigest: digest([]byte("a1\n")),
		},
		{
			decision:       Decision{Path: second, Action: UpdateFile},
			data:           []byte("b2\n"),
			expectedExists: true,
			expectedDigest: digest([]byte("b1\n")),
		},
	}

	err := applyAtomically(options, mutations)
	if !errors.Is(err, interruption) || !strings.Contains(err.Error(), "rollback incomplete") {
		t.Fatalf("expected interruption and incomplete rollback, got %v", err)
	}
	if got := read(t, repo, first); got != "concurrent edit\n" {
		t.Fatalf("rollback overwrote a concurrent edit: %q", got)
	}
	if got := read(t, repo, second); got != "b1\n" {
		t.Fatalf("unreached target changed: %q", got)
	}
	staging, err := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
	if err != nil || len(staging) != 1 {
		t.Fatalf("expected retained recovery staging, got %v err=%v", staging, err)
	}
	if got := read(t, staging[0], "old/000000"); got != "a1\n" {
		t.Fatalf("recovery staging lost pre-update original: %q", got)
	}
}

func TestTransactionRetainsRecoveryDataOnPanic(t *testing.T) {
	repo := t.TempDir()
	first := "memory-bank/dna/a.md"
	second := "memory-bank/dna/b.md"
	write(t, repo, first, "a1\n")
	write(t, repo, second, "b1\n")
	mutationCalls := 0
	options := Options{
		RepoRoot: repo,
		BeforeMutation: func(Decision) error {
			mutationCalls++
			if mutationCalls == 2 {
				panic("simulated panic")
			}
			return nil
		},
	}
	mutations := []mutation{
		{decision: Decision{Path: first, Action: UpdateFile}, data: []byte("a2\n"), expectedExists: true},
		{decision: Decision{Path: second, Action: UpdateFile}, data: []byte("b2\n"), expectedExists: true},
	}
	var recovered any
	func() {
		defer func() { recovered = recover() }()
		_ = applyAtomically(options, mutations)
	}()
	if recovered == nil {
		t.Fatal("expected simulated panic")
	}
	staging, err := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
	if err != nil || len(staging) != 1 {
		t.Fatalf("panic discarded recovery staging: %v err=%v", staging, err)
	}
	if got := read(t, staging[0], "old/000000"); got != "a1\n" {
		t.Fatalf("panic discarded original payload: %q", got)
	}
}

func TestTransactionRejectsTamperedStagedPayload(t *testing.T) {
	repo := t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, repo, path, "one\n")
	options := Options{
		RepoRoot: repo,
		BeforeMutation: func(Decision) error {
			staging, err := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
			if err != nil {
				return fmt.Errorf("locate staging: %w", err)
			}
			if len(staging) != 1 {
				return fmt.Errorf("locate staging: got %v", staging)
			}
			return os.WriteFile(filepath.Join(staging[0], "new", "000000"), []byte("tampered\n"), 0o644)
		},
	}
	mutations := []mutation{{
		decision:       Decision{Path: path, Action: UpdateFile},
		data:           []byte("two\n"),
		expectedExists: true,
		expectedDigest: digest([]byte("one\n")),
	}}

	err := applyAtomically(options, mutations)
	if err == nil || !strings.Contains(err.Error(), "staged payload changed") {
		t.Fatalf("expected staged-payload integrity error, got %v", err)
	}
	if got := read(t, repo, path); got != "one\n" {
		t.Fatalf("tampered staged payload replaced original: %q", got)
	}
	assertNoTransactionStaging(t, repo)
}

func TestTransactionReportsCleanupFailureAfterCommit(t *testing.T) {
	repo := t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, repo, path, "one\n")
	cleanupErr := errors.New("simulated cleanup failure")
	mutations := []mutation{{
		decision:       Decision{Path: path, Action: UpdateFile},
		data:           []byte("two\n"),
		expectedExists: true,
		expectedDigest: digest([]byte("one\n")),
	}}
	ops := osTransactionOps
	ops.removeAll = func(string) error { return cleanupErr }

	err := applyAtomicallyWithOps(Options{RepoRoot: repo}, mutations, ops)
	if !errors.Is(err, cleanupErr) || !strings.Contains(err.Error(), "update committed") {
		t.Fatalf("expected explicit committed-cleanup error, got %v", err)
	}
	var committed *committedError
	if !errors.As(err, &committed) {
		t.Fatalf("cleanup error did not preserve committed outcome: %v", err)
	}
	if got := read(t, repo, path); got != "two\n" {
		t.Fatalf("cleanup failure changed committed payload: %q", got)
	}
	staging, globErr := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
	if globErr != nil || len(staging) != 1 {
		t.Fatalf("expected retained staging, got %v err=%v", staging, globErr)
	}
}

func TestTransactionDoesNotReplaceDestinationThatAppearsDuringInstall(t *testing.T) {
	repo := t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, repo, path, "one\n")
	mutations := []mutation{{
		decision:       Decision{Path: path, Action: UpdateFile},
		data:           []byte("two\n"),
		expectedExists: true,
		expectedDigest: digest([]byte("one\n")),
	}}
	ops := osTransactionOps
	ops.link = func(oldPath, newPath string) error {
		if filepath.Base(filepath.Dir(oldPath)) == "new" {
			if err := os.WriteFile(newPath, []byte("concurrent\n"), 0o644); err != nil {
				return err
			}
		}
		return os.Link(oldPath, newPath)
	}

	err := applyAtomicallyWithOps(Options{RepoRoot: repo}, mutations, ops)
	if err == nil || !strings.Contains(err.Error(), "rollback incomplete") {
		t.Fatalf("expected no-clobber install failure, got %v", err)
	}
	if got := read(t, repo, path); got != "concurrent\n" {
		t.Fatalf("install replaced concurrently-created destination: %q", got)
	}
	staging, globErr := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
	if globErr != nil || len(staging) != 1 {
		t.Fatalf("expected retained recovery staging, got %v err=%v", staging, globErr)
	}
	if got := read(t, staging[0], "old/000000"); got != "one\n" {
		t.Fatalf("recovery staging lost original payload: %q", got)
	}
}

func TestLockCommitRejectsAndRestoresChangedOriginalBackup(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, source, path, "one\n")
	initialize(t, repo, source)
	write(t, source, path, "two\n")
	lockBefore := read(t, repo, LockFileName)

	options := opts(repo, source, "b")
	options.BeforeMutation = func(decision Decision) error {
		if decision.Path != LockFileName {
			return nil
		}
		staging, err := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
		if err != nil {
			return err
		}
		if len(staging) != 1 {
			return fmt.Errorf("expected one staging directory, got %v", staging)
		}
		return os.WriteFile(filepath.Join(staging[0], "old", "000000"), []byte("late drift\n"), 0o644)
	}

	_, err := Update(options)
	if err == nil || !strings.Contains(err.Error(), "backup content changed") {
		t.Fatalf("expected changed-backup error, got %v", err)
	}
	if got := read(t, repo, path); got != "late drift\n" {
		t.Fatalf("rollback lost the latest downstream bytes: %q", got)
	}
	if got := read(t, repo, LockFileName); got != lockBefore {
		t.Fatal("lock committed after the original backup changed")
	}
	assertNoTransactionStaging(t, repo)
}

func assertNoTransactionStaging(t *testing.T, repo string) {
	t.Helper()
	staging, err := filepath.Glob(filepath.Join(repo, ".memory-bank-update-*"))
	if err != nil {
		t.Fatal(err)
	}
	if len(staging) != 0 {
		t.Fatalf("transaction staging was not cleaned up: %v", staging)
	}
}
