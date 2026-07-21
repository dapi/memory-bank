package ownership

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type pinnedSource struct {
	root string
	info fs.FileInfo
}

func pinSourceRoot(root string) (pinnedSource, error) {
	if root == "" {
		return pinnedSource{}, errors.New("source root is required")
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return pinnedSource{}, fmt.Errorf("resolve source root: %w", err)
	}
	resolvedRoot, err := filepath.EvalSymlinks(absRoot)
	if err != nil {
		return pinnedSource{}, fmt.Errorf("resolve source root: %w", err)
	}
	info, err := os.Lstat(resolvedRoot)
	if err != nil {
		return pinnedSource{}, fmt.Errorf("inspect source root: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return pinnedSource{}, fmt.Errorf("source root is not a directory: %s", resolvedRoot)
	}
	return pinnedSource{root: resolvedRoot, info: info}, nil
}

func inspectSourceRoot(source pinnedSource) error {
	info, err := os.Lstat(source.root)
	if err != nil {
		return fmt.Errorf("inspect source root: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() || !os.SameFile(source.info, info) {
		return fmt.Errorf("unsafe source root: %s changed during update", source.root)
	}
	return nil
}

func rejectOverlappingRoots(repo pinnedRepo, source pinnedSource) error {
	if os.SameFile(repo.info, source.info) || pathContains(repo.root, source.root) || pathContains(source.root, repo.root) {
		return fmt.Errorf("source root and repo root overlap: source=%s repo=%s", source.root, repo.root)
	}
	return nil
}

func pathContains(parent, child string) bool {
	relative, err := filepath.Rel(parent, child)
	if err != nil {
		return false
	}
	return relative == "." || relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator))
}

func verifySourceCheckout(root, expectedRef string) error {
	topLevel, err := gitOutput(root, "rev-parse", "--show-toplevel")
	if err != nil {
		return fmt.Errorf("verify source checkout: %w", err)
	}
	topLevelInfo, err := os.Stat(topLevel)
	if err != nil {
		return fmt.Errorf("inspect Git checkout root %q: %w", topLevel, err)
	}
	rootInfo, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("inspect source root %q: %w", root, err)
	}
	if !os.SameFile(rootInfo, topLevelInfo) {
		return fmt.Errorf("source root must be the Git checkout root: source=%s checkout=%s", root, topLevel)
	}
	head, err := gitOutput(root, "rev-parse", "--verify", "HEAD^{commit}")
	if err != nil {
		return fmt.Errorf("resolve source checkout HEAD: %w", err)
	}
	if !strings.EqualFold(head, expectedRef) {
		return fmt.Errorf("source ref does not match checkout HEAD: got %s, want %s", expectedRef, head)
	}
	status, err := gitOutput(root, "status", "--porcelain=v1", "--untracked-files=all")
	if err != nil {
		return fmt.Errorf("inspect source checkout status: %w", err)
	}
	if status != "" {
		return errors.New("source checkout is dirty; commit or discard changes before init/update")
	}
	payloadStatus, err := gitOutput(root, "status", "--porcelain=v1", "--untracked-files=all", "--ignored=matching", "--", "memory-bank")
	if err != nil {
		return fmt.Errorf("inspect source template status: %w", err)
	}
	if payloadStatus != "" {
		return errors.New("source memory-bank tree contains uncommitted or ignored payloads")
	}
	if err := verifySourcePayload(root, expectedRef); err != nil {
		return err
	}
	return nil
}

func verifySourcePayload(root, expectedRef string) error {
	tree, err := gitOutput(root, "ls-tree", "-rz", "--full-tree", expectedRef, "--", "memory-bank")
	if err != nil {
		return fmt.Errorf("inspect pinned source payload: %w", err)
	}
	expected := make(map[string]string)
	for _, record := range strings.Split(tree, "\x00") {
		if record == "" {
			continue
		}
		header, path, found := strings.Cut(record, "\t")
		fields := strings.Fields(header)
		if !found || len(fields) != 3 {
			return errors.New("inspect pinned source payload: malformed Git tree entry")
		}
		mode, objectType, objectID := fields[0], fields[1], fields[2]
		if objectType != "blob" || mode != "100644" && mode != "100755" {
			return fmt.Errorf("pinned source payload contains unsupported entry: %q", path)
		}
		expected[path] = objectID
	}
	if len(expected) == 0 {
		return errors.New("pinned source commit has no memory-bank payload")
	}

	return nil
}

func gitOutput(root string, arguments ...string) (string, error) {
	output, err := gitBytes(root, arguments...)
	return strings.TrimSpace(string(output)), err
}

func gitBytes(root string, arguments ...string) ([]byte, error) {
	commandArguments := append([]string{"-C", root}, arguments...)
	command := exec.Command("git", commandArguments...)
	command.Env = append(os.Environ(), "GIT_OPTIONAL_LOCKS=0")
	output, err := command.CombinedOutput()
	if err != nil {
		result := strings.TrimSpace(string(output))
		if result == "" {
			return nil, err
		}
		return nil, fmt.Errorf("%s: %w", result, err)
	}
	return output, nil
}
