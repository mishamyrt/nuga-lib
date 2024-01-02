package hid

import (
	"log"
	"sync"
	"time"

	"github.com/sstallion/go-hid"
)

const requestRetries = 5

const (
	readReportID   = 0x06
	nuphyVendorID  = 0x05AC
	nuphyProductID = 0x024F
	nuphyUsagePage = 0xFF00
	nuphyUsage     = 0x01
)

// Device represents NuPhy device handle.
type Device struct {
	Info *DeviceInfo

	handle *hid.Device
	debug  bool
	lock   *sync.Mutex
}

// Debug returns instance of device with enabled verbose logging
func (d *Device) Debug() *Device {
	dev := *d
	dev.debug = true
	return &dev
}

// SendWithRetries sends the request and re-sends it if the request fails.
func (d *Device) SendWithRetries(payload []byte) error {
	var err error
	for i := 0; i < requestRetries; i++ {
		if d.debug {
			log.Println("Send attempt", i+1)
		}
		err = d.Send(payload)
		if err == nil {
			return nil
		}
	}
	if err != nil {
		return err
	}
	return ErrNotFound
}

// Send packet to the device.
func (d *Device) Send(payload []byte) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.debug {
		log.Printf("Send %v bytes: %v", len(payload), payload)
	}

	transferred, err := d.handle.SendFeatureReport(payload)
	if err != nil {
		return err
	}
	expected := len(payload)
	if transferred != len(payload) {
		return NewErrCountMismatch(expected, transferred)
	}

	d.waitSync()
	return nil
}

// Read packet from the device.
func (d *Device) Read(count int) ([]byte, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	buf := make([]byte, count)
	buf[0] = readReportID
	length, err := d.handle.GetFeatureReport(buf)
	if err != nil {
		return nil, err
	}
	packet := buf[1:]
	if d.debug {
		if length > 0 {
			log.Printf("Read %v", buf)
		} else {
			log.Println("Read 0")
		}
	}
	d.waitSync()
	return packet, nil
}

// Request sends a request to the device.
func (d *Device) Request(payload []byte, count int) ([]byte, error) {
	var resp []byte
	var err error
	for i := 0; i < requestRetries; i++ {
		if d.debug {
			log.Println("Read attempt", i+1)
		}
		resp, err = d.tryRequest(payload, count)
		if len(resp) > 0 && resp[0] != 0 {
			return resp, nil
		}
	}
	if err != nil {
		return resp, err
	}
	return resp, ErrNotFound
}

// Close device handle.
// The function should be called after the end of operation with the device.
func (d *Device) Close() error {
	return d.handle.Close()
}

func (d *Device) tryRequest(payload []byte, count int) ([]byte, error) {
	err := d.Send(payload)
	if err != nil {
		return nil, err
	}
	return d.Read(count)
}

func (d *Device) waitSync() {
	time.Sleep(time.Millisecond * 50)
}

// Open connection with the device
func Open() (*Device, error) {
	var path string
	err := hid.Enumerate(nuphyVendorID, nuphyProductID, func(info *hid.DeviceInfo) error {
		if info.Usage == nuphyUsage && info.UsagePage == nuphyUsagePage {
			path = info.Path
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(path) == 0 {
		return nil, ErrNotFound
	}
	handle, err := hid.OpenPath(path)
	if err != nil {
		return nil, err
	}
	info, err := GetDeviceInfo(handle)
	if err != nil {
		return nil, err
	}
	return &Device{
		Info:   info,
		handle: handle,
		lock:   new(sync.Mutex),
	}, nil
}
