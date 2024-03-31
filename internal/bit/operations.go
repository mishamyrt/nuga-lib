package bit

func Combine16(high uint8, low uint8) uint16 {
	return (uint16(high) << 8) + uint16(low)
}

func Split16(value uint16) (high uint8, low uint8) {
	return uint8(value >> 8), uint8(value)
}
