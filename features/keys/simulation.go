package keys

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/layout"
)

// State represents simulation data.
type State struct {
	Mac    []uint32 `json:"mac"`
	Win    []uint32 `json:"win"`
	Macros []uint8  `json:"macros,omitempty"`
}

// FeatureSimulation represents simulated keys feature.
type FeatureSimulation struct {
	defaultState State
	data         *State
	template     *layout.KeystrokeTemplate
}

// NewSimulation creates simulated keys from template.
func NewSimulation(t *State, model *device.Model) *FeatureSimulation {
	var template *layout.KeystrokeTemplate
	if model != nil {
		template = layout.GetKeystrokeTemplate(*model)
	}
	if t.Macros == nil {
		t.Macros = make([]uint8, 0)
	}
	return &FeatureSimulation{
		data:         t,
		defaultState: *t,
		template:     template,
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
func (f *FeatureSimulation) GetWin() (*KeyMap, error) {
	return ParseKeyMap(f.data.Win, f.template)
}

// GetMac returns mac keyboard keys
func (f *FeatureSimulation) GetMac() (*KeyMap, error) {
	return ParseKeyMap(f.data.Mac, f.template)
}

// GetMacros returns keyboard macros
func (f *FeatureSimulation) GetMacros() (Macros, error) {
	if len(f.data.Macros) > 0 {
		res := make([]uint8, len(f.data.Macros), 1032)
		res[0] = codeMacroHeader
		copy(res[7:], f.data.Macros)
		return ParseMacros(res)
	}
	return nil, nil
}

// SetMacros sets keyboard macros
func (f *FeatureSimulation) SetMacros(m Macros) error {
	res, err := m.Bytes()
	if err != nil {
		return err
	}
	f.data.Macros = res
	return nil
}

// SetWin sets win keyboard keys
func (f *FeatureSimulation) SetWin(keyMap *KeyMap) error {
	return keyMap.Apply(f.data.Win, f.template)
}

// SetMac sets mac keyboard keys
func (f *FeatureSimulation) SetMac(keyMap *KeyMap) error {
	return keyMap.Apply(f.data.Mac, f.template)
}

// Parse raw keys
func (f *FeatureSimulation) Parse(keys []uint32) (*KeyMap, error) {
	return ParseKeyMap(keys, f.template)
}
