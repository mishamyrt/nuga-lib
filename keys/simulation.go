package keys

import "github.com/mishamyrt/nuga-lib/keys/layout"

// SimulationTemplate represents simulation data.
type SimulationTemplate struct {
	Mac []uint32 `json:"mac"`
	Win []uint32 `json:"win"`
}

// FeatureSimulation represents simulated keys feature.
type FeatureSimulation struct {
	data     *SimulationTemplate
	template *layout.Template
}

// NewSimulation creates simulated keys from template.
func NewSimulation(t *SimulationTemplate, model string) *FeatureSimulation {
	return &FeatureSimulation{
		data:     t,
		template: layout.GetTemplate(model),
	}
}

func (f *FeatureSimulation) GetWinCodes() ([]uint32, error) {
	return f.data.Win, nil
}

func (f *FeatureSimulation) GetMacCodes() ([]uint32, error) {
	return f.data.Mac, nil
}

func (f *FeatureSimulation) SetWinCodes(keys []uint32) error {
	f.data.Win = keys
	return nil
}

func (f *FeatureSimulation) SetMacCodes(keys []uint32) error {
	f.data.Mac = keys
	return nil
}

func (f *FeatureSimulation) GetWin() (*layout.KeyMap, error) {
	return layout.Parse(f.data.Win, f.template)
}

func (f *FeatureSimulation) GetMac() (*layout.KeyMap, error) {
	return layout.Parse(f.data.Mac, f.template)
}

func (f *FeatureSimulation) SetWin(keyMap *layout.KeyMap) error {
	return keyMap.Apply(f.data.Win, f.template)
}

func (f *FeatureSimulation) SetMac(keyMap *layout.KeyMap) error {
	return keyMap.Apply(f.data.Mac, f.template)
}
