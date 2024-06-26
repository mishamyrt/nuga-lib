// Package light contains set of utils to work with keyboard.
package light

// BacklightColors represents keyboard color state.
type BacklightColors [ModesCount][ModeColorsCount]Color

// Set color to state.
func (b *BacklightColors) Set(modeIndex uint8, colorIndex uint8, color *Color) {
	b[modeIndex][colorIndex] = *color
}

// SetMac color to state.
func (b *BacklightColors) SetMac(modeIndex uint8, colorIndex uint8, color *Color) {
	b.Set(modeIndex+ModesCountPerOS, colorIndex, color)
}

// SetWin color to state.
func (b *BacklightColors) SetWin(modeIndex uint8, colorIndex uint8, color *Color) {
	b.Set(modeIndex, colorIndex, color)
}

// Get color from state.
func (b *BacklightColors) Get(modeIndex uint8, colorIndex uint8) *Color {
	return &b[modeIndex][colorIndex]
}

// GetMac returns colors for mac OS mode.
func (b *BacklightColors) GetMac() [][]Color {
	return b.toSlice(b[ModesCountPerOS:ModesCount])
}

// GetWin returns colors for win OS mode.
func (b *BacklightColors) GetWin() [][]Color {
	return b.toSlice(b[0:ModesCountPerOS])
}

// Bytes converts color state to raw byte slice.
func (b *BacklightColors) Bytes() []byte {
	payloadSize := ModesCount*ModeColorsCount*3 + len(ColorsSuffix)
	var buf = make([]byte, 0, payloadSize)
	var c Color
	for effect := range b {
		for i := range b[effect] {
			c = b[effect][i]
			buf = append(buf, c.Red(), c.Green(), c.Blue())
		}
	}
	buf = append(buf, ColorsSuffix...)
	return buf
}

// Slice returns colors as slice
func (b *BacklightColors) Slice() [][]Color {
	return b.toSlice(b[:])
}

func (b *BacklightColors) toSlice(modes [][7]Color) [][]Color {
	result := make([][]Color, len(modes))
	for i, colors := range modes {
		result[i] = make([]Color, 7)
		copy(result[i], colors[:])
	}
	return result
}

// ParseBacklightColors parses the raw byte slice into BacklightColors.
func ParseBacklightColors(data []byte) (*BacklightColors, error) {
	if len(data) < (ModesCount * 21) {
		return nil, ErrOutOfBounds
	}
	var (
		colors BacklightColors
		offset int
	)
	for effect := 0; effect < ModesCount; effect++ {
		for i := 0; i < ModeColorsCount; i++ {
			offset = (effect * (ModeColorsCount * 3)) + (i * 3)
			colors[effect][i] = FromRGB(data[offset], data[offset+1], data[offset+2])
		}
	}
	return &colors, nil
}

// BacklightColorsFromSlice loads color state from colors slice
func BacklightColorsFromSlice(modes [][]Color) *BacklightColors {
	var colors BacklightColors
	for effect := 0; effect < ModesCount; effect++ {
		for i := 0; i < ModeColorsCount; i++ {
			colors[effect][i] = modes[effect][i]
		}
	}
	return &colors
}
