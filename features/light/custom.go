package light

import "github.com/mishamyrt/nuga-lib/layout"

// CustomEffectMap represents custom backlight colors
type CustomEffectMap map[layout.KeyName]Color

// Bytes returns custom colors bytes
func (c CustomEffectMap) Bytes(tpl *layout.Template) []byte {
	payload := make([]byte, 1024)
	for keyName, color := range c {
		position := tpl.GetPosition(keyName)
		payload[position] = color.Red()
		payload[position+126] = color.Green()
		payload[position+252] = color.Blue()
	}
	return payload
}

// ParseCustomEffect parses custom effect color map
func ParseCustomEffect(payload []byte, tpl *layout.Template) *CustomEffectMap {
	backlightMap := make(CustomEffectMap)
	for keyName, position := range *tpl {
		backlightMap[keyName] = FromRGB(
			payload[position],
			payload[position+126],
			payload[position+252],
		)
	}
	return &backlightMap
}
