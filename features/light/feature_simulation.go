package light

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/layout"
)

// FeatureSimulation represents simulated light feature.
type FeatureSimulation struct {
	state    *State
	template *layout.Template
}

// NewSimulation creates simulated light from template.
func NewSimulation(s *StateData, model device.Model) *FeatureSimulation {
	template := layout.GetBacklightTemplate(model)
	state := s.Parse(template)
	return &FeatureSimulation{
		state:    state,
		template: template,
	}
}

// GetEffects returns current simulated effect.
func (f *FeatureSimulation) GetEffects() (*Effects, error) {
	return f.state.Effects, nil
}

// SetEffects sets current simulated effect.
func (f *FeatureSimulation) SetEffects(e *Effects) error {
	f.state.Effects = e
	return nil
}

// GetBacklightColors returns current simulated colors.
func (f *FeatureSimulation) GetBacklightColors() (*BacklightColors, error) {
	return f.state.Colors, nil
}

// SetBacklightColors sets current simulated colors.
func (f *FeatureSimulation) SetBacklightColors(c *BacklightColors) error {
	f.state.Colors = c
	return nil
}

// GetCustomEffect returns current simulated custom effect colors.
func (f *FeatureSimulation) GetCustomEffect() (*CustomEffectMap, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	return f.state.CustomEffect, nil
}

// SetCustomEffect sets current simulated custom effect colors.
func (f *FeatureSimulation) SetCustomEffect(c *CustomEffectMap) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	f.state.CustomEffect = c
	return nil
}

// GetStateData returns current simulated state.
func (f *FeatureSimulation) GetStateData() (*StateData, error) {
	return f.state.Data(f.template), nil
}

// SetStateData sets current simulated state.
func (f *FeatureSimulation) SetStateData(s *StateData) error {
	f.state = s.Parse(f.template)
	return nil
}
