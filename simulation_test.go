package nuga_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/dump"
)

func readTemplate(model string) (*dump.State, error) {
	path := fmt.Sprintf("dump/defaults/%v.nugafile", model)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var template dump.State
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
	device, err := nuga.FromTemplate(template)
	if err != nil {
		t.Errorf("Expected error on opening simulation: %v", err)
	}
	if device == nil {
		t.Error("Expected non-nil device, got nil")
		return
	}
	if string(device.Name) != model {
		t.Errorf("Unexpected device name '%v'. Expected '%v'", device.Name, model)
	}
}
