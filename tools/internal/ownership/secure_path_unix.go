//go:build darwin || linux

package ownership

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"
)

// openDestinationParent resolves ancestors relative to directory descriptors.
// In particular, an ancestor replaced after it was checked cannot redirect a
// later operation through a symlink.
func openDestinationParent(repo pinnedRepo, relative string, create bool, created *[]string) (int, string, error) {
	return openDestinationParentWithMkdir(repo, relative, create, created, unix.Mkdirat)
}

func openDestinationParentWithMkdir(repo pinnedRepo, relative string, create bool, created *[]string, mkdirat func(int, string, uint32) error) (int, string, error) {
	_, osRelative, err := destinationPathLexicalPinned(repo, relative)
	if err != nil {
		return -1, "", err
	}
	parts := strings.Split(osRelative, string(filepath.Separator))
	fd, err := unix.Open(repo.root, unix.O_RDONLY|unix.O_DIRECTORY|unix.O_CLOEXEC|unix.O_NOFOLLOW, 0)
	if err != nil {
		return -1, "", fmt.Errorf("open repository root: %w", err)
	}
	closeFD := func() { _ = unix.Close(fd) }
	for index, part := range parts[:len(parts)-1] {
		next, openErr := unix.Openat(fd, part, unix.O_RDONLY|unix.O_DIRECTORY|unix.O_CLOEXEC|unix.O_NOFOLLOW, 0)
		if errors.Is(openErr, unix.ENOENT) && create {
			mkdirErr := mkdirat(fd, part, 0o755)
			if mkdirErr != nil && !errors.Is(mkdirErr, unix.EEXIST) {
				closeFD()
				return -1, "", mkdirErr
			}
			if mkdirErr == nil && created != nil {
				createdPath := filepath.Join(repo.root, filepath.Join(parts[:index+1]...))
				if rel, relErr := filepath.Rel(repo.root, createdPath); relErr == nil {
					*created = append(*created, filepath.ToSlash(rel))
				}
			}
			next, openErr = unix.Openat(fd, part, unix.O_RDONLY|unix.O_DIRECTORY|unix.O_CLOEXEC|unix.O_NOFOLLOW, 0)
		}
		if openErr != nil {
			closeFD()
			return -1, "", fmt.Errorf("open destination parent %q: %w", relative, openErr)
		}
		_ = unix.Close(fd)
		fd = next
	}
	return fd, parts[len(parts)-1], nil
}

func secureEnsureDestinationParents(repo pinnedRepo, relative string, created *[]string) error {
	fd, _, err := openDestinationParent(repo, relative, true, created)
	if err != nil {
		return err
	}
	return unix.Close(fd)
}

func secureMkdirDestination(repo pinnedRepo, relative string, mode os.FileMode) error {
	fd, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer unix.Close(fd)
	return unix.Mkdirat(fd, leaf, uint32(mode.Perm()))
}

func secureRenameToDestination(repo pinnedRepo, relative, source string) error {
	fd, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer unix.Close(fd)
	return unix.Renameat(unix.AT_FDCWD, source, fd, leaf)
}

func secureRenameFromDestination(repo pinnedRepo, relative, destination string) error {
	fd, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer unix.Close(fd)
	return unix.Renameat(fd, leaf, unix.AT_FDCWD, destination)
}

func secureLinkToDestination(repo pinnedRepo, relative, source string) error {
	fd, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer unix.Close(fd)
	return unix.Linkat(unix.AT_FDCWD, source, fd, leaf, 0)
}

func secureRemoveDestination(repo pinnedRepo, relative string, directory bool) error {
	fd, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer unix.Close(fd)
	flags := 0
	if directory {
		flags = unix.AT_REMOVEDIR
	}
	return unix.Unlinkat(fd, leaf, flags)
}

func secureReadDestination(repo pinnedRepo, relative string) (os.FileInfo, []byte, error) {
	return secureReadDestinationWithParentOpened(repo, relative, nil)
}

func secureReadDestinationWithParentOpened(repo pinnedRepo, relative string, parentOpened func()) (os.FileInfo, []byte, error) {
	parent, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return nil, nil, err
	}
	defer unix.Close(parent)
	if parentOpened != nil {
		parentOpened()
	}

	fd, err := unix.Openat(parent, leaf, unix.O_RDONLY|unix.O_CLOEXEC|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("open destination file %q: %w", relative, err)
	}
	file := os.NewFile(uintptr(fd), relative)
	if file == nil {
		_ = unix.Close(fd)
		return nil, nil, fmt.Errorf("open destination file %q", relative)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return nil, nil, fmt.Errorf("unsupported destination file %q", relative)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}
	return info, data, nil
}
