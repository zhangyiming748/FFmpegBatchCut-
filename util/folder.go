package util

import (
	"os"
	"path/filepath"
)

func GetFoldersWithTimestamps(dir string) ([]string, error) {
	var folders []string

	// Walk the directory tree
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current path is a directory
		if info.IsDir() {
			// Check for the existence of timestamps.txt in the current directory
			timestampFile := filepath.Join(path, "timestamps.txt")
			if _, err := os.Stat(timestampFile); err == nil {
				// If the file exists, add the directory to the slice
				folders = append(folders, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return folders, nil
}
