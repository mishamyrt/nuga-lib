package layout_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys/layout"
)

func TestParseModifiers(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected layout.Modifiers
	}{
		{"None", 0, layout.Modifiers{}},
		{"Ctrl", layout.ModifierCtrl, layout.Modifiers{Ctrl: true}},
		{"Shift", layout.ModifierShift, layout.Modifiers{Shift: true}},
		{"Alt", layout.ModifierAlt, layout.Modifiers{Alt: true}},
		{"Meta", layout.ModifierMeta, layout.Modifiers{Meta: true}},
		{"CtrlShift", layout.ModifierCtrl | layout.ModifierShift, layout.Modifiers{Ctrl: true, Shift: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := layout.ParseModifiers(tt.input)
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
		modifier layout.Modifiers
		expected uint32
	}{
		{"AddCtrl", 0, layout.Modifiers{Ctrl: true}, layout.ModifierCtrl},
		{
			"AddShiftToCtrl",
			layout.ModifierCtrl,
			layout.Modifiers{Shift: true},
			layout.ModifierCtrl | layout.ModifierShift,
		},
		// Add more cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := layout.ApplyModifiers(tt.initial, &tt.modifier)
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
			result := layout.ClearModifiers(tt.value)
			if result != tt.expected {
				t.Errorf("%s: expected %#x, got %#x", tt.name, tt.expected, result)
			}
		})
	}
}
