package light

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/layout"
	"github.com/pkg/errors"
)

// State represents parsed lights state.
type State struct {
	Colors       *BacklightColors `json:"colors"`
	Effects      *Effects         `json:"effects"`
	CustomEffect *CustomEffectMap `json:"custom_effect"`
}

// Data returns raw state data.
func (s *State) Data(tpl *layout.Template) *device.LightsState {
	var customEffect []byte
	if tpl != nil {
		custom, err := s.CustomEffect.Bytes(tpl)
		if err == nil {
			customEffect = custom
		}
	}
	return &device.LightsState{
		Colors:       s.Colors.Bytes(),
		Params:       s.Effects.Bytes(),
		CustomEffect: customEffect,
	}
}

// Parse raw state data.
func ParseState(s *device.LightsState, tpl *layout.Template) (*State, error) {
	colors, err := ParseBacklightColors(s.Colors)
	if err != nil {
		return nil, errors.Wrap(err, "colors")
	}
	effects, err := ParseEffects(s.Params)
	if err != nil {
		return nil, errors.Wrap(err, "effects")
	}
	var customEffect *CustomEffectMap
	if tpl != nil {
		custom, err := ParseCustomEffect(s.CustomEffect, tpl)
		if err != nil {
			return nil, errors.Wrap(err, "custom effect")
		}
		customEffect = custom
	}
	return &State{
		Colors:       colors,
		Effects:      effects,
		CustomEffect: customEffect,
	}, nil
}
