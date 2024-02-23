package layout

// KeyName represents key name
type KeyName string

//revive:disable
const (
	KeyNone                    KeyName = "none"
	KeyCapsLock                KeyName = "capslock"
	KeyLCtrl                   KeyName = "lctrl"
	KeyLShift                  KeyName = "lshift"
	KeyLAlt                    KeyName = "lalt"
	KeyLMeta                   KeyName = "lmeta"
	KeyRCtrl                   KeyName = "rctrl"
	KeyRShift                  KeyName = "rshift"
	KeyRMeta                   KeyName = "rmeta"
	KeyFn                      KeyName = "fn"
	KeyIns                     KeyName = "ins"
	KeyHome                    KeyName = "home"
	KeyEnd                     KeyName = "end"
	KeyPgUp                    KeyName = "pgup"
	KeyPgDn                    KeyName = "pgdn"
	KeyEnter                   KeyName = "enter"
	KeyTab                     KeyName = "tab"
	KeySpace                   KeyName = "space"
	KeyBackspace               KeyName = "backspace"
	KeyEsc                     KeyName = "esc"
	KeyDel                     KeyName = "del"
	KeyGrave                   KeyName = "grave"
	KeyNum1                    KeyName = "num1"
	KeyNum2                    KeyName = "num2"
	KeyNum3                    KeyName = "num3"
	KeyNum4                    KeyName = "num4"
	KeyNum5                    KeyName = "num5"
	KeyNum6                    KeyName = "num6"
	KeyNum7                    KeyName = "num7"
	KeyNum8                    KeyName = "num8"
	KeyNum9                    KeyName = "num9"
	KeyNum0                    KeyName = "num0"
	KeyA                       KeyName = "a"
	KeyB                       KeyName = "b"
	KeyC                       KeyName = "c"
	KeyD                       KeyName = "d"
	KeyE                       KeyName = "e"
	KeyF                       KeyName = "f"
	KeyG                       KeyName = "g"
	KeyH                       KeyName = "h"
	KeyI                       KeyName = "i"
	KeyJ                       KeyName = "j"
	KeyK                       KeyName = "k"
	KeyL                       KeyName = "l"
	KeyM                       KeyName = "m"
	KeyN                       KeyName = "n"
	KeyO                       KeyName = "o"
	KeyP                       KeyName = "p"
	KeyQ                       KeyName = "q"
	KeyR                       KeyName = "r"
	KeyS                       KeyName = "s"
	KeyT                       KeyName = "t"
	KeyU                       KeyName = "u"
	KeyV                       KeyName = "v"
	KeyW                       KeyName = "w"
	KeyX                       KeyName = "x"
	KeyY                       KeyName = "y"
	KeyZ                       KeyName = "z"
	KeySemicolon               KeyName = "semicolon"
	KeyMinus                   KeyName = "minus"
	KeyEqual                   KeyName = "equal"
	KeyQuote                   KeyName = "quote"
	KeyLBracket                KeyName = "lbracket"
	KeyRBracket                KeyName = "rbracket"
	KeyBackslash               KeyName = "backslash"
	KeyFwdslash                KeyName = "fwdslash"
	KeyComma                   KeyName = "comma"
	KeyPeriod                  KeyName = "period"
	KeyScreenshot              KeyName = "screenshot"
	KeyFnUp                    KeyName = "fn_up"
	KeyFnDown                  KeyName = "fn_down"
	KeyFnLeft                  KeyName = "fn_left"
	KeyFnRight                 KeyName = "fn_right"
	KeyMacSearch               KeyName = "mac_search"
	KeyF1                      KeyName = "f1"
	KeyF2                      KeyName = "f2"
	KeyF3                      KeyName = "f3"
	KeyF4                      KeyName = "f4"
	KeyF5                      KeyName = "f5"
	KeyF6                      KeyName = "f6"
	KeyF7                      KeyName = "f7"
	KeyF8                      KeyName = "f8"
	KeyF9                      KeyName = "f9"
	KeyF10                     KeyName = "f10"
	KeyF11                     KeyName = "f11"
	KeyF12                     KeyName = "f12"
	KeyBrightnessDown          KeyName = "fn_f1"
	KeyBrightnessUp            KeyName = "fn_f2"
	KeyMissionControl          KeyName = "fn_f3"
	KeyLaunchpad               KeyName = "fn_f4"
	KeyVoiceCommand            KeyName = "fn_f5"
	KeySpotlight               KeyName = "fn_f6"
	KeyRewind                  KeyName = "fn_f7"
	KeyPlayPause               KeyName = "fn_f8"
	KeyForward                 KeyName = "fn_f9"
	KeyMute                    KeyName = "fn_f10"
	KeyVolumeDown              KeyName = "fn_f11"
	KeyVolumeUp                KeyName = "fn_f12"
	KeyBacklightEffectNext     KeyName = "backlight_effect_next"
	KeyBacklightEffectPrevious KeyName = "backlight_effect_previous"
	KeyBacklightBrightnessUp   KeyName = "backlight_brightness_up"
	KeyBacklightBrightnessDown KeyName = "backlight_brightness_down"
	KeyBacklightColorNext      KeyName = "backlight_color_next"
	KeyBacklightColorPrevious  KeyName = "backlight_color_previous"
)

