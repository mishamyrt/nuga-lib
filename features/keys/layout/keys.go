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

func FromShortKeyCode(code byte) uint32 {
	fullCode := uint32(code)
	for i := 0; i < 6; i++ {
		fullCode <<= 4
	}
	return fullCode
}

// IsRegularKey checks if code is regular
func IsRegularKey(code uint32) bool {
	return byte(code&0xFF) == 0x00
}
