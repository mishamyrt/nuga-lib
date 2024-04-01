package bit

// Combine16 combines two bytes into uint16
func Combine16(high uint8, low uint8) uint16 {
	return (uint16(high) << 8) + uint16(low)
}

// Split16 splits uint16 into two bytes
func Split16(value uint16) (high uint8, low uint8) {
	return uint8(value >> 8), uint8(value)
}