var Codes = map[KeyName]uint32{
	KeyCapsLock:                0x39000000,
	KeyLCtrl:                   0xe0000006,
	KeyLShift:                  0xe1000006,
	KeyLAlt:                    0xe2000006,
	KeyLMeta:                   0xe3000006,
	KeyRCtrl:                   0xe4000006,
	KeyRShift:                  0xe5000006,
	KeyRMeta:                   0xe7000006,
	KeyFn:                      0x00000020,
	KeyIns:                     0x49000000,
	KeyHome:                    0x4a000000,
	KeyEnd:                     0x4d000000,
	KeyPgUp:                    0x4b000000,
	KeyPgDn:                    0x4e000000,
	KeyEnter:                   0x28000000,
	KeyTab:                     0x2b000000,
	KeySpace:                   0x2c000000,
	KeyBackspace:               0x2a000000,
	KeyEsc:                     0x29000000,
	KeyDel:                     0x4c000000,
	KeyGrave:                   0x35000000,
	KeyNum1:                    0x1e000000,
	KeyNum2:                    0x1f000000,
	KeyNum3:                    0x20000000,
	KeyNum4:                    0x21000000,
	KeyNum5:                    0x22000000,
	KeyNum6:                    0x23000000,
	KeyNum7:                    0x24000000,
	KeyNum8:                    0x25000000,
	KeyNum9:                    0x26000000,
	KeyNum0:                    0x27000000,
	KeyA:                       0x04000000,
	KeyB:                       0x05000000,
	KeyC:                       0x06000000,
	KeyD:                       0x07000000,
	KeyE:                       0x08000000,
	KeyF:                       0x09000000,
	KeyG:                       0x0a000000,
	KeyH:                       0x0b000000,
	KeyI:                       0x0c000000,
	KeyJ:                       0x0d000000,
	KeyK:                       0x0e000000,
	KeyL:                       0x0f000000,
	KeyM:                       0x10000000,
	KeyN:                       0x11000000,
	KeyO:                       0x12000000,
	KeyP:                       0x13000000,
	KeyQ:                       0x14000000,
	KeyR:                       0x15000000,
	KeyS:                       0x16000000,
	KeyT:                       0x17000000,
	KeyU:                       0x18000000,
	KeyV:                       0x19000000,
	KeyW:                       0x1a000000,
	KeyX:                       0x1b000000,
	KeyY:                       0x1c000000,
	KeyZ:                       0x1d000000,
	KeySemicolon:               0x33000000,
	KeyMinus:                   0x2d000000,
	KeyEqual:                   0x2e000000,
	KeyQuote:                   0x34000000,
	KeyLBracket:                0x2f000000,
	KeyRBracket:                0x30000000,
	KeyBackslash:               0x31000000,
	KeyFwdslash:                0x38000000,
	KeyComma:                   0x36000000,
	KeyPeriod:                  0x37000000,
	KeyF1:                      0x3a000000,
	KeyF2:                      0x3b000000,
	KeyF3:                      0x3c000000,
	KeyF4:                      0x3d000000,
	KeyF5:                      0x3e000000,
	KeyF6:                      0x3f000000,
	KeyF7:                      0x40000000,
	KeyF8:                      0x41000000,
	KeyF9:                      0x42000000,
	KeyF10:                     0x43000000,
	KeyF11:                     0x44000000,
	KeyF12:                     0x45000000,
	KeyBrightnessDown:          0x6f000004,
	KeyBrightnessUp:            0x70000004,
	KeyRewind:                  0xb6000004,
	KeyPlayPause:               0xcd000004,
	KeyForward:                 0xb5000004,
	KeyMute:                    0xe2000004,
	KeyVolumeDown:              0xea000004,
	KeyVolumeUp:                0xe9000004,
	KeyVoiceCommand:            0xcf000004,
	KeySpotlight:               0x2000000e,
	KeyBacklightEffectPrevious: 0x0002000b,
	KeyBacklightEffectNext:     0x0003000b,
	KeyBacklightBrightnessDown: 0x0001000c,
	KeyBacklightBrightnessUp:   0x0002000c,
	KeyBacklightColorPrevious:  0x00020012,
	KeyBacklightColorNext:      0x00030012,
	KeyScreenshot:              0x18000002,
}

//revive:enable
