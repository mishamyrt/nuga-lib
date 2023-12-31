package device

import "strings"

const VendorPrefix = "NuPhy "

func TrimVendorPrefix(hidName string) (string, error) {
	if !strings.HasPrefix(hidName, VendorPrefix) {
		return "", ErrWrongVendor
	}
	return strings.TrimPrefix(hidName, VendorPrefix), nil
}
