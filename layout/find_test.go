package layout_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/layout"
)

func TestFindKeyName(t *testing.T) {
	tests := []struct {
		code     uint32
		expected layout.KeyName
	}{
		{layout.Keys[layout.KeyA].Code, layout.KeyA},
		{layout.Keys[layout.KeyBacklightBrightnessDown].Code, layout.KeyBacklightBrightnessDown},
		{layout.Keys[layout.KeyNone].Code, layout.KeyNone},
	}

	for _, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			result := layout.FindKeyName(tt.code)
			if result != tt.expected {
				t.Errorf("FindKeyName(%#x): expected %s, got %s", tt.code, tt.expected, result)
			}
		})
	}
}

func TestFindKeyCode(t *testing.T) {
	tests := []struct {
		name     layout.KeyName
		expected uint32
	}{
		{layout.KeyA, layout.Keys[layout.KeyA].Code},
		{layout.KeyB, layout.Keys[layout.KeyB].Code},
		{layout.KeyNone, layout.Keys[layout.KeyNone].Code},
	}

	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			result := layout.FindKeyCode(tt.name)
			if result != tt.expected {
				t.Errorf("FindKeyCode(%s): expected %#x, got %#x", tt.name, tt.expected, result)
			}
		})
	}
}
