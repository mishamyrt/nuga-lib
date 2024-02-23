package nuga_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mishamyrt/nuga-lib"
)

func readTemplate(model string) (*nuga.SimulationTemplate, error) {
	path := fmt.Sprintf("testdata/simulations/%v.json", strings.ToLower(model))
	content, err := os.ReadFile(path)
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

func TestOpenSimulation(t *testing.T) {
	t.Parallel()
	model := "Halo75"
	template, err := readTemplate(model)
	if err != nil {
		t.Errorf("Error while reading template: %v", err)
	}
	device, err := nuga.OpenSimulation(template)
	if err != nil {
		t.Errorf("Expected error on opening simulation: %v", err)
	}
	if device == nil {
		t.Error("Expected non-nil device, got nil")
		return
	}
	if device.Name != model {
		t.Errorf("Unexpected device name '%v'. Expected '%v'", device.Name, model)
	}
}

func TestNewFeaturesSimulation(t *testing.T) {
	t.Parallel()
	model := "Halo75"
	template, err := readTemplate(model)
	if err != nil {
		t.Errorf("Error while reading template: %v", err)
	}

	features := nuga.NewFeaturesSimulation(template, model)
	if features == nil {
		t.Error("Expected non-nil Features, got nil")
	}
}
