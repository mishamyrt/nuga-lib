package light

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/internal/slices"
	"github.com/mishamyrt/nuga-lib/layout"
)

// State represents raw lights state.
type State struct {
	Colors       []int `json:"colors"`
	Params       []int `json:"params"`
	CustomEffect []int `json:"custom_effect"`
}

// FeatureSimulation represents simulated light feature.
type FeatureSimulation struct {
	effects      *Effects
	colors       *BacklightColors
	template     *layout.BacklightTemplate
	customEffect []byte
}

// ParseColorsState parses raw colors state.
func ParseColorsState(s []int) *BacklightColors {
	colorsSubset := s[ColorsStartOffset : len(s)-ColorEndOffset]
	return ParseBacklightColors(slices.Cast[int, byte](colorsSubset))
}

// ParseParamsState parses raw params state.
func ParseParamsState(s []int) *Effects {
	paramsSubset := s[EffectsOffset : EffectsOffset+ParamsLength]
	return ParseEffects(slices.Cast[int, byte](paramsSubset))
}

// NewSimulation creates simulated light from template.
func NewSimulation(t *State, model *device.Model) *FeatureSimulation {
	var template *layout.BacklightTemplate
	if model != nil {
		template = layout.GetBacklightTemplate(*model)
	}
	return &FeatureSimulation{
		effects:  ParseParamsState(t.Params),
		colors:   ParseColorsState(t.Colors),
		template: template,
	}
}

// GetEffects returns current simulated effect.
func (f *FeatureSimulation) GetEffects() (*Effects, error) {
	return f.effects, nil
}

// SetEffects sets current simulated effect.
func (f *FeatureSimulation) SetEffects(effects *Effects) error {
	f.effects = effects
	return nil
}

// GetBacklightColors returns current simulated colors.
func (f *FeatureSimulation) GetBacklightColors() (*BacklightColors, error) {
	return f.colors, nil
}

// SetBacklightColors sets current simulated colors.
func (f *FeatureSimulation) SetBacklightColors(c *BacklightColors) error {
	f.colors = c
	return nil
}

// GetCustomEffectColors returns current simulated custom effect colors.
func (f *FeatureSimulation) GetCustomEffectColors() (*CustomBacklightMap, error) {
	if f.template == nil {
		return nil, ErrNoCustomColorsTemplate
	}
	return ParseCustomBacklight(f.customEffect, f.template)
}

// SetCustomEffectColors sets current simulated custom effect colors.
func (f *FeatureSimulation) SetCustomEffectColors(colors *CustomBacklightMap) error {
	if f.template == nil {
		return ErrNoCustomColorsTemplate
	}
	f.customEffect = colors.Bytes(f.template)
	return nil
}
