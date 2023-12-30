// Package features contains controllable keyboard features.
package features

import (
	"nuga/pkg/features/light"
	"nuga/pkg/hid"
)

// Repository represents keyboard features repository.
type Repository struct {
	Light LightFeature
}

// New creates repository instance
func New(handle *hid.Device) *Repository {
	return &Repository{
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
