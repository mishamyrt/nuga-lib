// Package device is the entrypoint for keyboard device control
package device

import (
	"nuga/pkg/features"
	"nuga/pkg/hid"
)

// Device represents keyboard with its controls
type Device struct {
	Name     string
	Path     string
	Firmware string
	Features *features.Repository
	handle   *hid.Device
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
	if !IsSupported(name) {
		return nil, ErrNotSupported
	}
	repo := features.New(handle)
	return &Device{
		Name:     name,
		Path:     info.Path,
		Firmware: info.Firmware,
		Features: repo,
		handle:   handle,
	}, nil
}

// Close connection with hid device
func (d *Device) Close() error {
	if d.handle == nil {
		return nil
	}
	return d.handle.Close()
}
