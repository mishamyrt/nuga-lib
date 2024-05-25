package light

import "github.com/mishamyrt/nuga-lib/layout"

// State represents parsed lights state.
type State struct {
	Colors       *BacklightColors `json:"colors"`
	Effects      *Effects         `json:"effects"`
	CustomEffect *CustomEffectMap `json:"custom_effect"`
}

// Data returns raw state data.
func (s *State) Data(tpl *layout.Template) *StateData {
	var customEffect []byte
	if tpl != nil {
		customEffect = s.CustomEffect.Bytes(tpl)
	}
	return &StateData{
		Colors:       s.Colors.Bytes(),
		Params:       s.Effects.Bytes(),
		CustomEffect: customEffect,
	}
}

// StateData represents raw lights state.
type StateData struct {
	Colors       []byte `json:"colors"`
	Params       []byte `json:"effects"`
	CustomEffect []byte `json:"custom_effect"`
}

// Parse raw state data.
func (s *StateData) Parse(tpl *layout.Template) *State {
	colors := ParseBacklightColors(s.Colors)
	effects := ParseEffects(s.Params)
	var customEffect *CustomEffectMap
	if tpl != nil {
		customEffect = ParseCustomEffect(s.CustomEffect, tpl)
	}
	return &State{
		Colors:       colors,
		Effects:      effects,
		CustomEffect: customEffect,
	}
}
