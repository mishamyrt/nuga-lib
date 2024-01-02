package light

// BacklightEffect represents keyboard backlight effect state.
type BacklightEffect struct {
	Mode   *Mode          `json:"mode"`
	Params []EffectParams `json:"params"`
}

// CurrentParams returns current effect params.
func (p *BacklightEffect) CurrentParams() *EffectParams {
	code := p.Mode.Code
	if code == 0 {
		return nil
	}
	return &p.Params[code-1]
}

// SetBrightness sets current effect brightness.
func (p *BacklightEffect) SetBrightness(brightness uint8) error {
	if p.Mode.Code == 0 {
		return ErrLightsOff
	}
	if brightness > 4 {
		return ErrOutOfRange
	}
	p.CurrentParams().Brightness = brightness
	return nil
}

// SetSpeed sets current effect speed.
func (p *BacklightEffect) SetSpeed(speed uint8) error {
	if p.Mode.Code == 0 {
		return ErrLightsOff
	}
	if !p.Mode.Features.IsSet(Speed) {
		return ErrNotSupported
	}
	if speed > 4 {
		return ErrOutOfRange
	}
	p.CurrentParams().Speed = speed
	return nil
}

// SetColor sets current effect color.
func (p *BacklightEffect) SetColor(color uint8) error {
	if p.Mode.Code == 0 {
		return ErrLightsOff
	}
	if !p.Mode.Features.IsAnySet(SpecificColor, RandomColor) {
		return ErrNotSupported
	}
	if color == 7 && !p.Mode.Features.IsSet(RandomColor) {
		return ErrNotSupported
	}
	p.CurrentParams().Color = color
	return nil
}
