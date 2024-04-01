package layout

// FromShortKeyCode converts short code to full code
func FromShortKeyCode(code byte) uint32 {
	fullCode := uint32(code)
	for i := 0; i < 6; i++ {
		fullCode <<= 4
	}
	return fullCode
}

// ToShortKeyCode converts full code to short code
func ToShortKeyCode(code uint32) byte {
	fullCode := uint32(code)
	for i := 0; i < 6; i++ {
		fullCode >>= 4
	}
	return byte(fullCode)
}
