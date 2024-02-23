package layout_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/keys/layout"
)

var testTemplate = layout.Template{
	layout.KeyM:            0,
	layout.KeyK:            1,
	layout.KeyBrightnessUp: 2,
}

var testKeyMap = layout.KeyMap{
	layout.KeyM:            {Name: layout.KeyM, Modifiers: &layout.Modifiers{Ctrl: true}},
	layout.KeyK:            {Name: layout.KeyK, Modifiers: &layout.Modifiers{}},
	layout.KeyBrightnessUp: {Name: layout.KeyBrightnessUp},
}

func TestApply(t *testing.T) {
	tests := []struct {
		expected layout.KeyName
		code     uint32
	}{
		{layout.KeyM, layout.ApplyModifiers(layout.Keys[layout.KeyM].Code, testKeyMap[layout.KeyM].Modifiers)},
		{layout.KeyK, layout.Keys[layout.KeyK].Code},
		{layout.KeyBrightnessUp, layout.Keys[layout.KeyBrightnessUp].Code},
	}
	source := make([]uint32, len(tests))
	err := testKeyMap.Apply(source, &testTemplate)
	if err != nil {
		t.Fatalf("Apply returned an error: %v", err)
	}

	for i, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			if source[i] != tt.code {
				t.Errorf("Expected source[%d] to be %#x, got %#x", i, tt.code, source[i])
			}
		})
	}
}

func TestParse(t *testing.T) {
	values := []uint32{
		layout.ApplyModifiers(layout.Keys[layout.KeyM].Code, &layout.Modifiers{Ctrl: true}),
		layout.Keys[layout.KeyK].Code,
		layout.Keys[layout.KeyBrightnessUp].Code,
	}
	keyMap, err := layout.Parse(values, &testTemplate)
	if err != nil {
		t.Fatalf("Parse returned an error: %v", err)
	}

	if len(*keyMap) != len(testKeyMap) {
		t.Errorf("Expected keyMap length to be %d, got %d", len(testKeyMap), len(*keyMap))
	}

	for k, v := range testKeyMap {
		if (*keyMap)[k].Name != v.Name {
			t.Errorf("Expected keyMap[%s].Name to be %#v, got %#v", k, v.Name, (*keyMap)[k].Name)
		}
		if v.Modifiers == nil || (*keyMap)[k].Modifiers == nil {
			if v.Modifiers == nil && (*keyMap)[k].Modifiers == nil {
				continue
			} else {
				t.Errorf("Unexpected one modifier to be nil: %v, %v", v.Modifiers, (*keyMap)[k].Modifiers)
			}
		}
		expectedModifiers := *v.Modifiers
		gotModifiers := *((*keyMap)[k].Modifiers)
		if expectedModifiers != gotModifiers {
			t.Errorf("Expected keyMap[%s].Modifiers to be %#v, got %#v", k, expectedModifiers, gotModifiers)
		}
	}
}
