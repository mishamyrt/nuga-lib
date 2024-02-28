package dump

import (
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Restore device state
func Restore(h hid.Handler, s *State) error {
	err := restoreLights(h, s.Lights)
	if err != nil {
		return err
	}
	err = restoreKeys(h, s.Keys)
	if err != nil {
		return err
	}
	return nil
}

func restoreLights(h hid.Handler, s light.State) error {
	f := light.New(h)
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

func restoreKeys(h hid.Handler, s keys.State) error {
	f := keys.New(h, nil)
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
