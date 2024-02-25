package dump

import (
	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/features"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
)

// OpenSimulation opens simulated keyboard
func OpenSimulation(s *State) (*nuga.Device, error) {
	capabilities := nuga.GetCapabilities(s.Name)
	repo := simulateFeatures(s)
	return &nuga.Device{
		Name:         s.Name,
		Path:         "/simulated/device/path",
		Firmware:     s.Firmware,
		Features:     repo,
		Capabilities: capabilities,
	}, nil
}

func simulateFeatures(s *State) *features.Features {
	return &features.Features{
		Light: light.NewSimulation(&s.Lights),
		Keys:  keys.NewSimulation(&s.Keys, &s.Name),
	}
}
