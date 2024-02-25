package dump

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
)

// State represents device state. It contains data of all supported features
type State struct {
	Name     device.Model `json:"name"`
	Firmware string       `json:"firmware"`
	Lights   light.State  `json:"lights"`
	Keys     keys.State   `json:"keys"`
}
