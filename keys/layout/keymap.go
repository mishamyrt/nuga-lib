package layout

// Key represents layout key
type Key struct {
	Name      KeyName    `json:"key"`
	Modifiers *Modifiers `json:"modifiers,omitempty"`
}

// KeyMap represents keyboard layout
type KeyMap map[KeyName]Key

// Apply layout to key slice
func (k KeyMap) Apply(source []uint32, tpl *Template) error {
	for keyName, v := range k {
		position := tpl.GetPosition(keyName)
		source[position] = FindKeyCode(v.Name)
		if IsRegularKey(source[position]) {
			source[position] = ApplyModifiers(source[position], v.Modifiers)
		}
	}
	return nil
}

// Parse key map from values
func Parse(values []uint32, tpl *Template) (*KeyMap, error) {
	keys := make(KeyMap)
	for i, position := range *tpl {
		code := values[position]
		var modifiers *Modifiers
		if IsRegularKey(code) {
			modifiers = ParseModifiers(code)
			code = ClearModifiers(code)
		}
		keys[i] = Key{
			Name:      FindKeyName(code),
			Modifiers: modifiers,
		}
	}
	return &keys, nil
}
