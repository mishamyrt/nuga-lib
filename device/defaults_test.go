package device_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/device"
)

func TestGetDefaults(t *testing.T) {
	// Define test cases
	tests := []struct {
		model          device.Model
		expectingError bool
	}{
		{device.Halo75, false},
		{device.Halo65, false},
		{"WrongModel", true},
	}

	// Run tests
	for _, test := range tests {
		t.Run(string(test.model), func(t *testing.T) {
			state, err := device.GetDefaults(test.model)

			if test.expectingError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}

			if state.Model != test.model {
				t.Errorf("Unexpected device name '%v'. Expected '%v'", state.Model, test.model)
			}

		})
	}
}
