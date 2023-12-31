// Package device is the entrypoint for keyboard device control
package device

import (
	"nuga/internal/bit"
	"nuga/pkg/features"
	"nuga/pkg/hid"
)

// Device represents keyboard with its controls
type Device struct {
	Name     string
	Path     string
	Firmware string
	Supports bit.Mask
	Features *features.Repository
	handle   *hid.Device
}

// Close connection with hid device
func (d *Device) Close() error {
	if d.handle == nil {
		return nil
	}
	return d.handle.Close()
}

// Open real keyboard USB handle
func Open() (*Device, error) {
	handle, err := hid.Open()
	if err != nil {
		return nil, err
	}
	info, err := handle.GetInfo()
	if err != nil {
		return nil, err
	}
	name, err := TrimVendorPrefix(info.Name)
	if err != nil {
		return nil, err
	}
	supports, err := GetSupports(name)
	if err != nil {
		return nil, err
	}
	repo := features.New(handle)
	return &Device{
		Name:     name,
		Path:     info.Path,
		Firmware: FormatVersion(info.Firmware),
		Features: repo,
		Supports: supports,
		handle:   handle,
	}, nil
}

// OpenSimulation opens simulated keyboard
func OpenSimulation(t *features.SimulationTemplate) (*Device, error) {
	name, err := TrimVendorPrefix(t.Name)
	if err != nil {
		return nil, err
	}
	supports, err := GetSupports(name)
	if err != nil {
		return nil, err
	}
	repo := features.NewSimulation(t)
	return &Device{
		Name:     name,
		Path:     "/simulated/device/path",
		Firmware: t.Firmware,
		Features: repo,
		Supports: supports,
	}, nil
}
