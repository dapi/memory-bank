package ownership

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestInitRejectsOverlappingSourceAndRepoRoots(t *testing.T) {
	for _, test := range []struct {
		name   string
		source func(*testing.T, string) string
	}{
		{name: "same root", source: func(_ *testing.T, repo string) string { return repo }},
		{name: "symlink alias", source: func(t *testing.T, repo string) string {
			alias := filepath.Join(t.TempDir(), "source-alias")
			symlinkForTest(t, repo, alias)
			return alias
		}},
		{name: "nested source", source: func(t *testing.T, repo string) string {
			nested := filepath.Join(repo, "template")
			if err := os.Mkdir(nested, 0o755); err != nil {
				t.Fatal(err)
			}
			return nested
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			repo := t.TempDir()
			source := test.source(t, repo)
			write(t, source, "memory-bank/dna/rule.md", "downstream\n")
			report, err := Init(opts(repo, source, "a"))
			if err == nil || !strings.Contains(err.Error(), "overlap") {
				t.Fatalf("expected source/repo overlap error, got report=%#v err=%v", report, err)
			}
			if _, err := os.Lstat(filepath.Join(repo, LockFileName)); !os.IsNotExist(err) {
				t.Fatalf("overlapping source created a lock: %v", err)
			}
		})
	}
}

func TestUpdateCannotUseDownstreamRepoAsTemplateSource(t *testing.T) {
	repo, source := t.TempDir(), t.TempDir()
	path := "memory-bank/dna/rule.md"
	write(t, source, path, "template\n")
	initialize(t, repo, source)
	write(t, repo, path, "downstream drift\n")
	lockBefore := read(t, repo, LockFileName)

	report, err := Update(opts(repo, repo, "b"))
	if err == nil || !strings.Contains(err.Error(), "overlap") {
		t.Fatalf("expected self-source rejection, got report=%#v err=%v", report, err)
	}
	if got := read(t, repo, path); got != "downstream drift\n" {
		t.Fatalf("self-source update changed downstream payload: %q", got)
	}
	if got := read(t, repo, LockFileName); got != lockBefore {
		t.Fatal("self-source update changed the lock")
	}
}

func TestSourceRefMustMatchCleanGitCheckout(t *testing.T) {
	source := t.TempDir()
	write(t, source, "memory-bank/dna/rule.md", "committed\n")
	commit := commitTestSource(t, source)

	options := Options{
		RepoRoot:        t.TempDir(),
		SourceRoot:      source,
		TemplateVersion: "v1",
		SourceRef:       strings.Repeat("f", 40),
		DryRun:          true,
	}
	if _, err := Init(options); err == nil || !strings.Contains(err.Error(), "does not match") {
		t.Fatalf("expected mismatched source ref error, got %v", err)
	}

	options.RepoRoot = t.TempDir()
	options.SourceRef = commit
	write(t, source, "memory-bank/dna/rule.md", "dirty\n")
	if _, err := Init(options); err == nil || !strings.Contains(err.Error(), "dirty") {
		t.Fatalf("expected dirty source checkout error, got %v", err)
	}

	write(t, source, "memory-bank/dna/rule.md", "committed\n")
	options.RepoRoot = t.TempDir()
	report, err := Init(options)
	if err != nil || !report.DryRun || report.Applied {
		t.Fatalf("clean matching checkout was rejected: report=%#v err=%v", report, err)
	}
}

