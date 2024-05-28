// Package nuga is the entrypoint for NuPhy keyboard control library
package nuga

import (
	"errors"

	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features"
	"github.com/mishamyrt/nuga-lib/hid"
)

// ErrNotSupported is returned when you try to open a keyboard that is not supported by the application
var ErrNotSupported = errors.New("device is not supported")

// Device represents keyboard with its controls
type Device struct {
	Name         device.Model
	Path         string
	Firmware     string
	Capabilities *Capability
	Features     *features.Features
	Handle       hid.Handler
}

// Close connection with hid device
func (d *Device) Close() error {
	if d.Handle == nil {
		return nil
	}
	return d.Handle.Close()
}

// Open real keyboard USB handle
func Open() (*Device, error) {
	handle, err := hid.Open()
	if err != nil {
		return nil, err
	}
	if !device.IsSupported(handle.Info.Model) {
		return nil, ErrNotSupported
	}
	model := device.Model(handle.Info.Model)
	capabilities := GetCapabilities(model)
	repo, err := features.New(handle, model)
	if err != nil {
		return nil, err
	}
	return &Device{
		Name:         model,
		Path:         handle.Info.Path,
		Firmware:     handle.Info.Firmware,
		Features:     repo,
		Capabilities: capabilities,
		Handle:       handle,
	}, nil
}
