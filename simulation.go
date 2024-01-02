package nuga

import (
	"nuga/light"
)

// SimulationTemplate represents the structure of the file that is used for simulation
type SimulationTemplate struct {
	Name     string
	Firmware string
	Lights   *light.SimulationTemplate
}

// NewFeaturesSimulation creates simulated feature repository
func NewFeaturesSimulation(t *SimulationTemplate) *Features {
	return &Features{
		Light: light.NewSimulation(t.Lights),
	}
}
