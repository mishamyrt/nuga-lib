package nuga

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
)

const (
	fakePath     = "/simulated/device/path"
	fakeFirmware = "1.0.0"
)

// FromTemplate creates simulated keyboard
func FromTemplate(t *device.State) (*Device, error) {
	f, err := features.NewSimulation(t)
	if err != nil {
		return nil, err
	}
	return &Device{
		Name:         t.Model,
		Path:         fakePath,
		Firmware:     fakeFirmware,
		Features:     f,
		Capabilities: GetCapabilities(t.Model),
	}, nil
}
