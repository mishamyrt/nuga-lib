package dump

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Restore device state
func Restore(dev *hid.Device, s *State) error {
	model := device.Model(dev.Info.Model)
	f := features.New(dev, model)
	if err := f.Light.SetStateData(s.Data.Lights); err != nil {
		return err
	}
	return f.Keys.SetStateData(s.Data.Keys)
}
