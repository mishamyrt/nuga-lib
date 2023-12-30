package features

import (
	"nuga/pkg/features/light"
)

type SimulationTemplate struct {
	Name     string
	Firmware string
	Lights   *light.SimulationTemplate
}

func NewSimulation(t *SimulationTemplate) *Repository {
	return &Repository{
		Light: light.NewSimulation(t.Lights),
	}
}
