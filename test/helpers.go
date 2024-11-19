package test

import (
	"os"
	"path/filepath"
)

func FixturePath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Walk up until we find go.mod
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			return ""
		}
		currentDir = parent
	}

	// Return path to test/fixtures
	return filepath.Join(currentDir, "test", "fixtures")
}

func Fixture(name string) string {
	path := FixturePath()
	if path == "" {
		return ""
	}

	return filepath.Join(path, name)
}
