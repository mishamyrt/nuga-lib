package light

const (
	// ModesCount represents total mode count
	ModesCount = 48
	// ModesCountPerOS represents mode count per OS (win/mac)
	ModesCountPerOS = ModesCount / 2
	// ModeColorsCount colors count per mode
	ModeColorsCount = 7
)

const (
	// ParamsLength represents Halo75 params count.
	ParamsLength = 123
	// EffectsOffset represents effect start offset.
	EffectsOffset = 15

	// ColorParamsOffset represents Halo75 color start offset.
	ColorParamsOffset = 22
	// ColorsStartOffset represents start offset of colors response.
	ColorsStartOffset = 7
	// ColorEndOffset represents end offset of colors response.
	ColorEndOffset = 18
)

// Commands
var (
	// CmdGetParams represents keyboard params reading command.
	CmdGetParams = []byte{0x05, 0x83, 0xb6, 0x00, 0x00, 0x00}
	// CmdSetParams represents keyboard params writing command header.
	CmdSetParams = []byte{
		0x06, 0x03, 0xb6, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5a, 0xa5,
	}
	// CmdGetColors represents keyboard color reading command.
	CmdGetColors = []byte{0x05, 0x88, 0xb8, 0x00, 0x00, 0x00}
	// CmdSetColors represents keyboard color writing command header.
	CmdSetColors = []byte{0x06, 0x08, 0xb8, 0x00, 0x40, 0x00, 0x00, 0x00}
)

// Command parts
var (
	// ParamsHeader represents keyboard params header.
	ParamsHeader = []byte{0x03, 0x03}
	// ParamsBacklightHeader represents keyboard backlight params header.
	ParamsBacklightHeader = []byte{0x00, 0x00}
	// ParamsMiscHeader represents keyboard misc params header.
	ParamsMiscHeader = []byte{0x20, 0x00, 0x00, 0x00, 0x00, 0x00}
	// ParamsSuffix represents keyboard params suffix.
	ParamsSuffix = []byte{
		0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x01, 0x01, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x5A, 0xA5,
	}
	// ColorsSuffix represents keyboard params suffix.
	ColorsSuffix = []byte{
		0x90, 0x0E, 0xBF, 0xEF, 0xF0, 0xC2, 0x05, 0xE4, 0xA3, 0xF0,
		0xE4, 0x90, 0x0E, 0xC1, 0xF0, 0x90,
	}
)
