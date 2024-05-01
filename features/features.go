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

// New creates Features instance with handle
func New(handle *hid.Device, model device.Model) *Features {
	return &Features{
		Light: light.New(handle, &model),
		Keys:  keys.New(handle, &model),
	}
}

// LightFeature represents keyboard light feature
type LightFeature interface {
	GetEffects() (*light.Effects, error)
	SetEffects(p *light.Effects) error
	GetBacklightColors() (*light.BacklightColors, error)
	SetBacklightColors(colors *light.BacklightColors) error
	GetCustomEffectColors() (*light.CustomBacklightMap, error)
	SetCustomEffectColors(colors *light.CustomBacklightMap) error
}

// KeysFeature represents keyboard keys feature
type KeysFeature interface {
	GetWinCodes() ([]uint32, error)
	GetMacCodes() ([]uint32, error)
	SetWinCodes(keys []uint32) error
	SetMacCodes(keys []uint32) error
	GetWin() (*keys.KeyMap, error)
	GetMac() (*keys.KeyMap, error)
	SetWin(keyMap *keys.KeyMap) error
	SetMac(keyMap *keys.KeyMap) error
	Parse(keys []uint32) (*keys.KeyMap, error)
	GetMacros() (keys.Macros, error)
	SetMacros(macros keys.Macros) error
}
