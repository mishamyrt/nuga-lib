package layout

//revive:disable
const (
	ModifierCtrl  = 0x00010000
	ModifierShift = 0x00020000
	ModifierAlt   = 0x00040000
	ModifierMeta  = 0x00080000
)

//revive:enable

// Modifiers represents keyboard modifiers
type Modifiers struct {
	Ctrl  bool `json:"ctrl"`
	Shift bool `json:"shift"`
	Alt   bool `json:"alt"`
	Meta  bool `json:"meta"`
}

// ParseModifiers parses modifiers from uint32 scan code
func ParseModifiers(v uint32) *Modifiers {
	return &Modifiers{
		Ctrl:  v&ModifierCtrl != 0,
		Shift: v&ModifierShift != 0,
		Alt:   v&ModifierAlt != 0,
		Meta:  v&ModifierMeta != 0,
	}
}

// ApplyModifiers applies modifiers to uint32 scan code
func ApplyModifiers(v uint32, m *Modifiers) uint32 {
	if m.Ctrl {
		v |= ModifierCtrl
	}
	if m.Shift {
		v |= ModifierShift
	}
	if m.Alt {
		v |= ModifierAlt
	}
	if m.Meta {
		v |= ModifierMeta
	}
	return v
}

// ClearModifiers clears modifiers from uint32 scan code
func ClearModifiers(value uint32) uint32 {
	return value & ^(uint32(0xFF << 16))
}
