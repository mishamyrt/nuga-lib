package layout

// FromShortKeyCode converts short code to full code
func FromShortKeyCode(code byte) uint32 {
	return uint32(code) << 24
}

// ToShortKeyCode converts full code to short code
func ToShortKeyCode(code uint32) byte {
	return byte(code >> 24)
}
