package features

import (
	"nuga/pkg/features/light"
	"nuga/pkg/hid"
)

type Repository struct {
	Light LightFeature
}

func New(handle *hid.Device) *Repository {
	return &Repository{
		Light: light.New(handle),
	}
}

type LightFeature interface {
	GetEffects() (*light.Effects, error)
	SetEffects(p *light.Effects) error
	SetBacklightColors(colors *light.BacklightColors) error
	GetBacklightColors() (*light.BacklightColors, error)
}
