// Package dump provides device state collector
package dump

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Collect device state
func Collect(dev *hid.Device) (*State, error) {
	model := device.Model(dev.Info.Model)
	f := features.New(dev, model)
	lights, err := f.Light.GetStateData()
	if err != nil {
		return nil, err
	}
	keys, err := f.Keys.GetStateData()
	if err != nil {
		return nil, err
	}
	return &State{
		Model:    model,
		Firmware: dev.Info.Firmware,
		Data: features.StateData{
			Lights: lights,
			Keys:   keys,
		},
	}, nil
}
