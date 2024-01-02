package light

import "github.com/mishamyrt/nuga-lib/internal/slices"

// SimulationTemplate represents simulation data.
type SimulationTemplate struct {
	Colors []int
	Params []int
}

// FeatureSimulation represents simulated light feature.
type FeatureSimulation struct {
	effects *Effects
	colors  *BacklightColors
}

// NewSimulation creates simulated light from template.
func NewSimulation(t *SimulationTemplate) *FeatureSimulation {
	feature := new(FeatureSimulation)
	setupSimulation(feature, t)
	return feature
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

func setupSimulation(f *FeatureSimulation, t *SimulationTemplate) {
	paramsSubset := t.Params[EffectsOffset : EffectsOffset+ParamsLength]
	colorsSubset := t.Colors[ColorsStartOffset : len(t.Colors)-ColorEndOffset]
	f.effects = ParseEffects(slices.Cast[int, byte](paramsSubset))
	f.colors = ParseBacklightColors(slices.Cast[int, byte](colorsSubset))
}
