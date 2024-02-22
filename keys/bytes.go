// Package keys provides helpers for working with keys
package keys

func packBytes(slice []byte) uint32 {
	var combined uint32
	total := len(slice)
	for i := 0; i < total; i++ {
		combined = (combined << 8) | uint32(slice[total-i-1])
	}
	return combined
}

func unpackBytes(value uint32, target []byte) {
	target[0] = byte(value & 0xFF)
	target[1] = byte((value >> 8) & 0xFF)
	target[2] = byte((value >> 16) & 0xFF)
	target[3] = byte((value >> 24) & 0xFF)
}
