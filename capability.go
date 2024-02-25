package nuga

import (
	"github.com/mishamyrt/nuga-lib/device"
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
	// KeysCapability represents that keyboard supports key mapping control
	KeysCapability Capability = 8 << iota
)

const haloSeriesCapabilities = BacklightCapability | SidelightCapability | HalolightCapability

var modelCapabilities = map[device.Model]Capability{
	device.Halo65: haloSeriesCapabilities,
	device.Halo75: haloSeriesCapabilities | KeysCapability,
	device.Halo96: haloSeriesCapabilities,
}

// GetCapabilities returns keyboard model capabilities pointer.
// If keyboard is not supported, returns nil.
func GetCapabilities(model device.Model) *Capability {
	if supports, ok := modelCapabilities[model]; ok {
		return &supports
	}
	return nil
}
