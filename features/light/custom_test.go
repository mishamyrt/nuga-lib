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

var colorMap = light.CustomBacklightMap{
	layout.KeyM: {R: 0xFF, G: 0x00, B: 0x00},
	layout.KeyY: {R: 0x00, G: 0xFF, B: 0x00},
	layout.KeyR: {R: 0x00, G: 0x00, B: 0xFF},
	layout.KeyT: {R: 0xFF, G: 0xFF, B: 0xFF},
}

func TestParseCustomBacklight(t *testing.T) {
	t.Parallel()
	payload := make([]byte, 1050)
	payload[0] = 0x89
	for keyName, position := range tpl {
		payload[7+position] = colorMap[keyName].R
		payload[7+position+126] = colorMap[keyName].G
		payload[7+position+252] = colorMap[keyName].B
	}
	customMap, err := light.ParseCustomBacklight(payload, &tpl)
	if err != nil {
		t.Fatal("unexpected error:", err)
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
	payload := colorMap.Bytes(&tpl)
	if len(payload) != 1024 {
		t.Fatalf("expected 1024 bytes, got %d", len(payload))
	}
	for keyName, position := range tpl {
		r := payload[position]
		g := payload[position+126]
		b := payload[position+252]
		if r != colorMap[keyName].R {
			t.Fatalf("expected R value %v for key %v, got %v", colorMap[keyName].R, keyName, r)
		}
		if g != colorMap[keyName].G {
			t.Fatalf("expected G value %v for key %v, got %v", colorMap[keyName].G, keyName, g)
		}
		if b != colorMap[keyName].B {
			t.Fatalf("expected B value %v for key %v, got %v", colorMap[keyName].B, keyName, b)
		}
	}
}
