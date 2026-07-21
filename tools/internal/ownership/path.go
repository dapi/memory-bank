package ownership

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type pinnedRepo struct {
	root string
	info fs.FileInfo
}

func pinRepoRoot(root string) (pinnedRepo, error) {
	if root == "" {
		return pinnedRepo{}, errors.New("repo root is required")
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return pinnedRepo{}, fmt.Errorf("resolve repo root: %w", err)
	}
	resolvedRoot, err := filepath.EvalSymlinks(absRoot)
	if err != nil {
		return pinnedRepo{}, fmt.Errorf("resolve repo root: %w", err)
	}
	info, err := inspectRepoRoot(resolvedRoot, nil)
	if err != nil {
		return pinnedRepo{}, err
	}
	return pinnedRepo{root: resolvedRoot, info: info}, nil
}

func inspectRepoRoot(root string, expected fs.FileInfo) (fs.FileInfo, error) {
	info, err := os.Lstat(root)
	if err != nil {
		return nil, fmt.Errorf("inspect repo root: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return nil, fmt.Errorf("unsafe repo root: %s is a symlink", root)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("repo root is not a directory: %s", root)
	}
	if expected != nil && !os.SameFile(expected, info) {
		return nil, fmt.Errorf("unsafe repo root: %s changed during update", root)
	}
	return info, nil
}

// destinationPath resolves a repository-relative path without following any
// symlink below repoRoot. A missing suffix is allowed so callers can safely
// plan creation of a new file.
func destinationPathPinned(repo pinnedRepo, relative string) (string, error) {
	target, osRelative, err := destinationPathLexicalPinned(repo, relative)
	if err != nil {
		return "", err
	}

	current := repo.root
	components := strings.Split(osRelative, string(filepath.Separator))
	for index, component := range components {
		current = filepath.Join(current, component)
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) {
			return target, nil
		}
		if err != nil {
			return "", fmt.Errorf("inspect destination path %q: %w", relative, err)
		}
		componentPath, relErr := filepath.Rel(repo.root, current)
		if relErr != nil {
			componentPath = current
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("unsafe destination path %q: component %q is a symlink", relative, filepath.ToSlash(componentPath))
		}
		if index < len(components)-1 && !info.IsDir() {
			return "", fmt.Errorf("unsafe destination path %q: component %q is not a directory", relative, filepath.ToSlash(componentPath))
		}
	}
	return target, nil
}

// destinationPathLexicalPinned validates a repository-relative path without
// requiring its current ancestors to be directories. Topology transitions use
// it while a clean managed ancestor is still waiting to be removed.
func destinationPathLexicalPinned(repo pinnedRepo, relative string) (string, string, error) {
	if _, err := inspectRepoRoot(repo.root, repo.info); err != nil {
		return "", "", err
	}
	if relative == "" || strings.Contains(relative, "\\") {
		return "", "", fmt.Errorf("unsafe destination path %q", relative)
	}
	osRelative := filepath.FromSlash(relative)
	if !filepath.IsLocal(osRelative) || filepath.ToSlash(filepath.Clean(osRelative)) != relative {
		return "", "", fmt.Errorf("unsafe destination path %q", relative)
	}
	return filepath.Join(repo.root, osRelative), osRelative, nil
}
