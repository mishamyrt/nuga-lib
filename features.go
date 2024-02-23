package nuga

import (
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/keys"
	"github.com/mishamyrt/nuga-lib/keys/layout"
	"github.com/mishamyrt/nuga-lib/light"
)

// Features represents keyboard features repository.
type Features struct {
	Light LightFeature
	Keys  KeysFeature
}

// NewFeaturesWith creates Features instance with handle
func NewFeaturesWith(handle *hid.Device, model string) *Features {
	return &Features{
		Light: light.New(handle),
		Keys:  keys.New(handle, model),
	}
}

// LightFeature represents keyboard light feature
type LightFeature interface {
	GetEffects() (*light.Effects, error)
	SetEffects(p *light.Effects) error
	SetBacklightColors(colors *light.BacklightColors) error
	GetBacklightColors() (*light.BacklightColors, error)
}

// KeysFeature represents keyboard keys feature
type KeysFeature interface {
	GetWinCodes() ([]uint32, error)
	GetMacCodes() ([]uint32, error)
	SetWinCodes(keys []uint32) error
	SetMacCodes(keys []uint32) error
	GetWin() (*layout.KeyMap, error)
	GetMac() (*layout.KeyMap, error)
	SetWin(keyMap *layout.KeyMap) error
	SetMac(keyMap *layout.KeyMap) error
}
