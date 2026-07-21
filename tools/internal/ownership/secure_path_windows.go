//go:build windows

package ownership

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

const fileCreated = 2

type fileNameInformation struct {
	ReplaceIfExists uint32
	RootDirectory   windows.Handle
	FileNameLength  uint32
	FileName        [1]uint16
}

func openRepoHandle(repo pinnedRepo) (windows.Handle, error) {
	name, err := windows.UTF16PtrFromString(repo.root)
	if err != nil {
		return 0, err
	}
	handle, err := windows.CreateFile(name, windows.FILE_GENERIC_READ, windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE, nil, windows.OPEN_EXISTING, windows.FILE_FLAG_BACKUP_SEMANTICS|windows.FILE_FLAG_OPEN_REPARSE_POINT, 0)
	if err != nil {
		return 0, err
	}
	file := os.NewFile(uintptr(handle), repo.root)
	info, statErr := file.Stat()
	if statErr != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 || !os.SameFile(repo.info, info) {
		_ = file.Close()
		if statErr != nil {
			return 0, statErr
		}
		return 0, fmt.Errorf("unsafe repo root: %s changed during update", repo.root)
	}
	// The caller owns the handle; do not close the temporary os.File wrapper.
	runtime.SetFinalizer(file, nil)
	return handle, nil
}

func ntOpenRelative(parent windows.Handle, name string, access, disposition, options uint32) (windows.Handle, uintptr, error) {
	objectName, err := windows.NewNTUnicodeString(name)
	if err != nil {
		return 0, 0, err
	}
	attributes := windows.OBJECT_ATTRIBUTES{
		Length:        uint32(unsafe.Sizeof(windows.OBJECT_ATTRIBUTES{})),
		RootDirectory: parent,
		ObjectName:    objectName,
		Attributes:    windows.OBJ_CASE_INSENSITIVE | windows.OBJ_DONT_REPARSE,
	}
	var status windows.IO_STATUS_BLOCK
	var handle windows.Handle
	err = windows.NtCreateFile(&handle, access, &attributes, &status, nil, 0,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE,
		disposition, options|windows.FILE_OPEN_REPARSE_POINT, 0, 0)
	return handle, status.Information, err
}

func openDestinationParent(repo pinnedRepo, relative string, create bool, created *[]string) (windows.Handle, string, error) {
	_, osRelative, err := destinationPathLexicalPinned(repo, relative)
	if err != nil {
		return 0, "", err
	}
	parts := strings.Split(osRelative, string(filepath.Separator))
	handle, err := openRepoHandle(repo)
	if err != nil {
		return 0, "", fmt.Errorf("open repository root: %w", err)
	}
	for index, part := range parts[:len(parts)-1] {
		disposition := uint32(windows.FILE_OPEN)
		if create {
			disposition = windows.FILE_OPEN_IF
		}
		next, information, openErr := ntOpenRelative(handle, part, windows.FILE_GENERIC_READ|windows.FILE_GENERIC_WRITE, disposition, windows.FILE_DIRECTORY_FILE)
		windows.CloseHandle(handle)
		if openErr != nil {
			return 0, "", fmt.Errorf("open destination parent %q: %w", relative, openErr)
		}
		handle = next
		if information == fileCreated && created != nil {
			*created = append(*created, filepath.ToSlash(filepath.Join(parts[:index+1]...)))
		}
	}
	return handle, parts[len(parts)-1], nil
}

func secureEnsureDestinationParents(repo pinnedRepo, relative string, created *[]string) error {
	handle, _, err := openDestinationParent(repo, relative, true, created)
	if err == nil {
		err = windows.CloseHandle(handle)
	}
	return err
}

func secureMkdirDestination(repo pinnedRepo, relative string, mode os.FileMode) error {
	parent, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(parent)
	handle, _, err := ntOpenRelative(parent, leaf, windows.FILE_GENERIC_READ|windows.FILE_GENERIC_WRITE, windows.FILE_CREATE, windows.FILE_DIRECTORY_FILE)
	if err == nil {
		err = windows.CloseHandle(handle)
	}
	_ = mode // Windows has no Unix directory permission mode.
	return err
}

