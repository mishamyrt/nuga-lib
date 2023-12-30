package device

import (
	"nuga/pkg/features"
)

// OpenSimulation creates simulated features
func OpenSimulation(t *features.SimulationTemplate) (*Device, error) {
	name, err := trimVendorPrefix(t.Name)
	if err != nil {
		return nil, err
	}
	if !isSupported(name) {
		return nil, ErrNotSupported
	}
	repo := features.NewSimulation(t)
	return &Device{
		Name:     name,
		Path:     "/simulated/device/path",
		Firmware: t.Firmware,
		Features: repo,
	}, nil
}
