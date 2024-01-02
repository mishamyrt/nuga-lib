package hid_test

import (
	"errors"
	"nuga/hid"
	"testing"
)

func TestTrimVendorPrefixValid(t *testing.T) {
	t.Parallel()
	hidName := "NuPhy Device"
	expected := "Device"

	result, err := hid.TrimVendorPrefix(hidName)

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

	_, err := hid.TrimVendorPrefix(invalidHIDName)

	if !errors.Is(err, hid.ErrWrongVendor) {
		t.Errorf("Expected: %v, Got: %v", hid.ErrWrongVendor, err)
	}
}

func TestFormatVersion(t *testing.T) {
	t.Parallel()
	version := uint16(0x105)
	expected := "1.0.5"

	result := hid.FormatVersion(version)

	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
