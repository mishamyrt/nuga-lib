package layout

// FindKeyName finds key name by code
func FindKeyName(code uint32) KeyName {
	for k, v := range Codes {
		if v == code {
			return k
		}
	}
	return KeyNone
}

// FindKeyCode finds key code by name
func FindKeyCode(name KeyName) uint32 {
	val, ok := Codes[name]
	if !ok {
		return 0
	}
	return val
}

// IsRegularKey checks if code is regular
func IsRegularKey(code uint32) bool {
	return byte(code&0xFF) == 0x00
}
