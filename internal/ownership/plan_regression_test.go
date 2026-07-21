package ownership

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestDeletedManagedFileIsDownstreamDrift(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, source, path, "managed\n")
	initialize(t, repo, source)
	lockBefore := read(t, repo, LockFileName)
	if err := os.Remove(filepath.Join(repo, filepath.FromSlash(path))); err != nil {
		t.Fatal(err)
	}

	report, err := Update(opts(repo, source, "b"))
	decision := decisionFor(t, report, path)
	if err != nil || report.Applied || report.ConflictCount != 1 || decision.Action != Conflict || decision.Reason != "managed file has downstream drift" {
		t.Fatalf("deleted managed file was not treated as drift: report=%#v err=%v", report, err)
	}
	if _, err := os.Stat(filepath.Join(repo, filepath.FromSlash(path))); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("deleted managed file was restored: %v", err)
	}
	if lockAfter := read(t, repo, LockFileName); lockAfter != lockBefore {
		t.Fatal("conflicting managed deletion changed the lock")
	}
}

func TestDeletedGeneratedFileIsRegenerated(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/.generated/index.json"
	write(t, source, path, "{\"generated\":true}\n")
	initialize(t, repo, source)
	if err := os.Remove(filepath.Join(repo, filepath.FromSlash(path))); err != nil {
		t.Fatal(err)
	}

	report, err := Update(opts(repo, source, "b"))
	decision := decisionFor(t, report, path)
	if err != nil || !report.Applied || report.ConflictCount != 0 || decision.Action != UpdateFile || decision.Ownership != Generated {
		t.Fatalf("deleted generated file was not regenerated: report=%#v err=%v", report, err)
	}
	if got := read(t, repo, path); got != "{\"generated\":true}\n" {
		t.Fatalf("unexpected regenerated payload: %q", got)
	}
}

func TestInitRegeneratesExistingGeneratedFile(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/.generated/index.json"
	write(t, repo, path, "stale\n")
	write(t, source, path, "generated\n")

	report, err := Init(opts(repo, source, "a"))
	decision := decisionFor(t, report, path)
	if err != nil || !report.Applied || decision.Action != UpdateFile || decision.Ownership != Generated {
		t.Fatalf("init did not regenerate deterministic payload: report=%#v err=%v", report, err)
	}
	if got := read(t, repo, path); got != "generated\n" {
		t.Fatalf("unexpected initialized generated payload: %q", got)
	}
}

func TestCollisionReportsAndPersistsUserOwnership(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	seed := "memory-bank/dna/seed.md"
	path := "memory-bank/dna/collision.md"
	write(t, source, seed, "seed\n")
	initialize(t, repo, source)
	write(t, repo, path, "downstream\n")
	write(t, source, path, "upstream\n")
	lockBefore := read(t, repo, LockFileName)

	dryRunOptions := opts(repo, source, "b")
	dryRunOptions.DryRun = true
	report, err := Update(dryRunOptions)
	decision := decisionFor(t, report, path)
	if err != nil || report.Applied || !report.DryRun || report.ConflictCount != 0 || decision.Action != Preserve || decision.Ownership != UserOwned {
		t.Fatalf("dry-run reported the wrong collision ownership: report=%#v err=%v", report, err)
	}
	if got := read(t, repo, path); got != "downstream\n" {
		t.Fatalf("dry-run overwrote the collision: %q", got)
	}
	if lockAfter := read(t, repo, LockFileName); lockAfter != lockBefore {
		t.Fatal("dry-run changed the lock")
	}

	report, err = Update(opts(repo, source, "b"))
	decision = decisionFor(t, report, path)
	if err != nil || !report.Applied || report.ConflictCount != 0 || decision.Action != Preserve || decision.Ownership != UserOwned {
		t.Fatalf("update reported the wrong collision ownership: report=%#v err=%v", report, err)
	}
	if got := read(t, repo, path); got != "downstream\n" {
		t.Fatalf("update overwrote the collision: %q", got)
	}
	lock, exists, err := ReadLock(repo)
	if err != nil || !exists {
		t.Fatalf("could not read updated lock: exists=%v err=%v", exists, err)
	}
	if got := lock.Files[path]; got != (File{Ownership: UserOwned}) {
		t.Fatalf("collision ownership was not persisted: %#v", got)
	}

	report, err = Update(opts(repo, source, "b"))
	decision = decisionFor(t, report, path)
	if err != nil || report.Applied || decision.Action != Preserve || decision.Ownership != UserOwned {
		t.Fatalf("repeated update reported the wrong collision ownership: report=%#v err=%v", report, err)
	}
}

