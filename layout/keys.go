package layout

// KeyName represents key name
type KeyName string

// KeyDescription represents key description
type KeyDescription struct {
	Title string `json:"title"`
	Code  uint32 `json:"code"`
	Group string `json:"group"`
}

//revive:disable
const (
	GroupBacklight  = "Backlight"
	GroupMultimedia = "Multimedia"
	GroupSpecial    = "Special"
	GroupSystem     = "System"
	GroupSymbols    = "Symbols"
	GroupModifiers  = "Modifiers"
	GroupNavigation = "Navigation"
	GroupNumPad     = "NumPad"
	GroupFunction   = "Function"
	GroupNumbers    = "Numbers"
	GroupLetters    = "Letters"
)

const (
	KeyNone                    KeyName = "none"
	KeyCapsLock                KeyName = "capslock"
	KeyLCtrl                   KeyName = "lctrl"
	KeyLShift                  KeyName = "lshift"
	KeyLAlt                    KeyName = "lalt"
	KeyLMeta                   KeyName = "lmeta"
	KeyLMetaAlias              KeyName = "lmeta_alias"
	KeyRAlt                    KeyName = "ralt"
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
	KenFNNum1                  KeyName = "fn_num1"
	KenFNNum2                  KeyName = "fn_num2"
	KenFNNum3                  KeyName = "fn_num3"
	KenFNNum4                  KeyName = "fn_num4"
	KenFNNum5                  KeyName = "fn_num5"
	KenFNNum6                  KeyName = "fn_num6"
	KenFNNum7                  KeyName = "fn_num7"
	KenFNNum8                  KeyName = "fn_num8"
	KenFNNum9                  KeyName = "fn_num9"
	KenFNNum0                  KeyName = "fn_num0"
	KeyFNMinus                 KeyName = "fn_minus"
	KeyFNEqual                 KeyName = "fn_equal"
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
	KeyForwardSlash            KeyName = "fwdslash"
	KeyComma                   KeyName = "comma"
	KeyPeriod                  KeyName = "period"
	ScrollLock                 KeyName = "scrolllock"
	NumLock                    KeyName = "numlock"
	NumPadDivide               KeyName = "numpad_div"
	NumPadMultiply             KeyName = "numpad_mul"
	NumPadSubtract             KeyName = "numpad_sub"
	NumPadAdd                  KeyName = "numpad_add"
	NumPadEnter                KeyName = "numpad_enter"
	NumPadDot                  KeyName = "numpad_dot"
	NumPadComma                KeyName = "numpad_comma"
	NumPad1                    KeyName = "numpad1"
	NumPad2                    KeyName = "numpad2"
	NumPad3                    KeyName = "numpad3"
	NumPad4                    KeyName = "numpad4"
	NumPad5                    KeyName = "numpad5"
	NumPad6                    KeyName = "numpad6"
	NumPad7                    KeyName = "numpad7"
	NumPad8                    KeyName = "numpad8"
	NumPad9                    KeyName = "numpad9"
	NumPad0                    KeyName = "numpad0"
	KeyScreenshot              KeyName = "screenshot"
	KeyPrintScreen             KeyName = "print_screen"
	KeyUp                      KeyName = "up"
	KeyDown                    KeyName = "down"
	KeyLeft                    KeyName = "left"
	KeyRight                   KeyName = "right"
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

var Keys = map[KeyName]KeyDescription{
	KeyNone:                    {"None", 0x00000000, GroupSpecial},
	KeyCapsLock:                {"Caps Lock", 0x39000000, GroupSpecial},
	KeyLCtrl:                   {"Left Control", 0xe0000006, GroupModifiers},
	KeyLShift:                  {"Left Shift", 0xe1000006, GroupModifiers},
	KeyLAlt:                    {"Left Alt", 0xe2000006, GroupModifiers},
	KeyLMeta:                   {"Left Meta", 0xe3000006, GroupModifiers},
	KeyLMetaAlias:              {"Left Meta Alias", 0x0c000002, GroupModifiers},
	KeyRAlt:                    {"Right Alt", 0xe6000006, GroupModifiers},
	KeyRCtrl:                   {"Right Control", 0xe4000006, GroupModifiers},
	KeyRShift:                  {"Right Shift", 0xe5000006, GroupModifiers},
	KeyRMeta:                   {"Right Meta", 0xe7000006, GroupModifiers},
	KeyFn:                      {"Function", 0x00000020, GroupSystem},
	KeyIns:                     {"Insert", 0x49000000, GroupNavigation},
	KeyHome:                    {"Home", 0x4a000000, GroupNavigation},
	KeyEnd:                     {"End", 0x4d000000, GroupNavigation},
	KeyPgUp:                    {"Page Up", 0x4b000000, GroupNavigation},
	KeyPgDn:                    {"Page Down", 0x4e000000, GroupNavigation},
	KeyEnter:                   {"Enter", 0x28000000, GroupSpecial},
	KeyTab:                     {"Tab", 0x2b000000, GroupSpecial},
	KeySpace:                   {"Space", 0x2c000000, GroupSpecial},
	KeyBackspace:               {"Backspace", 0x2a000000, GroupSpecial},
	KeyEsc:                     {"Escape", 0x29000000, GroupSpecial},
	KeyDel:                     {"Delete", 0x4c000000, GroupNavigation},
	KeyGrave:                   {"Grave Accent", 0x35000000, GroupSymbols},
	KeyNum1:                    {"Number 1", 0x1e000000, GroupNumbers},
	KeyNum2:                    {"Number 2", 0x1f000000, GroupNumbers},
	KeyNum3:                    {"Number 3", 0x20000000, GroupNumbers},
	KeyNum4:                    {"Number 4", 0x21000000, GroupNumbers},
	KeyNum5:                    {"Number 5", 0x22000000, GroupNumbers},
	KeyNum6:                    {"Number 6", 0x23000000, GroupNumbers},
	KeyNum7:                    {"Number 7", 0x24000000, GroupNumbers},
	KeyNum8:                    {"Number 8", 0x25000000, GroupNumbers},
	KeyNum9:                    {"Number 9", 0x26000000, GroupNumbers},
	KeyNum0:                    {"Number 0", 0x27000000, GroupNumbers},
	KeyA:                       {"A", 0x04000000, GroupLetters},
	KeyB:                       {"B", 0x05000000, GroupLetters},
	KeyC:                       {"C", 0x06000000, GroupLetters},
	KeyD:                       {"D", 0x07000000, GroupLetters},
	KeyE:                       {"E", 0x08000000, GroupLetters},
	KeyF:                       {"F", 0x09000000, GroupLetters},
	KeyG:                       {"G", 0x0a000000, GroupLetters},
	KeyH:                       {"H", 0x0b000000, GroupLetters},
	KeyI:                       {"I", 0x0c000000, GroupLetters},
	KeyJ:                       {"J", 0x0d000000, GroupLetters},
	KeyK:                       {"K", 0x0e000000, GroupLetters},
	KeyL:                       {"L", 0x0f000000, GroupLetters},
	KeyM:                       {"M", 0x10000000, GroupLetters},
	KeyN:                       {"N", 0x11000000, GroupLetters},
	KeyO:                       {"O", 0x12000000, GroupLetters},
	KeyP:                       {"P", 0x13000000, GroupLetters},
	KeyQ:                       {"Q", 0x14000000, GroupLetters},
	KeyR:                       {"R", 0x15000000, GroupLetters},
	KeyS:                       {"S", 0x16000000, GroupLetters},
	KeyT:                       {"T", 0x17000000, GroupLetters},
	KeyU:                       {"U", 0x18000000, GroupLetters},
	KeyV:                       {"V", 0x19000000, GroupLetters},
	KeyW:                       {"W", 0x1a000000, GroupLetters},
	KeyX:                       {"X", 0x1b000000, GroupLetters},
	KeyY:                       {"Y", 0x1c000000, GroupLetters},
	KeyZ:                       {"Z", 0x1d000000, GroupLetters},
	KeySemicolon:               {"Semicolon", 0x33000000, GroupSymbols},
	KeyMinus:                   {"Minus", 0x2d000000, GroupSymbols},
	KeyEqual:                   {"Equal", 0x2e000000, GroupSymbols},
	KeyQuote:                   {"Quote", 0x34000000, GroupSymbols},
	KeyLBracket:                {"Left Bracket", 0x2f000000, GroupSymbols},
	KeyRBracket:                {"Right Bracket", 0x30000000, GroupSymbols},
	KeyBackslash:               {"Backslash", 0x31000000, GroupSymbols},
	KeyForwardSlash:            {"Forward Slash", 0x38000000, GroupSymbols},
	KeyComma:                   {"Comma", 0x36000000, GroupSymbols},
	KeyPeriod:                  {"Period", 0x37000000, GroupSymbols},
	KeyF1:                      {"F1", 0x3a000000, GroupFunction},
	KeyF2:                      {"F2", 0x3b000000, GroupFunction},
	KeyF3:                      {"F3", 0x3c000000, GroupFunction},
	KeyF4:                      {"F4", 0x3d000000, GroupFunction},
	KeyF5:                      {"F5", 0x3e000000, GroupFunction},
	KeyF6:                      {"F6", 0x3f000000, GroupFunction},
	KeyF7:                      {"F7", 0x40000000, GroupFunction},
	KeyF8:                      {"F8", 0x41000000, GroupFunction},
	KeyF9:                      {"F9", 0x42000000, GroupFunction},
	KeyF10:                     {"F10", 0x43000000, GroupFunction},
	KeyF11:                     {"F11", 0x44000000, GroupFunction},
	KeyF12:                     {"F12", 0x45000000, GroupFunction},
	KeyBrightnessDown:          {"Brightness Down", 0x6f000004, GroupSpecial},
	KeyBrightnessUp:            {"Brightness Up", 0x70000004, GroupSpecial},
	KeyRewind:                  {"Rewind", 0xb6000004, GroupMultimedia},
	KeyPlayPause:               {"Play/Pause", 0xcd000004, GroupMultimedia},
	KeyForward:                 {"Forward", 0xb5000004, GroupMultimedia},
	KeyMute:                    {"Mute", 0xe2000004, GroupMultimedia},
	KeyVolumeDown:              {"Volume Down", 0xea000004, GroupMultimedia},
	KeyVolumeUp:                {"Volume Up", 0xe9000004, GroupMultimedia},
	KeyVoiceCommand:            {"Voice Command", 0xcf000004, GroupSpecial},
	KeySpotlight:               {"Spotlight", 0x2000000e, GroupSpecial},
	KeyBacklightEffectPrevious: {"Backlight Effect Previous", 0x0002000b, GroupBacklight},
	KeyBacklightEffectNext:     {"Backlight Effect Next", 0x0003000b, GroupBacklight},
	KeyBacklightBrightnessDown: {"Backlight Brightness Down", 0x0001000c, GroupBacklight},
	KeyBacklightBrightnessUp:   {"Backlight Brightness Up", 0x0002000c, GroupBacklight},
	KeyBacklightColorPrevious:  {"Backlight Color Previous", 0x00020012, GroupBacklight},
	KeyBacklightColorNext:      {"Backlight Color Next", 0x00030012, GroupBacklight},
	KeyScreenshot:              {"Screenshot", 0x18000002, GroupSpecial},
	KeyPrintScreen:             {"Print Screen", 0x46000000, GroupSpecial},
	KeyUp:                      {"Up", 0x52000000, GroupNavigation},
	KeyDown:                    {"Down", 0x51000000, GroupNavigation},
	KeyLeft:                    {"Left", 0x50000000, GroupNavigation},
	KeyRight:                   {"Right", 0x4f000000, GroupNavigation},
	ScrollLock:                 {"Scroll Lock", 0x47000000, GroupSpecial},
	NumLock:                    {"Num Lock", 0x53000000, GroupSpecial},
	NumPadDivide:               {"NumPad Divide", 0x54000000, GroupNumPad},
	NumPadMultiply:             {"NumPad Multiply", 0x55000000, GroupNumPad},
	NumPadSubtract:             {"NumPad Subtract", 0x56000000, GroupNumPad},
	NumPadAdd:                  {"NumPad Add", 0x57000000, GroupNumPad},
	NumPadEnter:                {"NumPad Enter", 0x58000000, GroupNumPad},
	NumPadDot:                  {"NumPad Dot", 0x63000000, GroupNumPad},
	NumPadComma:                {"NumPad Comma", 0x85000000, GroupNumPad},
	NumPad1:                    {"NumPad 1", 0x59000000, GroupNumPad},
	NumPad2:                    {"NumPad 2", 0x5a000000, GroupNumPad},
	NumPad3:                    {"NumPad 3", 0x5b000000, GroupNumPad},
	NumPad4:                    {"NumPad 4", 0x5c000000, GroupNumPad},
	NumPad5:                    {"NumPad 5", 0x5d000000, GroupNumPad},
	NumPad6:                    {"NumPad 6", 0x5e000000, GroupNumPad},
	NumPad7:                    {"NumPad 7", 0x5f000000, GroupNumPad},
	NumPad8:                    {"NumPad 8", 0x60000000, GroupNumPad},
	NumPad9:                    {"NumPad 9", 0x61000000, GroupNumPad},
	NumPad0:                    {"NumPad 0", 0x62000000, GroupNumPad},
}

//revive:enable
