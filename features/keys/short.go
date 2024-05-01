package keys

import "github.com/mishamyrt/nuga-lib/layout"

// FromShortKeyCode converts short code to full code
func FromShortKeyCode(code byte) uint32 {
	return uint32(code) << 24
}

// ToShortKeyCode converts full code to short code
func ToShortKeyCode(code uint32) byte {
	return byte(code >> 24)
}

// FindKeyNameByShortCode finds key name by short code
func FindKeyNameByShortCode(code byte) layout.KeyName {
	for k, v := range layout.Keys {
		if ToShortKeyCode(v.Code) == code {
			return k
		}
	}
	return layout.KeyNone
}

// FindShortKeyCode finds key short code by name
func FindShortKeyCode(name layout.KeyName) byte {
	code := layout.FindKeyCode(name)
	return ToShortKeyCode(code)
}
