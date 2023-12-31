package device

import (
	"nuga/internal/bit"
)

const (
	// SupportsBacklight represents that keyboard supports backlight control
	SupportsBacklight bit.Mask = 1 << iota
	// SupportsSidelight represents that keyboard supports sidelight control
	SupportsSidelight bit.Mask = 2 << iota
	// SupportsHalolight represents that keyboard supports halolight control
	SupportsHalolight bit.Mask = 4 << iota
)

var haloSeriesSupports = SupportsBacklight | SupportsSidelight | SupportsHalolight

var supportedModels = map[string]bit.Mask{
	"Halo96": haloSeriesSupports,
	"Halo75": haloSeriesSupports,
	"Halo65": haloSeriesSupports,
}

// GetSupports returns supports for keyboard model.
// If keyboard is not supported, returns device.ErrNotSupported.
func GetSupports(model string) (bit.Mask, error) {
	if supports, ok := supportedModels[model]; ok {
		return supports, nil
	}
	return 0, ErrNotSupported
}
