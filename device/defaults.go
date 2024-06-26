package device

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed all:defaults
var defaults embed.FS

// GetDefaults returns default state for given model
func GetDefaults(model Model) (*State, error) {
	filePath := fmt.Sprintf("defaults/%v.nugafile", string(model))
	data, err := defaults.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var template State
	err = json.Unmarshal(data, &template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}
