package layout

// FindKeyName finds key name by code
func FindKeyName(code uint32) KeyName {
	for k, v := range Keys {
		if v.Code == code {
			return k
		}
	}
	return KeyNone
}

// FindKeyCode finds key code by name
func FindKeyCode(name KeyName) uint32 {
	val, ok := Keys[name]
	if !ok {
		return 0
	}
	return val.Code
}

// FindKeyNameByShortCode finds key name by short code
func FindKeyNameByShortCode(code byte) KeyName {
	return FindKeyName(FromShortKeyCode(code))
}

// FindShortKeyCode finds key short code by name
func FindShortKeyCode(name KeyName) byte {
	code := FindKeyCode(name)
	return ToShortKeyCode(code)
}

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
