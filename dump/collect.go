// Package dump provides device state collector
package dump

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/internal/slices"
)

// Collect device state
func Collect(h hid.Handler) (State, error) {
	lightsState, err := collectLights(h)
	if err != nil {
		return State{}, err
	}
	keysState, err := collectKeys(h)
	if err != nil {
		return State{}, err
	}
	info := h.GetInfo()
	return State{
		Name:     device.Model(info.Model),
		Firmware: info.Firmware,
		Lights:   lightsState,
		Keys:     keysState,
	}, nil
}

func collectLights(h hid.Handler) (light.State, error) {
	var state light.State
	var err error
	lightsFeature := light.New(h)
	params, err := lightsFeature.GetRawEffects()
	if err != nil {
		return state, err
	}
	colors, err := lightsFeature.GetRawColors()
	if err != nil {
		return state, err
	}
	state.Params = slices.Cast[byte, int](params)
	state.Colors = slices.Cast[byte, int](colors)
	return state, nil
}

func collectKeys(h hid.Handler) (keys.State, error) {
	var state keys.State
	var err error
	keysFeature := keys.New(h, nil)
	state.Mac, err = keysFeature.GetMacCodes()
	if err != nil {
		return state, err
	}
	state.Win, err = keysFeature.GetWinCodes()
	if err != nil {
		return state, err
	}
	return state, nil
}
