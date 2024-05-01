// Package hex provides functions for formatting hex strings
package hex

import "fmt"

// FromByte converts byte to hex string
func FromByte(v byte) string {
	return fmt.Sprintf("%02x", v)
}

// FromUInt32 converts uint32 to hex string
func FromUInt32(v uint32) string {
	return fmt.Sprintf("0x%08x", v)
}
