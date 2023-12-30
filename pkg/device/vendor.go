package device

import (
	"slices"
	"strings"
)

const vendorPrefix = "NuPhy "

var supportedModels = []string{
	"Halo96",
	"Halo75",
	"Halo65",
}

func trimVendorPrefix(hidName string) (string, error) {
	if !strings.HasPrefix(hidName, vendorPrefix) {
		return "", ErrWrongVendor
	}
	return strings.TrimPrefix(hidName, vendorPrefix), nil
}

func isSupported(model string) bool {
	return slices.Contains(supportedModels, model)
}
