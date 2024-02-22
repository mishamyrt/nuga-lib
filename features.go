package nuga

import (
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/keys"
	"github.com/mishamyrt/nuga-lib/light"
)

// Features represents keyboard features repository.
type Features struct {
	Light LightFeature
	Keys  KeysFeature
}

// NewFeaturesWith creates Features instance with handle
func NewFeaturesWith(handle *hid.Device) *Features {
	return &Features{
		Light: light.New(handle),
		Keys:  keys.New(handle),
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
	GetWin() ([]uint32, error)
	GetMac() ([]uint32, error)
	SetWin(keys []uint32) error
	SetMac(keys []uint32) error
}
