package dump

import (
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Restore device state
func Restore(h hid.Handler, s *State) error {
	k := keys.New(h, nil)
	l := light.New(h, nil)
	colors := light.ParseColorsState(s.Lights.Colors)
	effects := light.ParseParamsState(s.Lights.Params)
	err := l.SetBacklightColors(colors)
	if err != nil {
		return err
	}
	macros, err := keys.ParseHeadlessMacros(s.Keys.Macros)
	if err != nil {
		return err
	}
	err = k.SetMacros(macros)
	if err != nil {
		return err
	}
	err = k.SetMacCodes(s.Keys.Mac)
	if err != nil {
		return err
	}
	err = k.SetWinCodes(s.Keys.Win)
	if err != nil {
		return err
	}
	err = l.SetEffects(effects)
	if err != nil {
		return err
	}
	return nil
}
