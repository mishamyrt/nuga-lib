package light

import (
	"nuga/pkg/color"
	"nuga/pkg/hid"
)

type Feature struct {
	handle *hid.Device
}

// Open BYK916 NuPhy keyboard lights.
func New(handle *hid.Device) *Feature {
	return &Feature{
		handle: handle,
	}
}

// GetRawEffects returns raw effects data.
func (f *Feature) GetRawEffects() ([]byte, error) {
	response, err := f.handle.Request(CmdGetParams, 270)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetEffectsSlice returns trimmed effects slice.
func (f *Feature) GetEffectsSlice() ([]byte, error) {
	raw, err := f.GetRawEffects()
	if err != nil {
		return nil, err
	}
	return raw[EffectsOffset : EffectsOffset+ParamsLength], nil
}

// GetEffects returns keyboard effects.
func (f *Feature) GetEffects() (*Effects, error) {
	params, err := f.GetEffectsSlice()
	if err != nil {
		return nil, err
	}
	effects := ParseEffects(params)
	return effects, err
}

// GetRawColors returns raw keyboard colors.
func (f *Feature) GetRawColors() ([]byte, error) {
	var colors []byte
	colors, err := f.handle.Request(CmdGetColors, 1050)
	if err != nil {
		return colors, err
	}
	return colors, err
}

// GetBacklightColors returns keyboard backlight colors state.
func (f *Feature) GetBacklightColors() (*BacklightColors, error) {
	raw, err := f.GetRawColors()
	if err != nil {
		return nil, err
	}
	colorSubset := raw[ColorsStartOffset : len(raw)-ColorEndOffset]
	return ParseBacklightColors(colorSubset), err
}

// SetBacklightColors sets keyboard backlight color state.
func (f *Feature) SetBacklightColors(colors *BacklightColors) error {
	colorsContent := colors.Bytes()
	cmdLength := len(CmdSetColors) + len(colorsContent)
	request := make([]byte, 0, cmdLength)
	request = append(request, CmdSetColors...)
	request = append(request, colorsContent...)
	return f.handle.SendWithRetries(request)
}

// ResetColors resets colors to defaults.
func (f *Feature) ResetColors() error {
	var colors BacklightColors
	for i := range colors {
		colors[i][0] = color.Red
		colors[i][1] = color.Green
		colors[i][2] = color.Blue
		colors[i][3] = color.Yellow
		colors[i][4] = color.Purple
		colors[i][5] = color.Cyan
		colors[i][6] = color.White
	}
	return f.SetBacklightColors(&colors)
}

// SetEffects sets keyboard effects.
func (f *Feature) SetEffects(p *Effects) error {
	currentParams, err := f.GetEffectsSlice()
	if err != nil {
		return err
	}
	paramsRequest := make([]byte, 0)
	paramsRequest = append(paramsRequest, CmdSetParams...)
	paramsRequest = append(paramsRequest, p.Bytes()...)
	paramsRequest = append(paramsRequest, currentParams...)
	paramsRequest = append(paramsRequest, make([]byte, 770)...)
	return f.handle.SendWithRetries(paramsRequest)
}
