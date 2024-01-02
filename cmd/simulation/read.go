package main

import (
	"encoding/json"
	"os"

	"github.com/mishamyrt/nuga-lib"
)

func readTemplate(filePath string) (*nuga.SimulationTemplate, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var template nuga.SimulationTemplate
	err = json.Unmarshal(content, &template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}
