package device

import "strings"

// VendorPrefix represents NuPhy HID name prefix
const VendorPrefix = "NuPhy "

// TrimVendorPrefix returns model name from vendor prefix
func TrimVendorPrefix(hidName string) (string, error) {
	if !strings.HasPrefix(hidName, VendorPrefix) {
		return "", ErrWrongVendor
	}
	return strings.TrimPrefix(hidName, VendorPrefix), nil
}
