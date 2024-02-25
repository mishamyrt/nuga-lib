package light

import "github.com/mishamyrt/nuga-lib/internal/bit"

const (
	// SpecificColor flag indicates that mode supports specific color.
	SpecificColor bit.Mask = 1 << iota
	// RandomColor flag indicates that mode supports random color
	RandomColor bit.Mask = 2 << iota
	// Speed flag indicates that mode supports speed.
	Speed bit.Mask = 4 << iota
)

// Mode represents keyboard light mode
type Mode struct {
	Name     string   `json:"name"`
	Code     byte     `json:"code"`
	Features bit.Mask `json:"features"`
}
