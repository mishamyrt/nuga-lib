package keys

// PackKeyCodes packs key codes from raw keys payload
func PackKeyCodes(p []byte) []uint32 {
	values := make([]uint32, 256)
	var offset int
	for i := 0; i < 256; i++ {
		for j := 0; j < 4; j++ {
			offset = (i * 4) + (3 - j)
			values[i] = (values[i] << 8) | uint32(p[offset])
		}
	}
	return values
}

// UnpackKeyCodes unpacks key codes to raw keys payload
func UnpackKeyCodes(v []uint32) []byte {
	payload := make([]byte, 1024)
	var offset int
	for i := 0; i < 256; i++ {
		offset = (i * 4)
		payload[offset] = byte(v[i] & 0xFF)
		payload[offset+1] = byte((v[i] >> 8) & 0xFF)
		payload[offset+2] = byte((v[i] >> 16) & 0xFF)
		payload[offset+3] = byte((v[i] >> 24) & 0xFF)
	}
	return payload
}
