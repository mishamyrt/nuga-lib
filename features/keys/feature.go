package keys

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/internal/assert"
	"github.com/mishamyrt/nuga-lib/layout"
)

// Feature represents keyboard keys feature
type Feature struct {
	handle   hid.Handler
	template *layout.Template
	defaults *device.KeysState
}

// New creates keys feature instance.
func New(handle hid.Handler, model device.Model) (*Feature, error) {
	template := layout.GetKeystrokeTemplate(model)
	state, err := device.GetDefaults(model)
	if err != nil {
		return nil, err
	}
	return &Feature{
		handle:   handle,
		template: template,
		defaults: state.Data.Keys,
	}, nil
}

// GetWin returns win keyboard keys
func (f *Feature) GetWin() (*KeyMap, error) {
	m, err := f.getKeys(cmdGetWinKeys)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// SetWin sets win keyboard keys
func (f *Feature) SetWin(keyMap *KeyMap) error {
	return f.setKeys(cmdSetWinKeys, keyMap, f.defaults.Win)
}

// GetMac returns mac keyboard keys
func (f *Feature) GetMac() (*KeyMap, error) {
	return f.getKeys(cmdGetMacKeys)
}

// SetMac sets mac keyboard keys
func (f *Feature) SetMac(keyMap *KeyMap) error {
	return f.setKeys(cmdSetMacKeys, keyMap, f.defaults.Mac)
}

// GetRawMacros returns raw keyboard macros
func (f *Feature) GetRawMacros() ([]byte, error) {
	resp, err := f.handle.Request(cmdGetMacro, 1032)
	if err != nil {
		return nil, err
	}
	return resp[7:], nil
}

// SetRawMacros sets raw keyboard macros
func (f *Feature) SetRawMacros(data []byte) error {
	request := make([]byte, 0, 1032)
	request = append(request, cmdSetMacro...)
	request = append(request, data...)
	return f.handle.Send(request)
}

// GetMacros returns macros
func (f *Feature) GetMacros() (Macros, error) {
	raw, err := f.GetRawMacros()
	if err != nil {
		return nil, err
	}
	return ParseMacros(raw)
}

// SetMacros sets macros
func (f *Feature) SetMacros(macros Macros) error {
	data, err := macros.Bytes()
	if err != nil {
		return err
	}
	return f.SetRawMacros(data)
}

// GetStateData returns current raw keys state
func (f *Feature) GetStateData() (*device.KeysState, error) {
	mac, err := f.getRawKeys(cmdGetMacKeys)
	if err != nil {
		return nil, err
	}
	win, err := f.getRawKeys(cmdGetWinKeys)
	if err != nil {
		return nil, err
	}
	macros, err := f.GetRawMacros()
	if err != nil {
		return nil, err
	}
	return &device.KeysState{
		Mac:    mac,
		Win:    win,
		Macros: macros,
	}, nil
}

// SetStateData sets current raw keys state
func (f *Feature) SetStateData(data *device.KeysState) error {
	if err := f.setRawKeys(cmdSetMacKeys, data.Mac); err != nil {
		return err
	}
	if err := f.setRawKeys(cmdSetWinKeys, data.Win); err != nil {
		return err
	}
	return f.SetRawMacros(data.Macros)
}

func (f *Feature) getRawKeys(cmd []byte) ([]byte, error) {
	response, err := f.handle.Request(cmd, 1032)
	if err != nil {
		return nil, err
	}
	err = assert.SliceValue(response, 1, cmd[2])
	if err != nil {
		return nil, err
	}
	return response[7:], nil
}

func (f *Feature) setRawKeys(cmd []byte, data []byte) error {
	request := make([]byte, 0, 1032)
	request = append(request, cmd...)
	request = append(request, data...)
	return f.handle.Send(request)
}

func (f *Feature) getKeys(cmd []byte) (*KeyMap, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	raw, err := f.getRawKeys(cmd)
	if err != nil {
		return nil, err
	}
	keys, err := ParseKeyMap(raw, f.template)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (f *Feature) setKeys(cmdSet []byte, keys *KeyMap, defaults []byte) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	raw := keys.Bytes(f.template, defaults)
	return f.setRawKeys(cmdSet, raw)
}
