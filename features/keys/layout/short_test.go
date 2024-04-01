package layout_test

import (
	"fmt"
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys/layout"
)

func TestFromShortKeyCode(t *testing.T) {
	tests := []struct {
		code     byte
		expected uint32
	}{
		{0x00, 0x00},
		{0x01, 0x01000000},
		{0x0F, 0x0F000000},
		{0x10, 0x10000000},
		{0x1F, 0x1F000000},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#x", tt.code), func(t *testing.T) {
			result := layout.FromShortKeyCode(tt.code)
			if result != tt.expected {
				t.Errorf("FromShortKeyCode(%#x): expected %#x, got %#x", tt.code, tt.expected, result)
			}
		})
	}
}

func TestToShortKeyCode(t *testing.T) {
	tests := []struct {
		code     uint32
		expected byte
	}{
		{0x00, 0x00},
		{0x01000099, 0x01},
		{0x0F022000, 0x0F},
		{0x10000033, 0x10},
		{0x1F000000, 0x1F},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#x", tt.code), func(t *testing.T) {
			result := layout.ToShortKeyCode(tt.code)
			if result != tt.expected {
				t.Errorf("ToShortKeyCode(%#x): expected %#x, got %#x", tt.code, tt.expected, result)
			}
		})
	}
}