func openAbsoluteForMutation(path string, directory bool) (windows.Handle, error) {
	name, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return 0, err
	}
	flags := uint32(windows.FILE_FLAG_OPEN_REPARSE_POINT)
	if directory {
		flags |= windows.FILE_FLAG_BACKUP_SEMANTICS
	}
	return windows.CreateFile(name, windows.FILE_GENERIC_READ|windows.FILE_GENERIC_WRITE|windows.DELETE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE, nil, windows.OPEN_EXISTING, flags, 0)
}

func nameInformation(root windows.Handle, leaf string, replace bool) ([]byte, error) {
	name, err := windows.UTF16FromString(leaf)
	if err != nil {
		return nil, err
	}
	name = name[:len(name)-1]
	var header fileNameInformation
	buffer := make([]byte, int(unsafe.Offsetof(header.FileName))+len(name)*2)
	info := (*fileNameInformation)(unsafe.Pointer(&buffer[0]))
	if replace {
		info.ReplaceIfExists = windows.FILE_RENAME_REPLACE_IF_EXISTS | windows.FILE_RENAME_POSIX_SEMANTICS
	}
	info.RootDirectory = root
	info.FileNameLength = uint32(len(name) * 2)
	copy(unsafe.Slice(&info.FileName[0], len(name)), name)
	return buffer, nil
}

func setName(handle, root windows.Handle, leaf string, class uint32, replace bool) error {
	buffer, err := nameInformation(root, leaf, replace)
	if err != nil {
		return err
	}
	var status windows.IO_STATUS_BLOCK
	return windows.NtSetInformationFile(handle, &status, &buffer[0], uint32(len(buffer)), class)
}

func secureRenameToDestination(repo pinnedRepo, relative, source string) error {
	parent, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(parent)
	sourceHandle, err := openAbsoluteForMutation(source, false)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(sourceHandle)
	return setName(sourceHandle, parent, leaf, windows.FileRenameInformation, false)
}

func secureRenameFromDestination(repo pinnedRepo, relative, destination string) error {
	parent, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(parent)
	source, _, err := ntOpenRelative(parent, leaf, windows.FILE_GENERIC_READ|windows.FILE_GENERIC_WRITE|windows.DELETE, windows.FILE_OPEN, 0)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(source)
	destinationParent, err := openAbsoluteForMutation(filepath.Dir(destination), true)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(destinationParent)
	return setName(source, destinationParent, filepath.Base(destination), windows.FileRenameInformation, false)
}

func secureLinkToDestination(repo pinnedRepo, relative, source string) error {
	parent, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(parent)
	sourceHandle, err := openAbsoluteForMutation(source, false)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(sourceHandle)
	return setName(sourceHandle, parent, leaf, windows.FileLinkInformation, false)
}

func secureRemoveDestination(repo pinnedRepo, relative string, directory bool) error {
	parent, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(parent)
	options := uint32(0)
	if directory {
		options = windows.FILE_DIRECTORY_FILE
	}
	handle, _, err := ntOpenRelative(parent, leaf, windows.DELETE, windows.FILE_OPEN, options)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)
	flags := uint32(windows.FILE_DISPOSITION_DELETE | windows.FILE_DISPOSITION_POSIX_SEMANTICS | windows.FILE_DISPOSITION_IGNORE_READONLY_ATTRIBUTE)
	var status windows.IO_STATUS_BLOCK
	return windows.NtSetInformationFile(handle, &status, (*byte)(unsafe.Pointer(&flags)), uint32(unsafe.Sizeof(flags)), windows.FileDispositionInformationEx)
}

func secureReadDestination(repo pinnedRepo, relative string) (os.FileInfo, []byte, error) {
	return secureReadDestinationWithParentOpened(repo, relative, nil)
}

func secureReadDestinationWithParentOpened(repo pinnedRepo, relative string, parentOpened func()) (os.FileInfo, []byte, error) {
	parent, leaf, err := openDestinationParent(repo, relative, false, nil)
	if err != nil {
		return nil, nil, err
	}
	defer windows.CloseHandle(parent)
	if parentOpened != nil {
		parentOpened()
	}

	handle, _, err := ntOpenRelative(parent, leaf, windows.FILE_GENERIC_READ, windows.FILE_OPEN, windows.FILE_NON_DIRECTORY_FILE)
	if err != nil {
		return nil, nil, fmt.Errorf("open destination file %q: %w", relative, err)
	}
	file := os.NewFile(uintptr(handle), relative)
	if file == nil {
		_ = windows.CloseHandle(handle)
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
