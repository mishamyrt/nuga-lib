package dump_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/dump"
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
		// Add more test cases as needed
	}

	// Run tests
	for _, test := range tests {
		t.Run(string(test.model), func(t *testing.T) {
			state, err := dump.GetDefaults(test.model)

			if test.expectingError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}

			if state.Name != test.model {
				t.Errorf("Unexpected device name '%v'. Expected '%v'", state.Name, test.model)
			}

		})
	}
}
