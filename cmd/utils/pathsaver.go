package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const jsonFileName = "paths.json"

func getGlobalFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}
	globalDir := filepath.Join(homeDir, ".gos")
	if err := os.MkdirAll(globalDir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}
	return filepath.Join(globalDir, jsonFileName)
}

func LoadPaths() map[string]string {
	paths := make(map[string]string)
	globalFilePath := getGlobalFilePath()
	file, err := os.Open(globalFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return paths
		}
		panic(err)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&paths); err != nil {
		panic(err)
	}
	return paths
}

func SavedPaths(paths map[string]string) error {
	globalFilePath := getGlobalFilePath()
	file, err := os.Create(globalFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(paths)
}
