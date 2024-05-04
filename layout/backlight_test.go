package layout_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/layout"
)

func TestGetBacklightTemplate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      device.Model
		expectNil bool
	}{
		{"Halo75", false},
		{"WrongModel", true},
	}
	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			result := layout.GetBacklightTemplate(tt.name)
			if tt.expectNil && result != nil {
				t.Errorf("Expected nil, got pointer")
			} else if !tt.expectNil && result == nil {
				t.Errorf("Unexpected nil")
			}
		})
	}
}
