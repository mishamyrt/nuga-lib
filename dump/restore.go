package dump

import (
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Restore device state
func Restore(handle *hid.Device, s *State) error {
	err := restoreLights(handle, s.Lights)
	if err != nil {
		return err
	}
	err = restoreKeys(handle, s.Keys)
	if err != nil {
		return err
	}
	return nil
}

func restoreLights(handle *hid.Device, s light.State) error {
	f := light.New(handle)
	colors := light.ParseColorsState(s.Colors)
	effects := light.ParseParamsState(s.Params)
	err := f.SetEffects(effects)
	if err != nil {
		return err
	}
	err = f.SetBacklightColors(colors)
	if err != nil {
		return err
	}
	return nil
}

func restoreKeys(handle *hid.Device, s keys.State) error {
	f := keys.New(handle, nil)
	err := f.SetMacCodes(s.Mac)
	if err != nil {
		return err
	}
	err = f.SetWinCodes(s.Win)
	if err != nil {
		return err
	}
	return nil
}
