package keys

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys/layout"
)

// State represents simulation data.
type State struct {
	Mac []uint32 `json:"mac"`
	Win []uint32 `json:"win"`
}

// FeatureSimulation represents simulated keys feature.
type FeatureSimulation struct {
	data     *State
	template *layout.Template
}

// NewSimulation creates simulated keys from template.
func NewSimulation(t *State, model *device.Model) *FeatureSimulation {
	var template *layout.Template
	if model != nil {
		template = layout.GetTemplate(*model)
	}
	return &FeatureSimulation{
		data:     t,
		template: template,
	}
}

// GetWinCodes returns win keyboard key codes
func (f *FeatureSimulation) GetWinCodes() ([]uint32, error) {
	return f.data.Win, nil
}

// GetMacCodes returns mac keyboard key codes
func (f *FeatureSimulation) GetMacCodes() ([]uint32, error) {
	return f.data.Mac, nil
}

// SetWinCodes sets win keyboard key codes
func (f *FeatureSimulation) SetWinCodes(keys []uint32) error {
	f.data.Win = keys
	return nil
}

// SetMacCodes sets mac keyboard key codes
func (f *FeatureSimulation) SetMacCodes(keys []uint32) error {
	f.data.Mac = keys
	return nil
}

// GetWin returns win keyboard keys
func (f *FeatureSimulation) GetWin() (*layout.KeyMap, error) {
	return layout.Parse(f.data.Win, f.template)
}

// GetMac returns mac keyboard keys
func (f *FeatureSimulation) GetMac() (*layout.KeyMap, error) {
	return layout.Parse(f.data.Mac, f.template)
}

// SetWin sets win keyboard keys
func (f *FeatureSimulation) SetWin(keyMap *layout.KeyMap) error {
	return keyMap.Apply(f.data.Win, f.template)
}

// SetMac sets mac keyboard keys
func (f *FeatureSimulation) SetMac(keyMap *layout.KeyMap) error {
	return keyMap.Apply(f.data.Mac, f.template)
}
