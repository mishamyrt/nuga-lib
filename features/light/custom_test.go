package light_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/layout"
)

var tpl = layout.Template{
	layout.KeyM: 0,
	layout.KeyY: 1,
	layout.KeyR: 2,
	layout.KeyT: 3,
}

var colorMap = light.CustomEffectMap{
	layout.KeyM: 0xFF0000,
	layout.KeyY: 0x00FF00,
	layout.KeyR: 0x0000FF,
	layout.KeyT: 0xFFFFFF,
}

func TestParseCustomBacklight(t *testing.T) {
	t.Parallel()
	payload := make([]byte, 1043)
	for keyName, position := range tpl {
		payload[position] = colorMap[keyName].Red()
		payload[position+126] = colorMap[keyName].Green()
		payload[position+252] = colorMap[keyName].Blue()
	}
	customMap, err := light.ParseCustomEffect(payload, &tpl)
	if err != nil {
		t.Fatalf("unexpected error on custom effect parsing: %v", err)
	}
	if len(*customMap) != 4 {
		t.Fatalf("expected 4 colors, got %d", len(*customMap))
	}
	for keyName, color := range colorMap {
		if (*customMap)[keyName] != color {
			t.Fatalf("expected color %v for key %v, got %v", color, keyName, (*customMap)[keyName])
		}
	}
}

func TestCustomBacklightToBytes(t *testing.T) {
	t.Parallel()
	payload, err := colorMap.Bytes(&tpl)
	if err != nil {
		t.Fatalf("unexpected error on color map serialization: %v", err)
	}
	if len(payload) != 1024 {
		t.Fatalf("expected 1024 bytes, got %d", len(payload))
	}
	for keyName, position := range tpl {
		r := payload[position]
		g := payload[position+126]
		b := payload[position+252]
		if r != colorMap[keyName].Red() {
			t.Fatalf("expected R value %v for key %v, got %v", colorMap[keyName].Red(), keyName, r)
		}
		if g != colorMap[keyName].Green() {
			t.Fatalf("expected G value %v for key %v, got %v", colorMap[keyName].Green(), keyName, g)
		}
		if b != colorMap[keyName].Blue() {
			t.Fatalf("expected B value %v for key %v, got %v", colorMap[keyName].Blue(), keyName, b)
		}
	}
}
