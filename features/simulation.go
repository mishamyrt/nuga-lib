package features

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
)

// StateData represents raw features state. It contains data of all supported features
type StateData struct {
	Lights *light.StateData `json:"lights"`
	Keys   *keys.StateData  `json:"keys"`
}

// NewSimulation creates simulated features from state
func NewSimulation(s *StateData, model device.Model) *Features {
	return &Features{
		Light: light.NewSimulation(s.Lights, model),
		Keys:  keys.NewSimulation(s.Keys, model),
	}
}
