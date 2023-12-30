package device

import (
	"errors"
	"slices"
	"strings"
)

// ErrNotSupported is returned when you try to open a keyboard that is not supported by the application
var ErrNotSupported = errors.New("device is not supported")

// ErrWrongVendor is returned when you try to open a keyboard not from NuPhy
var ErrWrongVendor = errors.New("device vendor is not NuPhy")

const vendorPrefix = "NuPhy "

var supportedModels = []string{
	"Halo96",
	"Halo75",
	"Halo65",
}

func TrimVendorPrefix(hidName string) (string, error) {
	if !strings.HasPrefix(hidName, vendorPrefix) {
		return "", ErrWrongVendor
	}
	return strings.TrimPrefix(hidName, vendorPrefix), nil
}

func IsSupported(model string) bool {
	return slices.Contains(supportedModels, model)
}
