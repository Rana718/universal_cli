package helper

import (
	"errors"
	"fmt"
	"os"
)

func ResolvePath(paths string, args []string, loadedPaths map[string]string) (string, error) {
	if paths != "" {
		path, exists := loadedPaths[paths]

		if !exists {
			return "", fmt.Errorf("Error: No path found for name %s", paths)
		}
		return path, nil
	}

	if len(args) > 0{
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