package keys

// PackKeyCodes packs key codes from raw keys payload
func PackKeyCodes(p []byte) []uint32 {
	count := len(p) / 4
	values := make([]uint32, count)
	var offset int
	for i := 0; i < count; i++ {
		for j := 0; j < 4; j++ {
			offset = (i * 4) + (3 - j)
			values[i] = (values[i] << 8) | uint32(p[offset])
		}
	}
	return values
}

// UnpackKeyCodes unpacks key codes to raw keys payload
func UnpackKeyCodes(v []uint32) []byte {
	payload := make([]byte, len(v)*4)
	var offset int
	for i := 0; i < len(v); i++ {
		offset = (i * 4)
		UnpackKeyCode(v[i], payload[offset:offset+4])
	}
	return payload
}

// UnpackKeyCode unpacks key code to target slice
func UnpackKeyCode(v uint32, target []byte) {
	target[0] = byte(v & 0xFF)
	target[1] = byte((v >> 8) & 0xFF)
	target[2] = byte((v >> 16) & 0xFF)
	target[3] = byte((v >> 24) & 0xFF)
}
