package layout

import "github.com/mishamyrt/nuga-lib/device"

// BacklightTemplate represents custom backlight template
type BacklightTemplate map[KeyName]uint32

// GetPosition returns key position
func (k BacklightTemplate) GetPosition(key KeyName) uint32 {
	return k[key]
}

// GetKey returns key name by position
func (k BacklightTemplate) GetKey(position uint32) KeyName {
	for k, v := range k {
		if v == position {
			return k
		}
	}
	return ""
}

var halo75Backlight = BacklightTemplate{
	KeyEsc:          0,
	KeyF1:           1,
	KeyF2:           2,
	KeyF3:           3,
	KeyF4:           4,
	KeyF5:           5,
	KeyF6:           6,
	KeyF7:           7,
	KeyF8:           8,
	KeyF9:           9,
	KeyF10:          10,
	KeyF11:          11,
	KeyF12:          12,
	KeyScreenshot:   13,
	KeyDel:          14,
	KeyIns:          15,
	KeyGrave:        21,
	KeyNum1:         22,
	KeyNum2:         23,
	KeyNum3:         24,
	KeyNum4:         25,
	KeyNum5:         26,
	KeyNum6:         27,
	KeyNum7:         28,
	KeyNum8:         29,
	KeyNum9:         30,
	KeyNum0:         31,
	KeyMinus:        32,
	KeyEqual:        33,
	KeyBackspace:    34,
	KeyHome:         36,
	KeyTab:          42,
	KeyQ:            43,
	KeyW:            44,
	KeyE:            45,
	KeyR:            46,
	KeyT:            47,
	KeyY:            48,
	KeyU:            49,
	KeyI:            50,
	KeyO:            51,
	KeyP:            52,
	KeyLBracket:     53,
	KeyRBracket:     54,
	KeyBackslash:    55,
	KeyEnd:          57,
	KeyCapsLock:     63,
	KeyA:            64,
	KeyS:            65,
	KeyD:            66,
	KeyF:            67,
	KeyG:            68,
	KeyH:            69,
	KeyJ:            70,
	KeyK:            71,
	KeyL:            72,
	KeySemicolon:    73,
	KeyQuote:        74,
	KeyEnter:        76,
	KeyPgUp:         78,
	KeyLShift:       84,
	KeyZ:            86,
	KeyX:            87,
	KeyC:            88,
	KeyV:            89,
	KeyB:            90,
	KeyN:            91,
	KeyM:            92,
	KeyComma:        93,
	KeyPeriod:       94,
	KeyForwardSlash: 95,
	KeyRShift:       97,
	KeyUp:           98,
	KeyPgDn:         99,
	KeyLCtrl:        105,
	KeyLAlt:         106,
	KeyLMeta:        107,
	KeySpace:        110,
	KeyRMeta:        113,
	KeyFn:           114,
	KeyLeft:         118,
	KeyDown:         119,
	KeyRight:        120,
}

var backlightTemplates = map[device.Model]BacklightTemplate{
	device.Halo75: halo75Backlight,
}

// GetBacklightTemplate returns backlight template by model
func GetBacklightTemplate(name device.Model) *BacklightTemplate {
	if t, ok := backlightTemplates[name]; ok {
		return &t
	}
	return nil
}
