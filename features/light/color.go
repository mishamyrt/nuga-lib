package light

import (
	"fmt"

	"github.com/mishamyrt/nuga-lib/internal/hex"
)

// Color represents light color
type Color uint32

// String returns the hex string representation of the color
func (c Color) String() string {
	return hex.FromUInt32Color(uint32(c))
}

// Red returns the red component of the color.
func (c Color) Red() uint8 {
	return uint8(c >> 16 & 0xFF)
}

// Green returns the green component of the color.
func (c Color) Green() uint8 {
	return uint8(c >> 8 & 0xFF)
}

// Blue returns the blue component of the color.
func (c Color) Blue() uint8 {
	return uint8(c & 0xFF)
}

// FromRGB converts RGB color to uint32
func FromRGB(r, g, b uint8) Color {
	return Color(r)<<16 | Color(g)<<8 | Color(b)
}

// Color presets
const (
	// ColorRed represents the RGB value for the red color.
	ColorRed Color = 0xFF0000
	// ColorGreen represents the RGB value for the green color.
	ColorGreen Color = 0x00FF00
	// ColorBlue represents the RGB value for the blue color.
	ColorBlue Color = 0x0000FF
	// ColorYellow represents the RGB value for the yellow color.
	ColorYellow Color = 0xFFFF00
	// ColorPurple represents the RGB value for the purple color.
	ColorPurple Color = 0xFF00FF
	// ColorCyan represents the RGB value for the cyan color.
	ColorCyan Color = 0x00FFFF
	// ColorWhite represents the RGB value for the white color.
	ColorWhite Color = 0xFFFFFF
)

// ParseHexColor converts a string with HEX color into an RGB structure
func ParseHexColor(s string) (Color, error) {
	var (
		err     error
		r, g, b uint8
	)
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &r, &g, &b)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &r, &g, &b)
		r *= 0x11
		g *= 0x11
		b *= 0x11
	default:
		err = ErrInvalidColorFormat
	}
	if err != nil {
		return 0, ErrInvalidColorFormat
	}
	return FromRGB(r, g, b), err
}
