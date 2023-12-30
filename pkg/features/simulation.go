package features

import (
	"nuga/pkg/features/light"
)

// SimulationTemplate represents the structure of the file that is used for simulation
type SimulationTemplate struct {
	Name     string
	Firmware string
	Lights   *light.SimulationTemplate
}

// NewSimulation creates simulated feature repository
func NewSimulation(t *SimulationTemplate) *Repository {
	return &Repository{
		Light: light.NewSimulation(t.Lights),
	}
}
