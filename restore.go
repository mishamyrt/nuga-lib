package nuga

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Restore device state
func Restore(dev hid.Handler, s *device.State) error {
	f, err := features.New(dev, s.Model)
	if err != nil {
		return err
	}
	if err := f.Light.SetStateData(s.Data.Lights); err != nil {
		return err
	}
	return f.Keys.SetStateData(s.Data.Keys)
}
