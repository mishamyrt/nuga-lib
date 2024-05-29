package nuga

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Collect device state
func Collect(dev hid.Handler, model device.Model) (*device.State, error) {
	f, err := features.New(dev, model)
	if err != nil {
		return nil, err
	}
	lights, err := f.Light.GetStateData()
	if err != nil {
		return nil, err
	}
	keys, err := f.Keys.GetStateData()
	if err != nil {
		return nil, err
	}
	return &device.State{
		Model: model,
		Data: device.StateData{
			Lights: lights,
			Keys:   keys,
		},
	}, nil
}
