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
func New(handle hid.Handler, model *device.Model) *Feature {
	var template *layout.Template
	if model != nil {
		template = layout.GetBacklightTemplate(*model)
	}
	return &Feature{
		handle:   handle,
		template: template,
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
	return f.handle.Send(request)
}

// ResetColors resets colors to defaults.
func (f *Feature) ResetColors() error {
	var colors BacklightColors
	for i := range colors {
		colors[i][0] = ColorRed
		colors[i][1] = ColorGreen
		colors[i][2] = ColorBlue
		colors[i][3] = ColorYellow
		colors[i][4] = ColorPurple
		colors[i][5] = ColorCyan
		colors[i][6] = ColorWhite
	}
	return f.SetBacklightColors(&colors)
}

// SetEffects sets keyboard effects.
func (f *Feature) SetEffects(p *Effects) error {
	currentParams, err := f.GetEffectsSlice()
	if err != nil {
		return err
	}
	paramsRequest := make([]byte, 0, 1032)
	paramsRequest = append(paramsRequest, CmdSetParams...)
	paramsRequest = append(paramsRequest, p.Bytes()...)
	paramsRequest = append(paramsRequest, currentParams...)
	paramsRequest = append(paramsRequest, make([]byte, 770)...)
	return f.handle.Send(paramsRequest)
}

// GetCustomEffectColors returns current custom effect colors.
func (f *Feature) GetCustomEffectColors() (*CustomBacklightMap, error) {
	if f.template == nil {
		return nil, ErrNoCustomColorsTemplate
	}
	raw, err := f.handle.Request(CmdGetCustomFirstPage, 1050)
	if err != nil {
		return nil, err
	}
	return ParseCustomBacklight(raw, f.template)
}

// SetCustomEffectColors sets current custom effect colors.
func (f *Feature) SetCustomEffectColors(colors *CustomBacklightMap) error {
	if f.template == nil {
		return ErrNoCustomColorsTemplate
	}
	payload := colors.Bytes(f.template)
	req := make([]byte, 0, len(CmdSetCustomFirstPage)+len(payload))
	req = append(req, CmdSetCustomFirstPage...)
	req = append(req, payload...)
	return f.handle.Send(req)
}
