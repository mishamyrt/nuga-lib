package light

import (
	"errors"
	"fmt"
)

// RGB represents light color
type RGB struct {
	R, G, B uint8
}

// Color presets
var (
	// ColorRed represents the RGB value for the red color.
	ColorRed = RGB{0xFF, 0x00, 0x00}
	// ColorGreen represents the RGB value for the green color.
	ColorGreen = RGB{0x00, 0xFF, 0x00}
	// ColorBlue represents the RGB value for the blue color.
	ColorBlue = RGB{0x00, 0x00, 0xFF}
	// ColorYellow represents the RGB value for the yellow color.
	ColorYellow = RGB{0xFF, 0xFF, 0x00}
	// ColorPurple represents the RGB value for the purple color.
	ColorPurple = RGB{0xFF, 0x00, 0xFF}
	// ColorCyan represents the RGB value for the cyan color.
	ColorCyan = RGB{0x00, 0xFF, 0xFF}
	// ColorWhite represents the RGB value for the white color.
	ColorWhite = RGB{0xFF, 0xFF, 0xFF}
)

// ErrInvalidColorFormat is returned when HEX color string format is unexpected
var ErrInvalidColorFormat = errors.New("invalid HEX color string format")

// ParseHexColor converts a string with HEX color into an RGB structure
func ParseHexColor(s string) (*RGB, error) {
	c := RGB{}
	var err error
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits
		c.R *= 0x11
		c.G *= 0x11
		c.B *= 0x11
	default:
		err = ErrInvalidColorFormat
	}
	if err != nil {
		return nil, ErrInvalidColorFormat
	}
	return &c, err
}
