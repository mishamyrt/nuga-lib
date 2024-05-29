package keys

import (
	"github.com/mishamyrt/nuga-lib/layout"
)

// ActionType represents action type
type ActionType string

const (
	// ActionMacro represents macro action
	ActionMacro ActionType = "macro"
	// ActionKeystroke represents keystroke action
	ActionKeystroke ActionType = "keystroke"
	// ActionNone represents no action
	ActionNone ActionType = "none"
)

// KeystrokeParams represents keystroke params
type KeystrokeParams struct {
	Name      layout.KeyName `json:"key"`
	Modifiers *Modifiers     `json:"modifiers,omitempty"`
}

// Key represents layout key
type Key struct {
	Type       ActionType       `json:"type"`
	Keystroke  *KeystrokeParams `json:"keystroke,omitempty"`
	MacroIndex *uint8           `json:"macro,omitempty"`
}

// KeyMap represents keyboard layout
type KeyMap map[layout.KeyName]Key

// SetKey sets key
func (k KeyMap) SetKey(name layout.KeyName, value Key) error {
	if _, ok := k[name]; !ok {
		return ErrKeyNotFound
	}
	k[name] = value
	return nil
}

// Bytes returns key map as bytes
func (k KeyMap) Bytes(tpl *layout.Template, defaults []byte) []byte {
	raw := make([]byte, len(defaults))
	copy(raw, defaults)
	for keyName, v := range k {
		var code uint32
		position := tpl.GetPosition(keyName)
		switch v.Type {
		case ActionKeystroke:
			code = layout.FindKeyCode(v.Keystroke.Name)
			if IsRegularKey(code) && v.Keystroke.Modifiers != nil {
				code = ApplyModifiers(code, v.Keystroke.Modifiers)
			}
		case ActionMacro:
			code = IndexToMacro(*v.MacroIndex)
		}
		rawPosition := position * 4
		UnpackKeyCode(code, raw[rawPosition:rawPosition+4])
	}
	return raw
}

// ParseKeyMap key map from values
func ParseKeyMap(payload []byte, tpl *layout.Template) (*KeyMap, error) {
	codes := PackKeyCodes(payload)
	keys := make(KeyMap)
	for key, position := range *tpl {
		code := codes[position]
		var actionType ActionType
		var keystroke *KeystrokeParams
		var macroIndex *uint8

		if IsMacroKey(code) {
			index := ExtractMacroIndex(code)
			actionType = ActionMacro
			macroIndex = &index
		} else if IsRegularKey(code) {
			modifiers := ParseModifiers(code)
			code = ClearModifiers(code)
			actionType = ActionKeystroke
			keystroke = &KeystrokeParams{
				Name:      layout.FindKeyName(code),
				Modifiers: modifiers,
			}
		} else {
			name := layout.FindKeyName(code)
			if name != layout.KeyNone {
				actionType = ActionKeystroke
				keystroke = &KeystrokeParams{
					Name:      name,
					Modifiers: nil,
				}
			} else {
				actionType = ActionNone
			}
		}

		keys[key] = Key{
			Type:       actionType,
			Keystroke:  keystroke,
			MacroIndex: macroIndex,
		}
	}
	return &keys, nil
}
