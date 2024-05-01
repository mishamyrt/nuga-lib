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
