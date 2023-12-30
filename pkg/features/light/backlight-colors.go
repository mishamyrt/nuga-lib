// Package light contains set of utils to work with keyboard.
package light

import (
	"nuga/pkg/color"
)

// ColorState represents keyboard color state.
type BacklightColors [ModesCount][ModeColorsCount]color.RGB

// Set color to state.
func (b *BacklightColors) Set(modeIndex uint8, colorIndex uint8, color *color.RGB) {
	b[modeIndex][colorIndex] = *color
}

// SetMacBacklight color to state.
func (b *BacklightColors) SetMac(modeIndex uint8, colorIndex uint8, color *color.RGB) {
	b.Set(modeIndex+ModesCount, colorIndex, color)
}

// SetWinBacklight color to state.
func (b *BacklightColors) SetWin(modeIndex uint8, colorIndex uint8, color *color.RGB) {
	b.Set(modeIndex, colorIndex, color)
}

// Get color from state.
func (b *BacklightColors) Get(modeIndex uint8, colorIndex uint8) *color.RGB {
	return &b[modeIndex][colorIndex]
}

// GetMac returns colors for mac OS mode.
func (b *BacklightColors) GetMac() [][]color.RGB {
	return b.toSlice(b[ModesCountPerOS:ModesCount])
}

// GetWin returns colors for win OS mode.
func (b *BacklightColors) GetWin() [][]color.RGB {
	return b.toSlice(b[0:ModesCountPerOS])
}

// Bytes converts color state to raw byte slice.
func (b *BacklightColors) Bytes() []byte {
	var buf = make([]byte, 0)
	var c color.RGB
	for effect := range b {
		for i := range b[effect] {
			c = b[effect][i]
			buf = append(buf, c.R, c.G, c.B)
		}
	}
	buf = append(buf, ColorsSuffix...)
	return buf
}

// Slice returns colors as slice
func (b *BacklightColors) Slice() [][]color.RGB {
	return b.toSlice(b[:])
}

func (b *BacklightColors) toSlice(modes [][7]color.RGB) [][]color.RGB {
	result := make([][]color.RGB, len(modes))
	for i, colors := range modes {
		result[i] = make([]color.RGB, 7)
		copy(result[i], colors[:])
	}
	return result
}

// ParseColors parses the raw byte slice into ColorState.
func ParseBacklightColors(data []byte) *BacklightColors {
	var colors BacklightColors
	var offset int
	for effect := 0; effect < 48; effect++ {
		for i := 0; i < 7; i++ {
			offset = (effect * 21) + (i * 3)
			colors[effect][i] = color.RGB{
				R: data[offset],
				G: data[offset+1],
				B: data[offset+2],
			}
		}
	}
	return &colors
}

// ColorsFromSlice loads color state from colors slice
func BacklightColorsFromSlice(modes [][]color.RGB) *BacklightColors {
	var colors BacklightColors
	for effect := 0; effect < ModesCount; effect++ {
		for i := 0; i < ModeColorsCount; i++ {
			colors[effect][i] = modes[effect][i]
		}
	}
	return &colors
}
