package keys

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/layout"
)

// Feature represents keyboard keys feature
type Feature struct {
	handle   hid.Handler
	template *layout.KeystrokeTemplate
}

// New creates keys feature instance.
func New(handle hid.Handler, model *device.Model) *Feature {
	var template *layout.KeystrokeTemplate
	if model != nil {
		template = layout.GetKeystrokeTemplate(*model)
	}
	return &Feature{
		handle:   handle,
		template: template,
	}
}

// GetWin returns win keyboard keys
func (f *Feature) GetWin() (*KeyMap, error) {
	return f.getKeys(cmdGetWinKeys)
}

// GetMac returns mac keyboard keys
func (f *Feature) GetMac() (*KeyMap, error) {
	return f.getKeys(cmdGetMacKeys)
}

// SetWin sets win keyboard keys
func (f *Feature) SetWin(keyMap *KeyMap) error {
	return f.setKeys(cmdGetWinKeys, cmdSetWinKeys, keyMap)
}

// SetMac sets mac keyboard keys
func (f *Feature) SetMac(keyMap *KeyMap) error {
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

// GetMacros returns macros
func (f *Feature) GetMacros() (Macros, error) {
	resp, err := f.handle.Request(cmdGetMacro, 1032)
	if err != nil {
		return nil, err
	}
	return ParseMacros(resp)
}

// SetMacros sets macros
func (f *Feature) SetMacros(macros Macros) error {
	request := make([]byte, 0, 1032)
	request = append(request, cmdSetMacro...)
	data, err := macros.Bytes()
	if err != nil {
		return err
	}
	request = append(request, data...)
	return f.handle.Send(request)
}

// Parse raw keys
func (f *Feature) Parse(keys []uint32) (*KeyMap, error) {
	return ParseKeyMap(keys, f.template)
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

func (f *Feature) getKeys(cmd []byte) (*KeyMap, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	codes, err := f.getKeyCodes(cmd)
	if err != nil {
		return nil, err
	}
	keys, err := ParseKeyMap(codes, f.template)
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

func (f *Feature) setKeys(cmdGet []byte, cmdSet []byte, keys *KeyMap) error {
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
