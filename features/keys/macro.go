package keys

import (
	"math"

	"github.com/mishamyrt/nuga-lib/features/keys/layout"
	"github.com/mishamyrt/nuga-lib/internal/bit"
)

// MacroEvent represents macro event
type MacroEvent string

const (
	// KeyDown represents key down action event
	KeyDown MacroEvent = "down"
	// KeyUp represents key up action event
	KeyUp MacroEvent = "up"
)

// MacroAction represents macro action
type MacroAction struct {
	Key   layout.KeyName `json:"key"`
	Event MacroEvent     `json:"type"`
	Delay uint16         `json:"delay,omitempty"`
}

// Macro represents keyboard macro
type Macro struct {
	Actions []MacroAction `json:"actions"`
	Repeats uint16        `json:"repeats"`
}

// Bytes returns macro bytes
func (m *Macro) Bytes() ([]byte, error) {
	b := make([]byte, 0, 128)
	high, low := bit.Split16(m.Repeats)
	b = append(b, high, low)
	for _, a := range m.Actions {
		delay := a.Delay
		keyCode := layout.FindShortKeyCode(a.Key)
		if delay >= 128 {
			high := byte(math.Floor(float64(a.Delay) / offsetHigh))
			rest := a.Delay - (offsetHigh * uint16(high))
			mid := byte(math.Floor(float64(a.Delay) / offsetMid))
			rest = rest - (offsetMid * uint16(mid))
			if rest > 255 {
				return nil, ErrWrongDelay
			}
			low := byte(rest)
			if low == 0 {
				low = 3
			}
			if a.Event == KeyUp {
				low += 128
			}
			b = append(b, low, keyCode, codeMacroLongFormatFirst, codeMacroLongFormatLast, high, mid)
		} else {
			delay := byte(a.Delay)
			if delay == 0 {
				delay = 3
			}
			if a.Event == KeyUp {
				delay += 128
			}
			b = append(b, delay, keyCode)
		}
	}
	return b, nil
}

// Macros represents keyboard macros
type Macros []Macro

// Bytes returns macros bytes
func (m Macros) Bytes() ([]byte, error) {
	b := make([]byte, 0, 1024)
	for _, macro := range m {
		bytes, err := macro.Bytes()
		if err != nil {
			return nil, err
		}
		padding := 128 - len(bytes)
		p := make([]byte, 0, 128)
		p = append(p, bytes...)
		p = append(p, make([]byte, padding)...)
		b = append(b, p...)
	}
	if len(b) < 1024 {
		b = append(b, make([]byte, 1024-len(b))...)
	}
	return b, nil
}

// ParseMacros parses macros payload
func ParseMacros(payload []byte) (Macros, error) {
	if payload[0] != codeMacroHeader {
		return nil, ErrWrongMacroHeader
	}
	macros := make([]Macro, 0)
	data := payload[7:]
	for i := 0; i < len(data); i += 128 {
		if i+128 > len(data) {
			break
		}
		macroData := data[i : i+128]
		length := countMacro(macroData)
		if length == 0 {
			break
		}
		macro, err := parseMacroPart(macroData[0:length])
		if err != nil {
			return nil, err
		}
		macros = append(macros, *macro)
	}
	return macros, nil
}

func parseMacroPart(payload []byte) (*Macro, error) {
	macro := &Macro{
		Actions: make([]MacroAction, 0),
	}
	rest := payload[:]
	if len(rest)%2 != 0 {
		macro.Repeats = uint16(rest[0])
		rest = rest[1:]
	} else {
		macro.Repeats = bit.Combine16(rest[0], rest[1])
		rest = rest[2:]
	}

	for i := 0; i < len(rest); i += 2 {
		var low byte
		name := layout.FindKeyNameByShortCode(rest[i+1])
		if name == layout.KeyNone {
			return nil, ErrWrongKeyCode
		}
		action := MacroAction{
			Key: name,
		}
		if rest[i] >= 128 {
			action.Event = KeyUp
			low = rest[i] - 128
		} else {
			action.Event = KeyDown
			low = rest[i]
		}
		if low == 3 {
			low = 0
		}
		if i+5 <= len(rest) &&
			rest[i+2] == codeMacroLongFormatFirst &&
			rest[i+3] == codeMacroLongFormatLast {
			mid := uint16(rest[i+5])
			high := uint16(rest[i+4])
			if low == 3 {
				low = 0
			}
			action.Delay = (offsetHigh * high) + (offsetMid * mid) + uint16(low)
			i += 4 // long format has 6 bytes, so add 4 to default 2
		} else {
			action.Delay = uint16(low)
		}
		macro.Actions = append(macro.Actions, action)
	}
	return macro, nil
}

func countMacro(v []byte) int {
	for i := 0; i < len(v); i++ {
		if v[i] == 0 && v[i+1] == 0 {
			return i
		}
	}
	return 128
}
