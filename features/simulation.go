package features

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/pkg/errors"
)

// StateData represents raw features state. It contains data of all supported features
type StateData struct {
	Lights *light.StateData `json:"lights"`
	Keys   *keys.StateData  `json:"keys"`
}

// NewSimulation creates simulated features from state
func NewSimulation(s *StateData, model device.Model) (*Features, error) {
	l, err := light.NewSimulation(s.Lights, model)
	if err != nil {
		return nil, errors.Wrap(err, "lights")
	}
	k, err := keys.NewSimulation(s.Keys, model)
	if err != nil {
		return nil, errors.Wrap(err, "keys")
	}
	return &Features{
		Light: l,
		Keys:  k,
	}, nil
}
