package keys

import (
	"github.com/mishamyrt/nuga-lib/internal/bit"
)

// MacroEvent represents macro event
type MacroEvent uint8

const (
	// KeyDown represents key down action
	KeyDown MacroEvent = 1
	// KeyUp represents key up action
	KeyUp MacroEvent = 0
)

// MacroAction represents macro action
type MacroAction struct {
	ShortCode byte       `json:"keyCode"`
	Event     MacroEvent `json:"type"`
	Delay     *uint16    `json:"delay,omitempty"`
}

type Macro struct {
	Actions []MacroAction `json:"actions"`
	Repeats uint16        `json:"repeats"`
}

func ParseMacro(payload []byte) ([]Macro, error) {
	if payload[0] != codeMacroHeader {
		return nil, ErrWrongMacroHeader
	}
	macros := make([]Macro, 0)
	data := payload[7:]
	for i := 0; i < len(data); i += 129 {
		if i+128 > len(data) {
			break
		}
		macroData := data[i : i+128]
		length := CountMacro(macroData)
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
		var (
			action MacroAction
			delay  uint16
		)
		if rest[i] == codeMacroKeyUp {
			action.Event = KeyUp
			action.ShortCode = rest[i+1]
		} else {
			// Detect long delay format
			if i+5 <= len(rest) &&
				rest[i+2] == codeMacroLongFormatFirst &&
				rest[i+3] == codeMacroLongFormatLast {
				low := uint16(rest[i])
				mid := uint16(rest[i+5])
				high := uint16(rest[i+4])
				if low == 3 {
					low = 0
				}
				delay = (offsetHigh * high) + (offsetMid * mid) + low
				action.Delay = &delay
				action.ShortCode = rest[i+1]
				i += 4 // long format has 6 bytes, so add 4 to default 2
			} else {
				delay = uint16(rest[i])
				action.Delay = &delay
				action.ShortCode = rest[i+1]
			}
			if delay == 0 {
				return nil, ErrWrongKeyCode
			} else if delay == 3 {
				delay = 0
			}
			if action.ShortCode == 0 {
				return nil, ErrWrongKeyCode
			}

			action.Event = KeyDown
		}
		macro.Actions = append(macro.Actions, action)
	}
	return macro, nil
}

func CountMacro(v []byte) int {
	for i := 0; i < len(v); i++ {
		if v[i] == 0 && v[i+1] == 0 {
			return i
		}
	}
	return 128
}
