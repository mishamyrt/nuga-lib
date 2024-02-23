// Package nuga is the entrypoint for NuPhy keyboard control library
package nuga

import (
	"errors"

	"github.com/mishamyrt/nuga-lib/hid"
)

// VendorPrefix represents NuPhy HID name prefix
const VendorPrefix = "NuPhy "

// ErrWrongVendor is returned when you try to open a keyboard not from NuPhy
var ErrWrongVendor = errors.New("device vendor is not NuPhy")

// Device represents keyboard with its controls
type Device struct {
	Name         string
	Path         string
	Firmware     string
	Capabilities *Capability
	Features     *Features
	handle       hid.Handler
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
	capabilities, err := GetCapabilities(handle.Info.Model)
	if err != nil {
		return nil, err
	}
	repo := NewFeaturesWith(handle, handle.Info.Model)
	return &Device{
		Name:         handle.Info.Model,
		Path:         handle.Info.Path,
		Firmware:     handle.Info.Firmware,
		Features:     repo,
		Capabilities: capabilities,
		handle:       handle,
	}, nil
}
