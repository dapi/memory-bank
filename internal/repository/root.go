// Package repository contains repository-level discovery shared by CLI commands.
package repository

import (
	"os"
	"path/filepath"
)

// ResolveRoot returns an explicit absolute root or discovers the nearest Git
// repository from the current directory. Outside Git it uses the current directory.
func ResolveRoot(explicitRoot string) (string, error) {
	if explicitRoot != "" {
		return filepath.Abs(explicitRoot)
	}

	currentDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if gitRoot, ok := findNearestGitRoot(currentDirectory); ok {
		return gitRoot, nil
	}
	return filepath.Abs(currentDirectory)
}

func findNearestGitRoot(startDirectory string) (string, bool) {
	currentDirectory, err := filepath.Abs(startDirectory)
	if err != nil {
		return "", false
	}
	for {
		if _, statErr := os.Stat(filepath.Join(currentDirectory, ".git")); statErr == nil {
			return currentDirectory, true
		}
		parent := filepath.Dir(currentDirectory)
		if parent == currentDirectory {
			return "", false
		}
		currentDirectory = parent
	}
}
