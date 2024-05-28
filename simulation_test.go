package nuga_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys"
)

func readTemplate(model string) (*device.State, error) {
	path := fmt.Sprintf("device/defaults/%v.nugafile", model)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var template device.State
	err = json.Unmarshal(content, &template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func TestOpenSimulation(t *testing.T) {
	tests := []struct {
		model           string
		expectKeysError bool
	}{
		{"Halo75", false},
		{"Halo65", false},
		{"Halo96", true},
	}
	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			template, err := readTemplate(tt.model)
			if err != nil {
				t.Errorf("Error while reading template: %v", err)
			}
			device, err := nuga.FromTemplate(template)
			if err != nil {
				passError := tt.expectKeysError
				if !passError {
					passError = passError && errors.Is(err, keys.ErrNoTemplate)
				}
				if !passError {
					t.Errorf("Unexpected error on opening simulation: %v", err)
				}
			}
			if device == nil {
				t.Error("Expected non-nil device, got nil")
				return
			}
			if string(device.Name) != tt.model {
				t.Errorf("Unexpected device name '%v'. Expected '%v'", device.Name, tt.model)
			}
		})
	}
}
