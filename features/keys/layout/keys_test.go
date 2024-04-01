package layout_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys/layout"
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

func TestIsRegularKey(t *testing.T) {
	tests := []struct {
		name     string
		code     uint32
		expected bool
	}{
		{"Empty", 0x0000, true},
		{"Regular large value", 0x9900, true},
		{"Empty not regular", 0x0001, false},
		{"Large not regular", 0x101b, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := layout.IsRegularKey(tt.code)
			if result != tt.expected {
				t.Errorf("IsRegularKey(%#x): expected %t, got %t", tt.code, tt.expected, result)
			}
		})
	}
}

func TestExtractMacroIndex(t *testing.T) {
	tests := []struct {
		name     string
		code     uint32
		expected uint8
	}{
		{"Empty", 0x00010010, 0},
		{"1", 0x01010010, 1},
		{"2", 0x02010010, 2},
		{"3", 0x03010010, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := layout.ExtractMacroIndex(tt.code)
			if result != tt.expected {
				t.Errorf("ExtractMacroIndex(%#x): expected %d, got %d", tt.code, tt.expected, result)
			}
		})
	}
}