func TestPinnedSourceObjectsIgnoreHiddenWorktreeChanges(t *testing.T) {
	for _, test := range []struct {
		name   string
		flag   string
		mutate func(*testing.T, string)
	}{
		{
			name: "assume unchanged modified file",
			flag: "--assume-unchanged",
			mutate: func(t *testing.T, source string) {
				write(t, source, "memory-bank/dna/rule.md", "modified but hidden\n")
			},
		},
		{
			name: "skip worktree missing file",
			flag: "--skip-worktree",
			mutate: func(t *testing.T, source string) {
				if err := os.Remove(filepath.Join(source, "memory-bank/dna/rule.md")); err != nil {
					t.Fatal(err)
				}
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			source := t.TempDir()
			write(t, source, "memory-bank/dna/rule.md", "committed\n")
			commit := commitTestSource(t, source)
			runGitTest(t, source, "update-index", test.flag, "memory-bank/dna/rule.md")
			test.mutate(t, source)
			if status := runGitTest(t, source, "status", "--porcelain=v1", "--untracked-files=all"); status != "" {
				t.Fatalf("test mutation was not hidden from porcelain status: %q", status)
			}

			options := Options{
				RepoRoot:        t.TempDir(),
				SourceRoot:      source,
				TemplateVersion: "v1",
				SourceRef:       commit,
			}
			report, err := Init(options)
			if err != nil || !report.Applied {
				t.Fatalf("pinned object install failed: report=%#v err=%v", report, err)
			}
			if got := read(t, options.RepoRoot, "memory-bank/dna/rule.md"); got != "committed\n" {
				t.Fatalf("installed worktree mutation instead of pinned blob: %q", got)
			}
		})
	}
}

func TestCleanCheckoutWithTextConversionUsesPinnedBlob(t *testing.T) {
	source := t.TempDir()
	write(t, source, "memory-bank/dna/rule.md", "canonical\n")
	write(t, source, ".gitattributes", "*.md text eol=crlf\n")
	commit := commitTestSource(t, source)
	if err := os.Remove(filepath.Join(source, "memory-bank/dna/rule.md")); err != nil {
		t.Fatal(err)
	}
	runGitTest(t, source, "checkout", "--", "memory-bank/dna/rule.md")
	if got := read(t, source, "memory-bank/dna/rule.md"); got != "canonical\r\n" {
		t.Fatalf("fixture did not apply CRLF checkout conversion: %q", got)
	}
	if status := runGitTest(t, source, "status", "--porcelain=v1"); status != "" {
		t.Fatalf("text-converted checkout is not clean: %q", status)
	}
	repo := t.TempDir()
	report, err := Init(Options{RepoRoot: repo, SourceRoot: source, TemplateVersion: "v1", SourceRef: commit})
	if err != nil || !report.Applied {
		t.Fatalf("clean text-converted checkout was rejected: report=%#v err=%v", report, err)
	}
	if got := read(t, repo, "memory-bank/dna/rule.md"); got != "canonical\n" {
		t.Fatalf("installed non-canonical worktree bytes: %q", got)
	}
}

func TestPinnedSourceExecutableModeIsInstalled(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not expose Unix executable permission bits")
	}
	source := t.TempDir()
	path := "memory-bank/flows/tool.md"
	write(t, source, path, "tool\n")
	if err := os.Chmod(filepath.Join(source, filepath.FromSlash(path)), 0o755); err != nil {
		t.Fatal(err)
	}
	commit := commitTestSource(t, source)
	repo := t.TempDir()
	report, err := Init(Options{RepoRoot: repo, SourceRoot: source, TemplateVersion: "v1", SourceRef: commit})
	if err != nil || !report.Applied {
		t.Fatalf("executable source install failed: report=%#v err=%v", report, err)
	}
	info, err := os.Stat(filepath.Join(repo, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode().Perm(); got != 0o755 {
		t.Fatalf("executable source mode was lost: got %04o", got)
	}
}

func commitTestSource(t *testing.T, root string) string {
	t.Helper()
	runGitTest(t, root, "init", "--quiet")
	runGitTest(t, root, "add", "--all")
	runGitTest(t, root, "-c", "user.name=Memory Bank Tests", "-c", "user.email=tests@example.invalid", "commit", "--quiet", "-m", "source")
	return runGitTest(t, root, "rev-parse", "HEAD")
}

func runGitTest(t *testing.T, root string, arguments ...string) string {
	t.Helper()
	command := exec.Command("git", append([]string{"-C", root}, arguments...)...)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v failed: %v\n%s", arguments, err, output)
	}
	return strings.TrimSpace(string(output))
}
