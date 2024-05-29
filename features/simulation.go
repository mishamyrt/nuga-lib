package features

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/pkg/errors"
)

// NewSimulation creates simulated features from state
func NewSimulation(s *device.State) (*Features, error) {
	l, err := light.NewSimulation(s.Data.Lights, s.Model)
	if err != nil {
		return nil, errors.Wrap(err, "lights")
	}
	k, err := keys.NewSimulation(s.Data.Keys, s.Model)
	if err != nil {
		return nil, errors.Wrap(err, "keys")
	}
	return &Features{
		Light: l,
		Keys:  k,
	}, nil
}
