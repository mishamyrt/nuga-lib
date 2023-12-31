package device

import (
	"nuga/internal/bit"
)

const (
	SupportsBacklight bit.Mask = 1 << iota
	SupportsSidelight bit.Mask = 2 << iota
	SupportsHalolight bit.Mask = 4 << iota
)

var haloSeriesSupports = SupportsBacklight | SupportsSidelight | SupportsHalolight

var supportedModels = map[string]bit.Mask{
	"Halo96": haloSeriesSupports,
	"Halo75": haloSeriesSupports,
	"Halo65": haloSeriesSupports,
}

func GetSupports(model string) (bit.Mask, error) {
	if supports, ok := supportedModels[model]; ok {
		return supports, nil
	}
	return 0, ErrNotSupported
}
