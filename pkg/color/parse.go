// Package color contains utils to work with colors
package color

import (
	"errors"
	"fmt"
)

// ErrInvalidValue is returned when HEX color string format is unexpected
var ErrInvalidFormat = errors.New("invalid hex color string")

// ParseHex converts a string with HEX color into an RGB structure
func ParseHex(s string) (*RGB, error) {
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
		err = ErrInvalidFormat
	}
	if err != nil {
		return nil, ErrInvalidFormat
	}
	return &c, err
}
