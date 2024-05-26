package nuga

import (
	"github.com/mishamyrt/nuga-lib/dump"
	"github.com/mishamyrt/nuga-lib/features"
)

const (
	fakePath     = "/simulated/device/path"
	fakeFirmware = "1.0.0"
)

// FromTemplate creates simulated keyboard
func FromTemplate(t *dump.State) (*Device, error) {
	f, err := features.NewSimulation(&t.Data, t.Model)
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
