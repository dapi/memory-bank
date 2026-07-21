package ownership

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type topologySnapshot struct {
	files       []destinationPrecondition
	directories []string
}

type removedDirectory struct {
	path string
	mode fs.FileMode
}

func inspectDestinationForPlan(repo pinnedRepo, relative string, cleanRemovals map[string]string) (string, bool, *topologySnapshot, error) {
	target, osRelative, err := destinationPathLexicalPinned(repo, relative)
	if err != nil {
		return "", false, nil, err
	}
	current := repo.root
	components := splitPath(osRelative)
	for index, component := range components {
		current = filepath.Join(current, component)
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) {
			return "", false, nil, nil
		}
		if err != nil {
			return "", false, nil, fmt.Errorf("inspect destination path %q: %w", relative, err)
		}
		componentPath, err := filepath.Rel(repo.root, current)
		if err != nil {
			return "", false, nil, err
		}
		componentPath = filepath.ToSlash(componentPath)
		if info.Mode()&os.ModeSymlink != 0 {
			return "", false, nil, fmt.Errorf("unsafe destination path %q: component %q is a symlink", relative, componentPath)
		}
		if index < len(components)-1 {
			if info.IsDir() {
				continue
			}
			if !info.Mode().IsRegular() || cleanRemovals[componentPath] == "" {
				return "", false, nil, fmt.Errorf("destination topology for %q is blocked by %q", relative, componentPath)
			}
			digest, err := digestRegularDestination(repo, componentPath, info)
			if err != nil {
				return "", false, nil, err
			}
			if digest != cleanRemovals[componentPath] {
				return "", false, nil, fmt.Errorf("destination topology changed at %q", componentPath)
			}
			return "", false, &topologySnapshot{files: []destinationPrecondition{{path: componentPath, digest: digest}}}, nil
		}
		if info.Mode().IsRegular() {
			digest, err := digestRegularDestination(repo, relative, info)
			return digest, true, nil, err
		}
		if !info.IsDir() {
			return "", false, nil, fmt.Errorf("unsupported destination file %q", relative)
		}
		snapshot, err := snapshotDirectory(repo, target)
		if err != nil {
			return "", false, nil, err
		}
		if len(snapshot.files) == 0 {
			return "", false, nil, fmt.Errorf("destination topology for %q is an untracked directory", relative)
		}
		for _, file := range snapshot.files {
			if cleanRemovals[file.path] != file.digest {
				return "", false, nil, fmt.Errorf("destination topology for %q contains non-removable file %q", relative, file.path)
			}
		}
		impliedDirectories := map[string]bool{relative: true}
		for _, file := range snapshot.files {
			for directory := filepath.ToSlash(filepath.Dir(filepath.FromSlash(file.path))); pathContainsRelative(relative, directory); directory = filepath.ToSlash(filepath.Dir(filepath.FromSlash(directory))) {
				impliedDirectories[directory] = true
				if directory == relative {
					break
				}
			}
		}
		for _, directory := range snapshot.directories {
			if !impliedDirectories[directory] {
				return "", false, nil, fmt.Errorf("destination topology for %q contains untracked directory %q", relative, directory)
			}
		}
		return "", false, snapshot, nil
	}
	return "", false, nil, nil
}

func pathContainsRelative(parent, child string) bool {
	return child == parent || strings.HasPrefix(child, parent+"/")
}

func splitPath(relative string) []string {
	return strings.Split(relative, string(filepath.Separator))
}

func snapshotDirectory(repo pinnedRepo, root string) (*topologySnapshot, error) {
	result := &topologySnapshot{}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(repo.root, path)
		if err != nil {
			return err
		}
		relative = filepath.ToSlash(relative)
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("unsafe destination path %q: component is a symlink", relative)
		}
		if info.IsDir() {
			result.directories = append(result.directories, relative)
			return nil
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("unsupported destination file %q", relative)
		}
		payloadDigest, err := digestRegularDestination(repo, relative, info)
		if err != nil {
			return err
		}
		result.files = append(result.files, destinationPrecondition{path: relative, digest: payloadDigest})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(result.files, func(i, j int) bool { return result.files[i].path < result.files[j].path })
	sort.Strings(result.directories)
	return result, nil
}

func digestRegularDestination(repo pinnedRepo, relative string, expected fs.FileInfo) (string, error) {
	info, data, err := secureReadDestination(repo, relative)
	if err != nil {
		return "", err
	}
	if !os.SameFile(expected, info) {
		return "", fmt.Errorf("destination file changed while reading: %s", relative)
	}
	return digest(data), nil
}

func verifyTopologySnapshot(repo pinnedRepo, relative string, expected *topologySnapshot) error {
	cleanRemovals := make(map[string]string, len(expected.files))
	for _, file := range expected.files {
		cleanRemovals[file.path] = file.digest
	}
	_, exists, current, err := inspectDestinationForPlan(repo, relative, cleanRemovals)
	if err != nil {
		return err
	}
	if exists || current == nil || !sameTopology(current, expected) {
		return errors.New("destination topology changed while update was being planned")
	}
	return nil
}

func sameTopology(left, right *topologySnapshot) bool {
	if len(left.files) != len(right.files) || len(left.directories) != len(right.directories) {
		return false
	}
	for index := range left.files {
		if left.files[index] != right.files[index] {
			return false
		}
	}
	for index := range left.directories {
		if left.directories[index] != right.directories[index] {
			return false
		}
	}
	return true
}

func prepareTopologyDestination(repo pinnedRepo, relative string, expected *topologySnapshot, removed *[]removedDirectory) error {
	target, _, err := destinationPathLexicalPinned(repo, relative)
	if err != nil {
		return err
	}
	info, err := os.Lstat(target)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return errors.New("destination topology did not clear after planned deletions")
	}
	current, err := snapshotDirectory(repo, target)
	if err != nil {
		return err
	}
	if len(current.files) != 0 || !sameStrings(current.directories, expected.directories) {
		return errors.New("destination directory topology changed during update")
	}
	for index := len(current.directories) - 1; index >= 0; index-- {
		directoryPath := current.directories[index]
		directoryTarget, err := destinationPathPinned(repo, directoryPath)
		if err != nil {
			return err
		}
		directoryInfo, err := os.Lstat(directoryTarget)
		if err != nil {
			return err
		}
		if !directoryInfo.IsDir() || directoryInfo.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("destination directory changed before removal: %s", directoryPath)
		}
		if err := secureRemoveDestination(repo, directoryPath, true); err != nil {
			return fmt.Errorf("remove replaced directory %s: %w", directoryPath, err)
		}
		*removed = append(*removed, removedDirectory{path: directoryPath, mode: directoryInfo.Mode().Perm()})
	}
	return nil
}

func sameStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
