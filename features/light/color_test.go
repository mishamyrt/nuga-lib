package light_test

import (
	"errors"
	"testing"

	"github.com/mishamyrt/nuga-lib/features/light"
)

func TestParseHexValidLong(t *testing.T) {
	t.Parallel()
	hexString := "#1A2B3C"
	var expected light.Color = 0x1A2B3C
	result, err := light.ParseHexColor(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestParseHexValidShort(t *testing.T) {
	t.Parallel()
	hexString := "#f00"
	var expected light.Color = 0xFF0000
	result, err := light.ParseHexColor(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestParseHexInvalid(t *testing.T) {
	t.Parallel()
	value, err := light.ParseHexColor("invalid_hex")
	expectedErr := light.ErrInvalidColorFormat
	if err == nil {
		t.Error("Expected error, but got none")
	}
	if value != 0 {
		t.Errorf("Expected 0 value, but got %v", value)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected: %v, Got: %v", expectedErr, err)
	}
}
