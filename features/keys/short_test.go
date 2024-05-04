package keys_test

import (
	"fmt"
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/layout"
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
			result := keys.FromShortKeyCode(tt.code)
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
			result := keys.ToShortKeyCode(tt.code)
			if result != tt.expected {
				t.Errorf("ToShortKeyCode(%#x): expected %#x, got %#x", tt.code, tt.expected, result)
			}
		})
	}
}

func TestFindKeyNameByShortCode(t *testing.T) {
	tests := []struct {
		code     byte
		expected layout.KeyName
	}{
		{0x10, layout.KeyM},
		{0x1C, layout.KeyY},
		{0x15, layout.KeyR},
		{0x17, layout.KeyT},
		{0x99, layout.KeyNone},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#x", tt.code), func(t *testing.T) {
			result := keys.FindKeyNameByShortCode(tt.code)
			if result != tt.expected {
				t.Errorf("FindKeyNameByShortCode(%#x): expected %s, got %s", tt.code, tt.expected, result)
			}
		})
	}
}

func TestFindShortKeyCode(t *testing.T) {
	tests := []struct {
		name     layout.KeyName
		expected byte
	}{
		{layout.KeyM, 0x10},
		{layout.KeyY, 0x1C},
		{layout.KeyR, 0x15},
		{layout.KeyT, 0x17},
	}
	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			result := keys.FindShortKeyCode(tt.name)
			if result != tt.expected {
				t.Errorf("FindShortKeyCode(%s): expected %#x, got %#x", tt.name, tt.expected, result)
			}
		})
	}
}
