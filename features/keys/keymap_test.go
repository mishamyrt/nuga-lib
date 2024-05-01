package keys_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/layout"
)

var testTemplate = layout.KeystrokeTemplate{
	layout.KeyM:            0,
	layout.KeyK:            1,
	layout.KeyBrightnessUp: 2,
}

var testKeyMap = keys.KeyMap{
	layout.KeyM: {
		Type: keys.ActionKeystroke,
		Keystroke: &keys.KeystrokeParams{
			Name:      layout.KeyM,
			Modifiers: &keys.Modifiers{Ctrl: true},
		},
	},
	layout.KeyK: {
		Type: keys.ActionKeystroke,
		Keystroke: &keys.KeystrokeParams{
			Name:      layout.KeyK,
			Modifiers: &keys.Modifiers{},
		},
	},
	layout.KeyBrightnessUp: {
		Type: keys.ActionKeystroke,
		Keystroke: &keys.KeystrokeParams{
			Name: layout.KeyBrightnessUp,
		},
	},
}

func TestApply(t *testing.T) {
	tests := []struct {
		expected layout.KeyName
		code     uint32
	}{
		{layout.KeyM, keys.ApplyModifiers(
			layout.Keys[layout.KeyM].Code,
			testKeyMap[layout.KeyM].Keystroke.Modifiers,
		)},
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

func TestParseKeyMap(t *testing.T) {
	values := []uint32{
		keys.ApplyModifiers(layout.Keys[layout.KeyM].Code, &keys.Modifiers{Ctrl: true}),
		layout.Keys[layout.KeyK].Code,
		layout.Keys[layout.KeyBrightnessUp].Code,
	}
	keyMap, err := keys.ParseKeyMap(values, &testTemplate)
	if err != nil {
		t.Fatalf("Parse returned an error: %v", err)
	}

	if len(*keyMap) != len(testKeyMap) {
		t.Errorf("Expected keyMap length to be %d, got %d", len(testKeyMap), len(*keyMap))
	}

	for k, v := range testKeyMap {
		if (*keyMap)[k].Keystroke.Name != v.Keystroke.Name {
			t.Errorf("Expected keyMap[%s].Name to be %#v, got %#v", k, v.Keystroke.Name, (*keyMap)[k].Keystroke.Name)
		}
		if v.Keystroke.Modifiers == nil || (*keyMap)[k].Keystroke.Modifiers == nil {
			if v.Keystroke.Modifiers == nil && (*keyMap)[k].Keystroke.Modifiers == nil {
				continue
			} else {
				t.Errorf(
					"Unexpected one modifier to be nil: %v, %v",
					v.Keystroke.Modifiers,
					(*keyMap)[k].Keystroke.Modifiers,
				)
			}
		}
		if v.Keystroke.Modifiers == nil && (*keyMap)[k].Keystroke.Modifiers != nil {
			t.Errorf("Expected keyMap[%s].Modifiers to be nil, got %#v", k, (*keyMap)[k].Keystroke.Modifiers)
		} else {
			expectedModifiers := *v.Keystroke.Modifiers
			gotModifiers := *((*keyMap)[k].Keystroke.Modifiers)
			if expectedModifiers != gotModifiers {
				t.Errorf("Expected keyMap[%s].Modifiers to be %#v, got %#v", k, expectedModifiers, gotModifiers)
			}
		}

	}
}
