package nuga_test

import (
	"testing"

	"nuga"
)

func TestCapabilityHas(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		cap       nuga.Capability
		feature   nuga.Capability
		expectHas bool
	}{
		{"BacklightCapability", nuga.BacklightCapability, nuga.BacklightCapability, true},
		{"SidelightCapability", nuga.SidelightCapability, nuga.BacklightCapability, false},
		{"HalolightCapability", nuga.HalolightCapability, nuga.HalolightCapability, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.cap.Has(tt.feature)
			if result != tt.expectHas {
				t.Errorf("Expected Has(%v) to be %v, got %v", tt.feature, tt.expectHas, result)
			}
		})
	}
}

func TestGetCapabilities(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		model       string
		expectError bool
	}{
		{"SupportedModel", "Halo96", false},
		{"UnsupportedModel", "NotSupportedModel", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capabilities, err := nuga.GetCapabilities(tt.model)

			if tt.expectError && err == nil {
				t.Errorf("Expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if capabilities != nil && *capabilities == 0 {
				t.Errorf("Expected non-zero capabilities, got zero")
			}
		})
	}
}
