package layout_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/layout"
)

var testTpl = layout.Template{
	"key1": 1,
	"key2": 2,
	"key3": 3,
}

func TestGetPosition(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     layout.KeyName
		expected uint32
	}{
		{"key1", 1},
		{"key2", 2},
		{"key3", 3},
		{"key4", 0},
	}
	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			if result := testTpl.GetPosition(tt.name); result != tt.expected {
				t.Errorf("GetPosition(%s): expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}

func TestGetKey(t *testing.T) {
	t.Parallel()
	tests := []struct {
		position uint32
		expected layout.KeyName
	}{
		{1, "key1"},
		{2, "key2"},
		{3, "key3"},
		{0, ""},
	}
	for _, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			if result := testTpl.GetKey(tt.position); result != tt.expected {
				t.Errorf("GetKey(%d): expected %s, got %s", tt.position, tt.expected, result)
			}
		})
	}
}
