package keys

// IsMacroKey checks if code is macros
func IsMacroKey(code uint32) bool {
	return byte(code&0xFF) == 0x10 && byte(code>>4) == 0x01 && byte(code>>16) == 0x01
}

// IsRegularKey checks if code is regular
func IsRegularKey(code uint32) bool {
	return byte(code&0xFF) == 0x00
}

// IndexToMacro converts index to macro
func IndexToMacro(index uint8) uint32 {
	result := uint32(index)
	result <<= 8
	result |= 0x01
	result <<= 8
	result |= 0x00
	result <<= 8
	result |= 0x10
	return result
}

// ExtractMacroIndex extracts index from macro
func ExtractMacroIndex(code uint32) uint8 {
	return uint8(code >> 24)
}
