package light

import "github.com/mishamyrt/nuga-lib/layout"

// CustomEffectMap represents custom backlight colors
type CustomEffectMap map[layout.KeyName]Color

// Bytes returns custom colors bytes
func (c CustomEffectMap) Bytes(tpl *layout.Template) ([]byte, error) {
	const payloadSize = 1024
	payload := make([]byte, payloadSize)
	for keyName, color := range c {
		position := tpl.GetPosition(keyName)
		if position+252 >= payloadSize {
			return nil, ErrOutOfBounds
		}
		payload[position] = color.Red()
		payload[position+126] = color.Green()
		payload[position+252] = color.Blue()
	}
	return payload, nil
}

// ParseCustomEffect parses custom effect color map
func ParseCustomEffect(payload []byte, tpl *layout.Template) (*CustomEffectMap, error) {
	backlightMap := make(CustomEffectMap)
	for keyName, position := range *tpl {
		if int(position+252) >= len(payload) {
			return nil, ErrOutOfBounds
		}
		backlightMap[keyName] = FromRGB(
			payload[position],
			payload[position+126],
			payload[position+252],
		)
	}
	return &backlightMap, nil
}
