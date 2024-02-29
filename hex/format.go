package hex

import "fmt"

func FromByte(v byte) string {
	return fmt.Sprintf("%02x", v)
}

func FromUInt32(v uint32) string {
	return fmt.Sprintf("0x%08x", v)
}
