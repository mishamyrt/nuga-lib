package hex_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/internal/hex"
)

func TestFromByte(t *testing.T) {
	tests := []struct {
		input  byte
		output string
	}{
		{0x00, "00"},
		{0x01, "01"},
		{0xFF, "ff"},
	}

	for _, tt := range tests {
		t.Run(tt.output, func(t *testing.T) {
			if result := hex.FromByte(tt.input); result != tt.output {
				t.Errorf("Expected FromByte(%#x) to be %v, got %v", tt.input, tt.output, result)
			}
		})
	}
}

func TestFromUInt32(t *testing.T) {
	tests := []struct {
		input  uint32
		output string
	}{
		{0x00000000, "0x00000000"},
		{0x00000001, "0x00000001"},
		{0xffffffff, "0xffffffff"},
	}

	for _, tt := range tests {
		t.Run(tt.output, func(t *testing.T) {
			if result := hex.FromUInt32(tt.input); result != tt.output {
				t.Errorf("Expected FromUInt32(%#x) to be %v, got %v", tt.input, tt.output, result)
			}
		})
	}
}
