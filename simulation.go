package nuga

import (
	"github.com/mishamyrt/nuga-lib/keys"
	"github.com/mishamyrt/nuga-lib/light"
)

// OpenSimulation opens simulated keyboard
func OpenSimulation(t *SimulationTemplate) (*Device, error) {
	capabilities, err := GetCapabilities(t.Name)
	if err != nil {
		return nil, err
	}
	repo := NewFeaturesSimulation(t, t.Name)
	return &Device{
		Name:         t.Name,
		Path:         "/simulated/device/path",
		Firmware:     t.Firmware,
		Features:     repo,
		Capabilities: capabilities,
	}, nil
}

// SimulationTemplate represents the structure of the file that is used for simulation
type SimulationTemplate struct {
	Name     string
	Firmware string
	Lights   *light.SimulationTemplate
	Keys     *keys.SimulationTemplate
}

// NewFeaturesSimulation creates simulated feature repository
func NewFeaturesSimulation(t *SimulationTemplate, model string) *Features {
	return &Features{
		Light: light.NewSimulation(t.Lights),
		Keys:  keys.NewSimulation(t.Keys, model),
	}
}