func TestManagedEditAfterPlanningIsNotOverwritten(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, source, path, "one\n")
	initialize(t, repo, source)
	write(t, source, path, "two\n")
	lockBefore := read(t, repo, LockFileName)

	options := opts(repo, source, "b")
	options.Now = func() time.Time {
		write(t, repo, path, "late drift\n")
		return fixedTime
	}
	if _, err := Update(options); err == nil || !strings.Contains(err.Error(), "content changed") {
		t.Fatalf("expected late-drift error, got %v", err)
	}
	if got := read(t, repo, path); got != "late drift\n" {
		t.Fatalf("late downstream edit was overwritten: %q", got)
	}
	if got := read(t, repo, LockFileName); got != lockBefore {
		t.Fatal("late downstream edit changed the lock")
	}
}

func TestManagedEditAfterStagingIsNotOverwritten(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, source, path, "one\n")
	initialize(t, repo, source)
	write(t, source, path, "two\n")
	lockBefore := read(t, repo, LockFileName)

	options := opts(repo, source, "b")
	options.BeforeMutation = func(decision Decision) error {
		if decision.Path == path {
			write(t, repo, path, "late drift\n")
		}
		return nil
	}
	if _, err := Update(options); err == nil || !strings.Contains(err.Error(), "content changed") {
		t.Fatalf("expected late-drift error, got %v", err)
	}
	if got := read(t, repo, path); got != "late drift\n" {
		t.Fatalf("late downstream edit was overwritten: %q", got)
	}
	if got := read(t, repo, LockFileName); got != lockBefore {
		t.Fatal("late downstream edit changed the lock")
	}
}

func TestPreservedManagedEditBeforeLockCommitIsRejected(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, source, path, "managed\n")
	initialize(t, repo, source)
	lockBefore := read(t, repo, LockFileName)

	options := opts(repo, source, "b")
	options.Now = func() time.Time {
		write(t, repo, path, "late drift\n")
		return fixedTime
	}
	if _, err := Update(options); err == nil || !strings.Contains(err.Error(), "before lock commit") {
		t.Fatalf("expected lock-precondition error, got %v", err)
	}
	if got := read(t, repo, path); got != "late drift\n" {
		t.Fatalf("late downstream edit was overwritten: %q", got)
	}
	if got := read(t, repo, LockFileName); got != lockBefore {
		t.Fatal("lock was committed after a preserved managed file drifted")
	}
}

func TestIncomingPayloadResolvesManagedAndAdaptedConflicts(t *testing.T) {
	for _, test := range []struct {
		name      string
		path      string
		ownership Class
	}{
		{name: "managed", path: "memory-bank/dna/rule.md", ownership: Managed},
		{name: "adapted", path: "memory-bank/domain/model.md", ownership: Adapted},
	} {
		t.Run(test.name, func(t *testing.T) {
			repo, source := t.TempDir(), t.TempDir()
			write(t, source, test.path, "base\n")
			initialize(t, repo, source)
			write(t, repo, test.path, "downstream\n")
			write(t, source, test.path, "incoming\n")

			report, err := Update(opts(repo, source, "b"))
			if err != nil || report.ConflictCount != 1 || decisionFor(t, report, test.path).Action != Conflict {
				t.Fatalf("expected initial conflict: report=%#v err=%v", report, err)
			}

			write(t, repo, test.path, "incoming\n")
			report, err = Update(opts(repo, source, "b"))
			decision := decisionFor(t, report, test.path)
			if err != nil || !report.Applied || report.ConflictCount != 0 || decision.Action != Preserve || !strings.Contains(decision.Reason, "matches incoming") {
				t.Fatalf("incoming resolution was not accepted: report=%#v err=%v", report, err)
			}
			lock, exists, err := ReadLock(repo)
			if err != nil || !exists {
				t.Fatalf("could not read resolved lock: exists=%v err=%v", exists, err)
			}
			want := File{Ownership: test.ownership, BaseDigest: digest([]byte("incoming\n")), BaseMode: "100644"}
			if test.ownership == Managed {
				want.PayloadDigest = want.BaseDigest
				want.PayloadMode = want.BaseMode
			}
			if got := lock.Files[test.path]; got != want {
				t.Fatalf("resolved lock did not advance to incoming base: got=%#v want=%#v", got, want)
			}

			report, err = Update(opts(repo, source, "b"))
			if err != nil || report.Applied || report.ConflictCount != 0 {
				t.Fatalf("resolved update was not idempotent: report=%#v err=%v", report, err)
			}
		})
	}
}

