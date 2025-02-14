package helper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ResolvePath(paths string, args []string, loadedPaths map[string]string) (string, error) {
	if paths != "" {
		parts := strings.Split(paths, "\\")
		basePath, exists := loadedPaths[parts[0]]

		if !exists {
			return "", fmt.Errorf("no path found for name %s", parts[0])
		}

		if len(parts) > 1 {
			return filepath.Join(basePath, filepath.Join(parts[1:]...)), nil
		}
		return basePath, nil
	}

	if len(args) > 0 {
		path := args[0]
		if path == "." {
			currentDir, err := os.Getwd()
			if err != nil {
				return "", err
			}

			return currentDir, nil
		}

		return path, nil
	}
	return "", errors.New("either provide a path as an argument or use the -p flag")
}

func GetPathCompletions(paths string, loadedPaths map[string]string) []string {
	var completions []string
	parts := strings.Split(paths, "\\")

	if len(parts) == 1 {
		for key := range loadedPaths {
			if strings.HasPrefix(key, paths) {
				completions = append(completions, key)
			}
		}
		return completions
	}

	basePath, exists := loadedPaths[parts[0]]
	if !exists {
		return completions
	}

	searchPath := filepath.Join(basePath, filepath.Join(parts[1:]...))
	parentDir := filepath.Dir(searchPath)

	entries, err := os.ReadDir(parentDir)
	if err != nil {
		return completions
	}

	prefix := parts[len(parts)-1]
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), prefix) {
			suggestion := fmt.Sprintf("%s\\%s", parts[0], filepath.Join(parts[1:len(parts)-1]...))
			if suggestion != "" {
				suggestion += "\\"
			}
			suggestion += entry.Name()
			completions = append(completions, suggestion)
		}
	}

	return completions
}
