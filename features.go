package nuga

import (
	"nuga/hid"
	"nuga/light"
)

// Features represents keyboard features repository.
type Features struct {
	Light LightFeature
}

// NewFeaturesWith creates Features instance with handle
func NewFeaturesWith(handle *hid.Device) *Features {
	return &Features{
		Light: light.New(handle),
	}
}

// LightFeature represents keyboard light features
type LightFeature interface {
	GetEffects() (*light.Effects, error)
	SetEffects(p *light.Effects) error
	SetBacklightColors(colors *light.BacklightColors) error
	GetBacklightColors() (*light.BacklightColors, error)
}
