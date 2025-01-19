package utils

import (
	"encoding/json"
	"os"
)

const jsonFile = "paths.json"

func LoadPaths() map[string]string {
	paths := make(map[string]string)

	file, err := os.Open(jsonFile)

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
	file, err := os.Create(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(paths)
}
