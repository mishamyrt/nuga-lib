package light

import (
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/layout"
)

// Feature represents keyboard light feature
type Feature struct {
	handle   hid.Handler
	template *layout.Template
}

// New creates light feature instance.
func New(handle hid.Handler, model device.Model) *Feature {
	return &Feature{
		handle:   handle,
		template: layout.GetBacklightTemplate(model),
	}
}

// GetRawEffects returns raw effects data.
func (f *Feature) GetRawEffects() ([]byte, error) {
	raw, err := f.handle.Request(CmdGetParams, 270)
	if err != nil {
		return nil, err
	}
	return raw[effectsOffset : effectsOffset+effectsLength], nil
}

// SetRawEffects sets raw effects data.
func (f *Feature) SetRawEffects(data []byte) error {
	previous, err := f.GetRawEffects()
	if err != nil {
		return err
	}
	paramsRequest := make([]byte, 0, 1032)
	paramsRequest = append(paramsRequest, CmdSetParams...)
	paramsRequest = append(paramsRequest, data...)
	paramsRequest = append(paramsRequest, previous...)
	paramsRequest = append(paramsRequest, make([]byte, 770)...)
	return f.handle.Send(paramsRequest)
}

// GetEffects returns keyboard effects.
func (f *Feature) GetEffects() (*Effects, error) {
	raw, err := f.GetRawEffects()
	if err != nil {
		return nil, err
	}
	return ParseEffects(raw)
}

// SetEffects sets keyboard effects.
func (f *Feature) SetEffects(e *Effects) error {
	return f.SetRawEffects(e.Bytes())
}

// GetRawBacklightColors returns raw keyboard colors.
func (f *Feature) GetRawBacklightColors() ([]byte, error) {
	resp, err := f.handle.Request(CmdGetColors, 1050)
	if err != nil {
		return nil, err
	}
	return resp[lightsStartOffset : len(resp)-lightsEndOffset], nil
}

// SetRawBacklightColors sets raw keyboard colors.
func (f *Feature) SetRawBacklightColors(data []byte) error {
	cmdLength := len(CmdSetColors) + len(data)
	request := make([]byte, 0, cmdLength)
	request = append(request, CmdSetColors...)
	request = append(request, data...)
	return f.handle.Send(request)
}

// GetBacklightColors returns keyboard backlight colors state.
func (f *Feature) GetBacklightColors() (*BacklightColors, error) {
	raw, err := f.GetRawBacklightColors()
	if err != nil {
		return nil, err
	}
	return ParseBacklightColors(raw)
}

// SetBacklightColors sets keyboard backlight color state.
func (f *Feature) SetBacklightColors(colors *BacklightColors) error {
	return f.SetRawBacklightColors(colors.Bytes())
}

// GetRawCustomEffect returns current custom effect raw colors.
func (f *Feature) GetRawCustomEffect() ([]byte, error) {
	resp, err := f.handle.Request(CmdGetCustomFirstPage, 1050)
	if err != nil {
		return nil, err
	}
	return resp[lightsStartOffset : len(resp)-lightsEndOffset], nil
}

// SetRawCustomEffect applies raw custom effect.
func (f *Feature) SetRawCustomEffect(p []byte) error {
	req := make([]byte, 0, len(CmdSetCustomFirstPage)+len(p))
	req = append(req, CmdSetCustomFirstPage...)
	req = append(req, p...)
	return f.handle.Send(req)
}

// GetCustomEffect returns current custom effect colors.
func (f *Feature) GetCustomEffect() (*CustomEffectMap, error) {
	if f.template == nil {
		return nil, ErrNoTemplate
	}
	raw, err := f.GetRawCustomEffect()
	if err != nil {
		return nil, err
	}
	return ParseCustomEffect(raw, f.template)
}

// SetCustomEffect sets current custom effect colors.
func (f *Feature) SetCustomEffect(colors *CustomEffectMap) error {
	if f.template == nil {
		return ErrNoTemplate
	}
	payload, err := colors.Bytes(f.template)
	if err != nil {
		return err
	}
	return f.SetRawCustomEffect(payload)
}

// GetStateData returns current keyboard light state.
func (f *Feature) GetStateData() (*device.LightsState, error) {
	colors, err := f.GetRawBacklightColors()
	if err != nil {
		return nil, err
	}
	params, err := f.GetRawEffects()
	if err != nil {
		return nil, err
	}
	customEffect, err := f.GetRawCustomEffect()
	if err != nil {
		return nil, err
	}
	return &device.LightsState{
		Colors:       colors,
		Params:       params,
		CustomEffect: customEffect,
	}, nil
}

// SetStateData sets keyboard light state.
func (f *Feature) SetStateData(data *device.LightsState) error {
	if err := f.SetRawBacklightColors(data.Colors); err != nil {
		return err
	}
	if err := f.SetRawEffects(data.Params); err != nil {
		return err
	}
	return f.SetRawCustomEffect(data.CustomEffect)
}
