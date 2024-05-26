package device_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/device"
)

func TestIsSupported(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   device.Model
		expect bool
	}{
		{"UnknownModel", false},
		{"Halo65", true},
		{"Halo75", true},
		{"Halo96", true},
	}

	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			if device.IsSupported(tt.name) != tt.expect {
				t.Errorf("Expected %v, got %v", tt.expect, device.IsSupported(tt.name))
			}
		})
	}
}
