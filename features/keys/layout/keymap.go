package layout

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
	Name      KeyName    `json:"key"`
	Modifiers *Modifiers `json:"modifiers,omitempty"`
}

// Key represents layout key
type Key struct {
	Type       ActionType       `json:"type"`
	Keystroke  *KeystrokeParams `json:"keystroke,omitempty"`
	MacroIndex *uint8           `json:"macro,omitempty"`
}

// KeyMap represents keyboard layout
type KeyMap map[KeyName]Key

// Apply layout to key slice
func (k KeyMap) Apply(source []uint32, tpl *Template) error {
	for keyName, v := range k {
		position := tpl.GetPosition(keyName)
		switch v.Type {
		case ActionKeystroke:
			source[position] = FindKeyCode(v.Keystroke.Name)
			if IsRegularKey(source[position]) && v.Keystroke.Modifiers != nil {
				source[position] = ApplyModifiers(source[position], v.Keystroke.Modifiers)
			}
		case ActionMacro:
			source[position] = IndexToMacro(*v.MacroIndex)
		case ActionNone:
			source[position] = 0
		}

	}
	return nil
}

// Parse key map from values
func Parse(values []uint32, tpl *Template) (*KeyMap, error) {
	keys := make(KeyMap)
	for key, position := range *tpl {
		code := values[position]
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
				Name:      FindKeyName(code),
				Modifiers: modifiers,
			}
		} else {
			name := FindKeyName(code)
			if name != KeyNone {
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
