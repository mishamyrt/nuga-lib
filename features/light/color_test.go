package light_test

import (
	"errors"
	"github.com/mishamyrt/nuga-lib/features/light"
	"testing"
)

func TestParseHexValidLong(t *testing.T) {
	t.Parallel()
	hexString := "#1a2b3c"
	expectedRGB := light.RGB{R: 0x1a, G: 0x2b, B: 0x3c}
	result, err := light.ParseHexColor(hexString)
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
	expectedRGB := light.RGB{R: 0xFF, G: 0, B: 0}
	result, err := light.ParseHexColor(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *result != expectedRGB {
		t.Errorf("Expected: %v, Got: %v", expectedRGB, result)
	}
}

func TestParseHexInvalid(t *testing.T) {
	t.Parallel()
	value, err := light.ParseHexColor("invalid_hex")
	expectedErr := light.ErrInvalidColorFormat
	if err == nil {
		t.Error("Expected error, but got none")
	}
	if value != nil {
		t.Errorf("Expected nil value, but got %v", value)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected: %v, Got: %v", expectedErr, err)
	}
}
