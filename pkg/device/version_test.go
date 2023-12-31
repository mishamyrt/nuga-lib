package device_test

import (
	"nuga/pkg/device"
	"testing"
)

func TestFormatVersion(t *testing.T) {
	t.Parallel()
	version := uint16(0x105)
	expected := "1.0.5"

	result := device.FormatVersion(version)

	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
