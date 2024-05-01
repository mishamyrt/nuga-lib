package keys_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys"
)

func TestParseModifiers(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected keys.Modifiers
	}{
		{"None", 0, keys.Modifiers{}},
		{"Ctrl", keys.ModifierCtrl, keys.Modifiers{Ctrl: true}},
		{"Shift", keys.ModifierShift, keys.Modifiers{Shift: true}},
		{"Alt", keys.ModifierAlt, keys.Modifiers{Alt: true}},
		{"Meta", keys.ModifierMeta, keys.Modifiers{Meta: true}},
		{"CtrlShift", keys.ModifierCtrl | keys.ModifierShift, keys.Modifiers{Ctrl: true, Shift: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keys.ParseModifiers(tt.input)
			if *result != tt.expected {
				t.Errorf("ParseModifiers(%d): expected %v, got %v", tt.input, tt.expected, *result)
			}
		})
	}
}

func TestApplyModifiers(t *testing.T) {
	tests := []struct {
		name     string
		initial  uint32
		modifier keys.Modifiers
		expected uint32
	}{
		{"AddCtrl", 0, keys.Modifiers{Ctrl: true}, keys.ModifierCtrl},
		{
			"AddShiftToCtrl",
			keys.ModifierCtrl,
			keys.Modifiers{Shift: true},
			keys.ModifierCtrl | keys.ModifierShift,
		},
		// Add more cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keys.ApplyModifiers(tt.initial, &tt.modifier)
			if result != tt.expected {
				t.Errorf("%s: expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}

func TestClearModifiers(t *testing.T) {
	tests := []struct {
		name     string
		value    uint32
		expected uint32
	}{
		{"NoModifiers", 0x00000000, 0x00000000},
		{"Modifiers", 0x12345678, 0x12005678},
		{"Ctrl", 0x10010000, 0x10000000},
		{"Shift", 0x10020000, 0x10000000},
		{"Alt", 0x20040000, 0x20000000},
		{"Meta", 0x30080000, 0x30000000},
		{"All", 0x400F0000, 0x40000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := keys.ClearModifiers(tt.value)
			if result != tt.expected {
				t.Errorf("%s: expected %#x, got %#x", tt.name, tt.expected, result)
			}
		})
	}
}
