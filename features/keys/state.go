package keys

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/layout"
)

// State represents parsed keys state.
type State struct {
	Mac    *KeyMap `json:"mac"`
	Win    *KeyMap `json:"win"`
	Macros Macros  `json:"macros"`
}

// Data returns raw keys state.
func (s *State) Data(tpl *layout.Template, defaults *device.KeysState) (*device.KeysState, error) {
	macros, err := s.Macros.Bytes()
	if err != nil {
		return nil, err
	}
	mac := s.Mac.Bytes(tpl, defaults.Mac)
	win := s.Win.Bytes(tpl, defaults.Win)
	return &device.KeysState{
		Mac:    mac,
		Win:    win,
		Macros: macros,
	}, nil
}

// ParseState parses raw state data.
func ParseState(s *device.KeysState, tpl *layout.Template) (*State, error) {
	if tpl == nil {
		return nil, ErrNoTemplate
	}
	mac, err := ParseKeyMap(s.Mac, tpl)
	if err != nil {
		return nil, err
	}
	win, err := ParseKeyMap(s.Win, tpl)
	if err != nil {
		return nil, err
	}
	macros, err := ParseMacros(s.Macros)
	if err != nil {
		return nil, err
	}
	return &State{
		Mac:    mac,
		Win:    win,
		Macros: macros,
	}, nil
}
