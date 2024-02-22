package keys

import (
	"github.com/mishamyrt/nuga-lib/hid"
)

// Feature represents keyboard keys feature
type Feature struct {
	handle hid.Handler
}

// New creates keys feature instance.
func New(handle hid.Handler) *Feature {
	return &Feature{
		handle: handle,
	}
}

// GetMac returns mac keyboard keys
func (f *Feature) GetMac() ([]uint32, error) {
	return f.getKeys(cmdGetMacKeys)
}

// GetWin returns win keyboard keys
func (f *Feature) GetWin() ([]uint32, error) {
	return f.getKeys(cmdGetWinKeys)
}

// SetMac sets mac keyboard keys
func (f *Feature) SetMac(keys []uint32) error {
	return f.setKeys(cmdSetMacKeys, keys)
}

// SetWin sets win keyboard keys
func (f *Feature) SetWin(keys []uint32) error {
	return f.setKeys(cmdSetWinKeys, keys)
}

func (f *Feature) setKeys(cmd []byte, keys []uint32) error {
	request := make([]byte, len(cmd)+len(keys)*4)
	copy(request, cmd)
	for i := 0; i < len(keys); i++ {
		offset := len(cmd) + (i * 4)
		unpackBytes(keys[i], request[offset:offset+4])
	}

	return f.handle.Send(request)
}

func (f *Feature) getKeys(cmd []byte) ([]uint32, error) {
	response, err := f.handle.Request(cmd, 1035)
	if err != nil {
		return nil, err
	}
	values := make([]uint32, 256)
	for i := 0; i < 256; i++ {
		offset := 7 + (i * 4)
		values[i] = packBytes(response[offset : offset+4])
	}
	return values, nil
}
