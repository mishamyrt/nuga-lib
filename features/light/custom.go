package light

import "github.com/mishamyrt/nuga-lib/layout"

// CustomBacklightMap represents custom backlight colors
type CustomBacklightMap map[layout.KeyName]RGB

// Bytes returns custom colors bytes
func (c CustomBacklightMap) Bytes(tpl *layout.Template) []byte {
	payload := make([]byte, 1024)
	for keyName, color := range c {
		position := tpl.GetPosition(keyName)
		payload[position] = color.R
		payload[position+126] = color.G
		payload[position+252] = color.B
	}
	return payload
}

// ParseCustomBacklight parses custom colors
func ParseCustomBacklight(payload []byte, tpl *layout.Template) (*CustomBacklightMap, error) {
	if payload[0] != byte(codeCustomEffectHeader) {
		return nil, ErrWrongCustomColorsHeader
	}
	raw := payload[7:]
	backlightMap := make(CustomBacklightMap)
	var color RGB
	for keyName, position := range *tpl {
		color = RGB{
			R: raw[position],
			G: raw[position+126],
			B: raw[position+252],
		}
		if color.R == 0 && color.G == 0 && color.B == 0 {
			continue
		}
		backlightMap[keyName] = color
	}
	return &backlightMap, nil
}
