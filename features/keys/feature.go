package keys

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/features/keys/layout"
	"github.com/mishamyrt/nuga-lib/hid"
)

// Feature represents keyboard keys feature
type Feature struct {
	handle   hid.Handler
	template *layout.Template
}

// New creates keys feature instance.
func New(handle hid.Handler, model *device.Model) *Feature {
	var template *layout.Template
	if model != nil {
		template = layout.GetTemplate(*model)
	}
	return &Feature{
		handle:   handle,
		template: template,
	}
}

// GetWin returns win keyboard keys
func (f *Feature) GetWin() (*layout.KeyMap, error) {
	return f.getKeys(cmdGetWinKeys)
}

// GetMac returns mac keyboard keys
func (f *Feature) GetMac() (*layout.KeyMap, error) {
	return f.getKeys(cmdGetMacKeys)
}

// SetWin sets win keyboard keys
func (f *Feature) SetWin(keyMap *layout.KeyMap) error {
	return f.setKeys(cmdGetWinKeys, cmdSetWinKeys, keyMap)
}

// SetMac sets mac keyboard keys
func (f *Feature) SetMac(keyMap *layout.KeyMap) error {
	return f.setKeys(cmdGetMacKeys, cmdSetMacKeys, keyMap)
}

// GetMacCodes returns mac keyboard key codes
func (f *Feature) GetMacCodes() ([]uint32, error) {
	return f.getKeyCodes(cmdGetMacKeys)
}

// GetWinCodes returns win keyboard key codes
func (f *Feature) GetWinCodes() ([]uint32, error) {
	return f.getKeyCodes(cmdGetWinKeys)
}

// SetMacCodes sets mac keyboard key codes
func (f *Feature) SetMacCodes(keys []uint32) error {
	return f.setKeyCodes(cmdSetMacKeys, keys)
}

// SetWinCodes sets win keyboard key codes
func (f *Feature) SetWinCodes(keys []uint32) error {
	return f.setKeyCodes(cmdSetWinKeys, keys)
}

func (f *Feature) GetMacros() ([]Macro, error) {
	resp, err := f.handle.Request(cmdGetMacro, 1032)
	if err != nil {
		return nil, err
	}
	return ParseMacro(resp)
}

// Parse raw keys
func (f *Feature) Parse(keys []uint32) (*layout.KeyMap, error) {
	return layout.Parse(keys, f.template)
}

func (f *Feature) getKeyCodes(cmd []byte) ([]uint32, error) {
	response, err := f.handle.Request(cmd, 1032)
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

func (f *Feature) getKeys(cmd []byte) (*layout.KeyMap, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	codes, err := f.getKeyCodes(cmd)
	if err != nil {
		return nil, err
	}
	keys, err := layout.Parse(codes, f.template)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (f *Feature) setKeyCodes(cmd []byte, keys []uint32) error {
	request := make([]byte, len(cmd)+len(keys)*4)
	copy(request, cmd)
	for i := 0; i < len(keys); i++ {
		offset := len(cmd) + (i * 4)
		unpackBytes(keys[i], request[offset:offset+4])
	}

	return f.handle.Send(request)
}

func (f *Feature) setKeys(cmdGet []byte, cmdSet []byte, keys *layout.KeyMap) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	codes, err := f.getKeyCodes(cmdGet)
	if err != nil {
		return err
	}
	err = keys.Apply(codes, f.template)
	if err != nil {
		return err
	}
	return f.setKeyCodes(cmdSet, codes)
}
