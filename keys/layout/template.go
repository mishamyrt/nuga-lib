// Package layout provides helpers for working with keys layout
package layout

// Template represents keyboard layout template
type Template map[KeyName]uint32

// GetPosition returns key position
func (l Template) GetPosition(key KeyName) uint32 {
	return l[key]
}

// GetKey returns key name by position
func (l Template) GetKey(position uint32) KeyName {
	for k, v := range l {
		if v == position {
			return k
		}
	}
	return ""
}

// Halo75 represents Halo 75 keyboard
var Halo75 = Template{
	KeyCapsLock:       3,
	KeyLCtrl:          5,
	KeyLShift:         4,
	KeyLAlt:           11,
	KeyLMeta:          17,
	KeyRCtrl:          65,
	KeyRShift:         82,
	KeyRMeta:          53,
	KeyFn:             59,
	KeyIns:            90,
	KeyHome:           91,
	KeyEnd:            92,
	KeyPgUp:           93,
	KeyPgDn:           94,
	KeyEnter:          81,
	KeyTab:            2,
	KeySpace:          35,
	KeyBackspace:      79,
	KeyEsc:            147,
	KeyDel:            84,
	KeyGrave:          1,
	KeyNum1:           164,
	KeyNum2:           165,
	KeyNum3:           166,
	KeyNum4:           167,
	KeyNum5:           31,
	KeyNum6:           37,
	KeyNum7:           43,
	KeyNum8:           49,
	KeyNum9:           55,
	KeyNum0:           61,
	KeyA:              9,
	KeyB:              169,
	KeyC:              162,
	KeyD:              21,
	KeyE:              20,
	KeyF:              176,
	KeyG:              145,
	KeyH:              168,
	KeyI:              50,
	KeyJ:              45,
	KeyK:              51,
	KeyL:              57,
	KeyM:              46,
	KeyN:              170,
	KeyO:              56,
	KeyP:              62,
	KeyQ:              8,
	KeyR:              26,
	KeyS:              15,
	KeyT:              32,
	KeyU:              44,
	KeyV:              163,
	KeyW:              14,
	KeyX:              161,
	KeyY:              38,
	KeyZ:              10,
	KeySemicolon:      63,
	KeyMinus:          174,
	KeyEqual:          175,
	KeyQuote:          69,
	KeyLBracket:       178,
	KeyRBracket:       179,
	KeyBackslash:      177,
	KeyFwdslash:       173,
	KeyComma:          171,
	KeyPeriod:         172,
	KeyScreenshot:     78,
	KeyFnUp:           197,
	KeyFnDown:         196,
	KeyFnLeft:         198,
	KeyFnRight:        199,
	KeyF1:             132,
	KeyF2:             133,
	KeyF3:             134,
	KeyF4:             135,
	KeyF5:             136,
	KeyF6:             137,
	KeyF7:             138,
	KeyF8:             139,
	KeyF9:             140,
	KeyF10:            141,
	KeyF11:            142,
	KeyF12:            143,
	KeyBrightnessDown: 180,
	KeyBrightnessUp:   181,
	KeyMissionControl: 182,
	KeyLaunchpad:      183,
	KeyVoiceCommand:   184,
	KeySpotlight:      185,
	KeyRewind:         186,
	KeyPlayPause:      187,
	KeyForward:        188,
	KeyMute:           189,
	KeyVolumeDown:     190,
	KeyVolumeUp:       191,
}

// Templates contains all keyboard templates
var Templates = map[string]Template{
	"Halo75": Halo75,
}
