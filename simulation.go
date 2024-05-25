package nuga

import (
	"github.com/mishamyrt/nuga-lib/dump"
	"github.com/mishamyrt/nuga-lib/features"
)

const fakePath = "/simulated/device/path"

// FromTemplate creates simulated keyboard
func FromTemplate(t *dump.State) (*Device, error) {
	return &Device{
		Name:         t.Model,
		Path:         fakePath,
		Firmware:     t.Firmware,
		Features:     features.NewSimulation(&t.Data, t.Model),
		Capabilities: GetCapabilities(t.Model),
	}, nil
}
