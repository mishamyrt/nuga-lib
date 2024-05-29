// Package features provides keyboard features repository
package features

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Features represents keyboard features repository.
type Features struct {
	Light LightFeature
	Keys  KeysFeature
}

// LightFeature represents keyboard light feature
type LightFeature interface {
	GetEffects() (*light.Effects, error)
	SetEffects(*light.Effects) error
	GetBacklightColors() (*light.BacklightColors, error)
	SetBacklightColors(*light.BacklightColors) error
	GetCustomEffect() (*light.CustomEffectMap, error)
	SetCustomEffect(*light.CustomEffectMap) error
	GetStateData() (*device.LightsState, error)
	SetStateData(*device.LightsState) error
}

// KeysFeature represents keyboard keys feature
type KeysFeature interface {
	GetWin() (*keys.KeyMap, error)
	SetWin(keyMap *keys.KeyMap) error
	GetMac() (*keys.KeyMap, error)
	SetMac(keyMap *keys.KeyMap) error
	GetMacros() (keys.Macros, error)
	SetMacros(macros keys.Macros) error
	GetStateData() (*device.KeysState, error)
	SetStateData(*device.KeysState) error
}

// New creates Features instance with handle
func New(dev hid.Handler, model device.Model) (*Features, error) {
	k, err := keys.New(dev, model)
	if err != nil {
		return nil, err
	}
	return &Features{
		Light: light.New(dev, model),
		Keys:  k,
	}, nil
}
