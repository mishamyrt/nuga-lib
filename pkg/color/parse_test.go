package color_test

import (
	"errors"
	"nuga/pkg/color"
	"testing"
)

func TestParseHexValidLong(t *testing.T) {
	t.Parallel()
	hexString := "#1a2b3c"
	expectedRGB := color.RGB{R: 0x1a, G: 0x2b, B: 0x3c}
	result, err := color.ParseHex(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *result != expectedRGB {
		t.Errorf("Expected: %v, Got: %v", expectedRGB, result)
	}
}

func TestParseHexValidShort(t *testing.T) {
	t.Parallel()
	hexString := "#f00"
	expectedRGB := color.RGB{R: 0xFF, G: 0, B: 0}
	result, err := color.ParseHex(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *result != expectedRGB {
		t.Errorf("Expected: %v, Got: %v", expectedRGB, result)
	}
}

func TestParseHexInvalid(t *testing.T) {
	t.Parallel()
	value, err := color.ParseHex("invalid_hex")
	if err == nil {
		t.Error("Expected error, but got none")
	}
	if value != nil {
		t.Errorf("Expected nil value, but got %v", value)
	}
	if !errors.Is(err, color.ErrInvalidFormat) {
		t.Errorf("Expected: %v, Got: %v", color.ErrInvalidFormat, err)
	}
}
