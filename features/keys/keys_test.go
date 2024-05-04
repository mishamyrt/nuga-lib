package keys_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys"
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

func TestIsRegularKey(t *testing.T) {
	tests := []struct {
		name     string
		code     uint32
		expected bool
	}{
		{"None", 0x0000, false},
		{"Regular large value", 0x9900, true},
		{"Empty not regular", 0x0001, false},
		{"Large not regular", 0x101b, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keys.IsRegularKey(tt.code)
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
			result := keys.ExtractMacroIndex(tt.code)
			if result != tt.expected {
				t.Errorf("ExtractMacroIndex(%#x): expected %d, got %d", tt.code, tt.expected, result)
			}
		})
	}
}

func TestIsMacroKey(t *testing.T) {
	tests := []struct {
		name     string
		code     uint32
		expected bool
	}{
		{"Empty  ", 0x00000000, false},
		{"Regular", 0x01000000, false},
		{"0", 0x00010010, true},
		{"1", 0x01010010, true},
		{"2", 0x02010010, true},
		{"3", 0x03010010, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keys.IsMacroKey(tt.code)
			if result != tt.expected {
				t.Errorf("IsMacroKey(%#x): expected %t, got %t", tt.code, tt.expected, result)
			}
		})
	}
}

func TestIndexToMacro(t *testing.T) {
	tests := []struct {
		name     string
		index    uint8
		expected uint32
	}{
		{"0", 0, 0x00010010},
		{"1", 1, 0x01010010},
		{"2", 2, 0x02010010},
		{"3", 3, 0x03010010},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keys.IndexToMacro(tt.index)
			if result != tt.expected {
				t.Errorf("IndexToMacro(%d): expected %#x, got %#x", tt.index, tt.expected, result)
			}
		})
	}
}
