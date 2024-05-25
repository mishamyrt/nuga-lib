package keys_test

import (
	"fmt"
	"testing"

	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/layout"
)

func ref[T any](v T) *T {
	return &v
}

var testTemplate = layout.Template{
	layout.KeyM:            0,
	layout.KeyK:            1,
	layout.KeyBrightnessUp: 2,
	layout.KeyY:            3,
	layout.KeyR:            4,
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
	layout.KeyY: {
		Type:       keys.ActionMacro,
		MacroIndex: ref(uint8(1)),
	},
	layout.KeyR: {
		Type: keys.ActionNone,
	},
}

func TestBytes(t *testing.T) {
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
		{layout.KeyY, keys.IndexToMacro(1)},
		{layout.KeyR, layout.Keys[layout.KeyNone].Code},
	}
	raw := testKeyMap.Bytes(&testTemplate)
	codes := keys.PackKeyCodes(raw)

	if len(codes) < len(testTemplate) {
		t.Errorf("Expected codes length to be more or equal %d, got %d", len(testTemplate), len(codes))
	}

	for i, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			if codes[i] != tt.code {
				t.Errorf("Expected source[%d] to be %#x, got %#x", i, tt.code, codes[i])
			}
		})
	}
}

func TestParseKeyMap(t *testing.T) {
	values := keys.UnpackKeyCodes([]uint32{
		keys.ApplyModifiers(layout.Keys[layout.KeyM].Code, &keys.Modifiers{Ctrl: true}),
		layout.Keys[layout.KeyK].Code,
		layout.Keys[layout.KeyBrightnessUp].Code,
		keys.IndexToMacro(1),
		layout.Keys[layout.KeyNone].Code,
	})
	keyMap, err := keys.ParseKeyMap(values, &testTemplate)
	if err != nil {
		t.Fatalf("Parse returned an error: %v", err)
	}

	if len(*keyMap) != len(testKeyMap) {
		t.Errorf("Expected keyMap length to be %d, got %d", len(testKeyMap), len(*keyMap))
	}
	for k, v := range testKeyMap {
		switch v.Type {
		case keys.ActionKeystroke:
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
		case keys.ActionMacro:
			if (*keyMap)[k].MacroIndex == nil || *(*keyMap)[k].MacroIndex != 1 {
				t.Errorf("Expected keyMap[%s].MacroIndex to be 1, got %#v", k, (*keyMap)[k].MacroIndex)
			}
		case keys.ActionNone:
			if (*keyMap)[k].Type != keys.ActionNone {
				fmt.Println(*(*keyMap)[k].Keystroke)
				t.Errorf("Expected keyMap[%s].Type to be ActionNone, got %#v", k, (*keyMap)[k].Type)
			}
		}

	}
}
