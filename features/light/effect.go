package light

import (
	"math"
)

// EffectParams represents keyboard effect parameters.
type EffectParams struct {
	// Color represents effect color. Number from 0 to 7
	Color uint8 `json:"color"`
	// Speed represents effect speed. Number from 0 to 4
	Speed uint8 `json:"speed"`
	// Speed represents effect brightness. Number from 0 to 4
	Brightness uint8 `json:"brightness"`
}

// Effects represents keyboard effects state.
type Effects struct {
	Backlight BacklightEffect `json:"backlight"`
	Sidelight MiscEffect      `json:"sidelight"`
	Halo      MiscEffect      `json:"halo"`
	// Debounce represents key press debounce time in milliseconds.
	// This parameter has nothing to do with backlight,
	// A NuPhy representative explained it as follows:
	// Putting it in the lighting setting is a matter of design, for the sake of beauty.
	Debounce uint8 `json:"debounce"`
}

// Bytes returns effects as a raw byte slice.
func (b *Effects) Bytes() []byte {
	buf := make([]byte, 0, 123)
	buf = append(buf, ParamsHeader...)
	buf = append(buf, b.Debounce)
	buf = append(buf, ParamsBacklightHeader...)
	buf = append(buf, b.Backlight.Mode.Code)
	buf = append(buf, ParamsMiscHeader...)
	buf = append(
		buf,
		b.Sidelight.Mode.Code,
		b.Sidelight.Color,
		b.Sidelight.Brightness,
		b.Sidelight.Speed,
	)
	buf = append(
		buf,
		b.Halo.Mode.Code,
		b.Halo.Color,
		b.Halo.Brightness,
		b.Halo.Speed,
	)
	buf = append(buf, 0xFF, 0xFF)
	for _, param := range b.Backlight.Params {
		value := (param.Speed * 16) + param.Brightness
		buf = append(buf, param.Color, value)
	}
	buf = append(buf, ParamsSuffix...)
	return buf
}

// ParseEffects parses raw bytes to effects struct.
func ParseEffects(data []byte) *Effects {
	result := &Effects{}
	result.Debounce = data[2]
	result.Sidelight = MiscEffect{
		Mode: SidelightDomain.Find(data[12]),
		EffectParams: EffectParams{
			Color:      data[13],
			Brightness: data[14],
			Speed:      data[15],
		},
	}
	result.Halo = MiscEffect{
		Mode: HaloDomain.Find(data[16]),
		EffectParams: EffectParams{
			Color:      data[17],
			Brightness: data[18],
			Speed:      data[19],
		},
	}
	result.Backlight.Mode = BacklightDomain.Find(data[5])
	result.Backlight.Params = make([]EffectParams, 29)
	for i := range result.Backlight.Params {
		offset := ColorParamsOffset + (i * 2)
		value := data[offset+1]
		result.Backlight.Params[i] = EffectParams{
			Color:      data[offset],
			Speed:      uint8(math.Floor(float64(value) / 16)),
			Brightness: value % 16,
		}
	}
	return result
}
