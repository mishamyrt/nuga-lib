package dump

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
)

// State represents raw device state. It contains data of all supported features
type State struct {
	Model    device.Model       `json:"model"`
	Firmware string             `json:"firmware"`
	Data     features.StateData `json:"state"`
}
