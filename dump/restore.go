package dump

import (
	"github.com/mishamyrt/nuga-lib/features"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Restore device state
func Restore(dev hid.Handler, s *State) error {
	f := features.New(dev, s.Model)
	if err := f.Light.SetStateData(s.Data.Lights); err != nil {
		return err
	}
	return f.Keys.SetStateData(s.Data.Keys)
}
