package device

import "strconv"

// FormatVersion returns string version from hid format
func FormatVersion(v uint16) string {
	symbols := strconv.FormatInt(int64(v), 16)
	version := ""
	for i := range symbols {
		version += string(symbols[i])
		if i != len(symbols)-1 {
			version += "."
		}
	}
	return version
}
