package light

// Domain is a type of keyboard lights with its mode
type Domain struct {
	Name  string `json:"name"`
	Modes []Mode `json:"modes"`
}

// Find mode in list. Returns nil if not found
func (d *Domain) Find(code byte) *Mode {
	for i := range d.Modes {
		if d.Modes[i].Code == code {
			return &d.Modes[i]
		}
	}
	return nil
}

var (
	// BacklightOff represents keyboard backlight off mode.
	BacklightOff = Mode{
		"Lights off",
		0,
		0,
	}
	// BacklightStatic represents keyboard backlight fixed mode.
	BacklightStatic = Mode{
		"Static",
		1,
		SpecificColor | RandomColor,
	}
	// BacklightRespire represents keyboard backlight respire mode.
	BacklightRespire = Mode{
		"Respire",
		2,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightRainbow represents keyboard backlight rainbow mode.
	BacklightRainbow = Mode{
		"Rainbow",
		3,
		Speed,
	}
	// BacklightFlashAway represents keyboard backlight flash away mode.
	BacklightFlashAway = Mode{
		"Flash away",
		4,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightRainbowWheel represents keyboard backlight rainbow wheel mode.
	BacklightRainbowWheel = Mode{
		"Rainbow wheel",
		6,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightRipplesShining represents keyboard backlight ripples shining mode.
	BacklightRipplesShining = Mode{
		"Ripples shining",
		7,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightStarsTwinkle represents keyboard backlight ripples shining mode.
	BacklightStarsTwinkle = Mode{
		"Stars twinkle",
		8,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightNeonStream represents keyboard backlight neon stream mode.
	BacklightNeonStream = Mode{
		"Neon stream",
		11,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightReaction represents keyboard backlight reaction mode.
	BacklightReaction = Mode{
		"Reaction",
		12,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightSineWave represents keyboard backlight sine wave mode.
	BacklightSineWave = Mode{
		"Sine wave",
		13,
		SpecificColor | RandomColor | Speed,
	}
	// BacklightRotatingWindmill represents keyboard backlight rotating windmill mode.
	BacklightRotatingWindmill = Mode{
		"Rotating windmill",
		15,
		Speed,
	}
	// BacklightColorfulWaterfall represents keyboard backlight colorful waterfall mode.
	BacklightColorfulWaterfall = Mode{
		"Colorful waterfall",
		16,
		Speed,
	}
	// BacklightBlossoming represents keyboard backlight blossoming mode.
	BacklightBlossoming = Mode{
		"Blossoming",
		17,
		Speed,
	}
	// BacklightCustom represents keyboard backlight game mode.
	BacklightCustom = Mode{
		"Custom",
		18,
		0,
	}
)

var (
	// HalolightOff represents keyboard halolight off mode.
	HalolightOff = Mode{
		"Lights off",
		0,
		0,
	}
	// HalolightRainbowWheel represents keyboard halolight rainbow wheel mode.
	HalolightRainbowWheel = Mode{
		"Rainbow wheel",
		1,
		Speed | RandomColor,
	}
	// HalolightRainbow represents keyboard halolight rainbow mode.
	HalolightRainbow = Mode{
		"Rainbow",
		2,
		Speed | RandomColor,
	}
	// HalolightFixed represents keyboard halolight fixed mode.
	HalolightFixed = Mode{
		"Fixed",
		3,
		SpecificColor,
	}
	// HalolightRespire represents keyboard halolight respire mode.
	HalolightRespire = Mode{
		"Respire",
		4,
		SpecificColor | Speed,
	}
)

var (
	// SidelightOff represents keyboard sidelight off mode.
	SidelightOff = Mode{
		"Lights off",
		0,
		0,
	}
	// SidelightRainbowStream represents keyboard sidelight rainbow stream mode.
	SidelightRainbowStream = Mode{
		"Rainbow stream",
		1,
		Speed | RandomColor,
	}
	// SidelightRainbow represents keyboard sidelight rainbow mode.
	SidelightRainbow = Mode{
		"Rainbow",
		2,
		Speed | RandomColor,
	}
	// SidelightFixed represents keyboard sidelight fixed mode.
	SidelightFixed = Mode{
		"Fixed on",
		3,
		SpecificColor,
	}
	// SidelightRespire represents keyboard sidelight respire mode.
	SidelightRespire = Mode{
		"Respire",
		4,
		SpecificColor | Speed,
	}
	// SidelightNeonStream represents keyboard sidelight neon stream mode.
	SidelightNeonStream = Mode{
		"Neon stream",
		5,
		Speed | RandomColor,
	}
)

// BacklightDomain is a domain of backlight effects.
var BacklightDomain = Domain{
	Name: "Backlight",
	Modes: []Mode{
		BacklightOff,
		BacklightStatic,
		BacklightRespire,
		BacklightRainbow,
		BacklightFlashAway,
		BacklightRainbowWheel,
		BacklightRipplesShining,
		BacklightStarsTwinkle,
		BacklightNeonStream,
		BacklightReaction,
		BacklightSineWave,
		BacklightRotatingWindmill,
		BacklightColorfulWaterfall,
		BacklightBlossoming,
		BacklightCustom,
	},
}

// HaloDomain is a domain of halolight effects.
var HaloDomain = Domain{
	Name: "Halo",
	Modes: []Mode{
		HalolightOff,
		HalolightRainbowWheel,
		HalolightRainbow,
		HalolightFixed,
		HalolightRespire,
	},
}

// SidelightDomain is a domain of sidelight effects.
var SidelightDomain = Domain{
	Name: "Sidelight",
	Modes: []Mode{
		SidelightOff,
		SidelightRainbowStream,
		SidelightRainbow,
		SidelightFixed,
		SidelightRespire,
		SidelightNeonStream,
	},
}
