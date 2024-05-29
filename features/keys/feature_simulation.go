package keys

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/layout"
)

// FeatureSimulation represents simulated keys feature.
type FeatureSimulation struct {
	state    *State
	defaults *device.KeysState
	template *layout.Template
}

// NewSimulation creates simulated keys from template.
func NewSimulation(s *device.KeysState, model device.Model) (*FeatureSimulation, error) {
	var (
		f   FeatureSimulation
		err error
	)

	f.defaults = s
	f.template = layout.GetKeystrokeTemplate(model)
	if f.template != nil {
		f.state, err = ParseState(s, f.template)
	}
	if err != nil {
		return nil, err
	}
	return &f, nil
}

// GetMac returns mac keyboard keys
func (f *FeatureSimulation) GetMac() (*KeyMap, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	return f.state.Mac, nil
}

// SetMac sets mac keyboard keys
func (f *FeatureSimulation) SetMac(m *KeyMap) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	f.state.Mac = m
	return nil
}

// GetWin returns win keyboard keys
func (f *FeatureSimulation) GetWin() (*KeyMap, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	return f.state.Win, nil
}

// SetWin sets win keyboard keys
func (f *FeatureSimulation) SetWin(m *KeyMap) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	f.state.Win = m
	return nil
}

// GetMacros returns keyboard macros
func (f *FeatureSimulation) GetMacros() (Macros, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	return f.state.Macros, nil
}

// SetMacros sets keyboard macros
func (f *FeatureSimulation) SetMacros(m Macros) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	f.state.Macros = m
	return nil
}

// GetStateData returns current simulated state.
func (f *FeatureSimulation) GetStateData() (*device.KeysState, error) {
	if f.template == nil {
		return &device.KeysState{
			Mac:    make([]byte, 1024),
			Win:    make([]byte, 1024),
			Macros: make([]byte, 1024),
		}, nil
	}
	return f.state.Data(f.template, f.defaults)
}

// SetStateData sets current simulated state.
func (f *FeatureSimulation) SetStateData(s *device.KeysState) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	state, err := ParseState(s, f.template)
	if err != nil {
		return err
	}
	f.state = state
	return nil
}
