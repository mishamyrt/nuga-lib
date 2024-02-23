// Package jsonfile provides helpers for working with json file
package jsonfile

import (
	"encoding/json"
	"os"
)

// Read content from json file
func Read[T any](path string) (*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var dump T
	err = json.Unmarshal(data, &dump)
	if err != nil {
		return nil, err
	}
	return &dump, nil
}

// Write content to json file
func Write[T any](path string, dump *T) error {
	data, err := json.Marshal(&dump)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
