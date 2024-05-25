package nuga

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
)

const fakePath = "/simulated/device/path"

// SimulationTemplate represents template for simulated keyboard
type SimulationTemplate struct {
	Name     device.Model        `json:"name"`
	Firmware string              `json:"firmware"`
	Data     *features.StateData `json:"data"`
}

// FromTemplate creates simulated keyboard
func FromTemplate(t *SimulationTemplate) (*Device, error) {
	return &Device{
		Name:         t.Name,
		Path:         fakePath,
		Firmware:     t.Firmware,
		Features:     features.NewSimulation(t.Data, t.Name),
		Capabilities: GetCapabilities(t.Name),
	}, nil
}