func TestNestedArtifactREADMEStaysUserOwned(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/features/FT-001/README.md"
	write(t, source, path, "seed\n")
	initialize(t, repo, source)
	write(t, repo, path, "project package\n")
	write(t, source, path, "changed seed\n")

	report, err := Update(opts(repo, source, "b"))
	decision := decisionFor(t, report, path)
	if err != nil || !report.Applied || decision.Action != Preserve || decision.Ownership != UserOwned {
		t.Fatalf("nested README was not preserved as user-owned: report=%#v err=%v", report, err)
	}
	if got := read(t, repo, path); got != "project package\n" {
		t.Fatalf("nested README was overwritten: %q", got)
	}
	lock, exists, err := ReadLock(repo)
	if err != nil || !exists || lock.Files[path].Ownership != UserOwned {
		t.Fatalf("nested README lock ownership is wrong: lock=%#v exists=%v err=%v", lock.Files[path], exists, err)
	}

	if err := os.Remove(filepath.Join(source, filepath.FromSlash(path))); err != nil {
		t.Fatal(err)
	}
	report, err = Update(opts(repo, source, "c"))
	if err != nil || !report.Applied || decisionFor(t, report, path).Action != Preserve {
		t.Fatalf("nested README was not preserved after upstream removal: report=%#v err=%v", report, err)
	}
	if got := read(t, repo, path); got != "project package\n" {
		t.Fatalf("nested README was deleted after upstream removal: %q", got)
	}
}

func TestCleanManagedPathTopologyTransitions(t *testing.T) {
	for _, test := range []struct {
		name        string
		initialPath string
		initialData string
		updatedPath string
		updatedData string
	}{
		{
			name:        "file to directory",
			initialPath: "memory-bank/dna/topic",
			initialData: "topic index\n",
			updatedPath: "memory-bank/dna/topic/page.md",
			updatedData: "topic page\n",
		},
		{
			name:        "directory to file",
			initialPath: "memory-bank/dna/topic/page.md",
			initialData: "topic page\n",
			updatedPath: "memory-bank/dna/topic",
			updatedData: "topic index\n",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			repo, source := t.TempDir(), t.TempDir()
			write(t, source, test.initialPath, test.initialData)
			initialize(t, repo, source)
			if err := os.RemoveAll(filepath.Join(source, "memory-bank", "dna", "topic")); err != nil {
				t.Fatal(err)
			}
			write(t, source, test.updatedPath, test.updatedData)

			report, err := Update(opts(repo, source, "b"))
			if err != nil || !report.Applied || report.ConflictCount != 0 {
				t.Fatalf("clean topology transition failed: report=%#v err=%v", report, err)
			}
			if decisionFor(t, report, test.initialPath).Action != Delete {
				t.Fatalf("old topology was not deleted: %#v", report.Decisions)
			}
			if decisionFor(t, report, test.updatedPath).Action != Create {
				t.Fatalf("new topology was not created: %#v", report.Decisions)
			}
			if got := read(t, repo, test.updatedPath); got != test.updatedData {
				t.Fatalf("unexpected transitioned payload: %q", got)
			}
			if strings.HasPrefix(test.initialPath, test.updatedPath+"/") {
				if _, err := os.Lstat(filepath.Join(repo, filepath.FromSlash(test.initialPath))); err == nil {
					t.Fatal("old descendant still exists after directory-to-file transition")
				}
			}
			lock, exists, err := ReadLock(repo)
			if err != nil || !exists {
				t.Fatalf("could not read transitioned lock: exists=%v err=%v", exists, err)
			}
			if _, exists := lock.Files[test.initialPath]; exists {
				t.Fatalf("old topology remains in lock: %#v", lock.Files[test.initialPath])
			}
			if got := lock.Files[test.updatedPath]; got.Ownership != Managed || got.PayloadDigest != digest([]byte(test.updatedData)) {
				t.Fatalf("new topology missing from lock: %#v", got)
			}
		})
	}
}

func TestDirectoryToFileTransitionRejectsUntrackedTopology(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	oldPath := "memory-bank/dna/topic/page.md"
	newPath := "memory-bank/dna/topic"
	write(t, source, oldPath, "old child\n")
	initialize(t, repo, source)
	lockBefore := read(t, repo, LockFileName)
	if err := os.Mkdir(filepath.Join(repo, "memory-bank", "dna", "topic", "project-only"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(filepath.Join(source, "memory-bank", "dna", "topic")); err != nil {
		t.Fatal(err)
	}
	write(t, source, newPath, "new file\n")

	if _, err := Update(opts(repo, source, "b")); err == nil || !strings.Contains(err.Error(), "untracked directory") {
		t.Fatalf("expected untracked topology rejection, got %v", err)
	}
	if got := read(t, repo, oldPath); got != "old child\n" {
		t.Fatalf("rejected transition changed managed payload: %q", got)
	}
	if got := read(t, repo, LockFileName); got != lockBefore {
		t.Fatal("rejected transition changed the lock")
	}
}
