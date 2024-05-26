package hid

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mishamyrt/nuga-lib/device"
	"github.com/sstallion/go-hid"
)

// VendorPrefix represents NuPhy HID name prefix
const VendorPrefix = "NuPhy "

// ErrWrongVendor is returned when you try to open a keyboard not from NuPhy
var ErrWrongVendor = errors.New("device vendor is not NuPhy")

// DeviceInfo represents NuPhy keyboard HID information
type DeviceInfo struct {
	Model    device.Model
	Firmware string
	Path     string
}

// GetDeviceInfo returns HID device information.
func GetDeviceInfo(d *hid.Device) (*DeviceInfo, error) {
	name, err := d.GetProductStr()
	if err != nil {
		return nil, err
	}
	model, err := TrimVendorPrefix(name)
	if err != nil {
		return nil, err
	}
	info, err := d.GetDeviceInfo()
	if err != nil {
		return nil, err
	}
	return &DeviceInfo{
		Model:    device.Model(model),
		Path:     info.Path,
		Firmware: FormatVersion(info.ReleaseNbr),
	}, nil
}

// TrimVendorPrefix returns model name from vendor prefix
func TrimVendorPrefix(hidName string) (string, error) {
	if !strings.HasPrefix(hidName, VendorPrefix) {
		return "", ErrWrongVendor
	}
	return strings.TrimPrefix(hidName, VendorPrefix), nil
}

// FormatVersion returns string version from hid format
func FormatVersion(v uint16) string {
	symbols := strconv.FormatInt(int64(v), 16)
	version := ""
	for i := range symbols {
		version += string(symbols[i])
		if i != len(symbols)-1 {
			version += "."
		}
	}
	return version
}
