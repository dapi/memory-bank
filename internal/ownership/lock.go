package ownership

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

var ErrLockNotFound = errors.New("memory-bank lock not found")
var digestPattern = regexp.MustCompile(`^sha256:[0-9a-f]{64}$`)
var modePattern = regexp.MustCompile(`^100(644|755)$`)

func ReadLock(repoRoot string) (Lock, bool, error) {
	repo, err := pinRepoRoot(repoRoot)
	if err != nil {
		return Lock{}, false, err
	}
	lock, exists, _, err := readLockSnapshot(repo)
	return lock, exists, err
}

func readLockSnapshot(repo pinnedRepo) (Lock, bool, string, error) {
	_, destinationInfo, exists, err := inspectDestination(repo, LockFileName)
	if err != nil {
		return Lock{}, false, "", err
	}
	if !exists {
		return Lock{}, false, "", nil
	}
	readInfo, data, err := secureReadDestination(repo, LockFileName)
	if err != nil {
		return Lock{}, false, "", err
	}
	if !os.SameFile(destinationInfo, readInfo) {
		return Lock{}, false, "", fmt.Errorf("%s changed while reading", LockFileName)
	}
	lockDigest := digest(data)
	var lock Lock
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&lock); err != nil {
		return Lock{}, false, "", fmt.Errorf("read %s: %w", LockFileName, err)
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return Lock{}, false, "", fmt.Errorf("read %s: trailing JSON content", LockFileName)
	}
	// Schema 0 was the unversioned prototype. Its fields have v1 semantics and
	// are rewritten as v1 after the next successful update.
	if lock.SchemaVersion != 0 && lock.SchemaVersion != CurrentSchemaVersion {
		return Lock{}, false, "", fmt.Errorf("unsupported memory-bank lock schema %d (supported: %d)", lock.SchemaVersion, CurrentSchemaVersion)
	}
	if lock.Files == nil {
		lock.Files = make(map[string]File)
	}
	if lock.Template.Version == "" || !immutableRefPattern.MatchString(lock.Template.SourceRef) {
		return Lock{}, false, "", fmt.Errorf("invalid template identity in %s", LockFileName)
	}
	if lock.LastUpdate.Version == "" || lock.LastUpdate.At.IsZero() {
		return Lock{}, false, "", fmt.Errorf("invalid last update in %s", LockFileName)
	}
	for filePath, file := range lock.Files {
		if filePath == LockFileName || strings.Contains(filePath, "\\") || path.Clean(filePath) != filePath || !strings.HasPrefix(filePath, "memory-bank/") {
			return Lock{}, false, "", fmt.Errorf("invalid path %q in %s", filePath, LockFileName)
		}
		switch file.Ownership {
		case Managed, Generated:
			if !digestPattern.MatchString(file.BaseDigest) || !digestPattern.MatchString(file.PayloadDigest) {
				return Lock{}, false, "", fmt.Errorf("invalid digest contract for %s", filePath)
			}
			if lock.SchemaVersion == CurrentSchemaVersion && (!modePattern.MatchString(file.BaseMode) || !modePattern.MatchString(file.PayloadMode)) {
				return Lock{}, false, "", fmt.Errorf("invalid mode contract for %s", filePath)
			}
		case Adapted:
			if !digestPattern.MatchString(file.BaseDigest) {
				return Lock{}, false, "", fmt.Errorf("invalid base digest for %s", filePath)
			}
			if lock.SchemaVersion == CurrentSchemaVersion && !modePattern.MatchString(file.BaseMode) {
				return Lock{}, false, "", fmt.Errorf("invalid base mode for %s", filePath)
			}
		case UserOwned:
		default:
			return Lock{}, false, "", fmt.Errorf("invalid ownership %q for %s", file.Ownership, filePath)
		}
	}
	return lock, true, lockDigest, nil
}

func marshalLock(lock Lock) ([]byte, error) {
	data, err := json.MarshalIndent(lock, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}
