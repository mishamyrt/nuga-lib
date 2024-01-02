package nuga

import (
	"errors"
	"github.com/mishamyrt/nuga-lib/internal/bit"
)

// Capability represents supported keyboard features
type Capability bit.Mask

// Has checks if capability is set
func (c Capability) Has(feature Capability) bool {
	return bit.Mask(c).IsSet(bit.Mask(feature))
}

const (
	// BacklightCapability represents that keyboard supports backlight control
	BacklightCapability Capability = 1 << iota
	// SidelightCapability represents that keyboard supports sidelight control
	SidelightCapability Capability = 2 << iota
	// HalolightCapability represents that keyboard supports halolight control
	HalolightCapability Capability = 4 << iota
)

const haloSeriesCapabilities = BacklightCapability | SidelightCapability | HalolightCapability

var supportedModels = map[string]Capability{
	"Halo96": haloSeriesCapabilities,
	"Halo75": haloSeriesCapabilities,
	"Halo65": haloSeriesCapabilities,
}

// ErrNotSupported is returned when you try to open a keyboard that is not supported by the application
var ErrNotSupported = errors.New("device is not supported")

// GetCapabilities returns keyboard model capabilities.
// If keyboard is not supported, returns ErrNotSupported.
func GetCapabilities(model string) (*Capability, error) {
	if supports, ok := supportedModels[model]; ok {
		return &supports, nil
	}
	return nil, ErrNotSupported
}
