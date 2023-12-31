package device_test

import (
	"errors"
	"nuga/pkg/device"
	"testing"
)

func TestTrimVendorPrefixValid(t *testing.T) {
	t.Parallel()
	hidName := "NuPhy Device"
	expected := "Device"

	result, err := device.TrimVendorPrefix(hidName)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestTrimVendorPrefixInvalid(t *testing.T) {
	t.Parallel()
	invalidHIDName := "OtherVendor Device"

	_, err := device.TrimVendorPrefix(invalidHIDName)

	if !errors.Is(err, device.ErrWrongVendor) {
		t.Errorf("Expected: %v, Got: %v", device.ErrWrongVendor, err)
	}
}
