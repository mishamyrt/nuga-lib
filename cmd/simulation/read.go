package main

import (
	"encoding/json"
	"nuga/pkg/features"
	"os"
)

func readTemplate(filePath string) (*features.SimulationTemplate, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var template features.SimulationTemplate
	err = json.Unmarshal(content, &template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}
