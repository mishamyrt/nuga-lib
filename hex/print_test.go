package hex_test

import (
	"bytes"
	"testing"

	"github.com/mishamyrt/nuga-lib/hex"
)

func TestPrintBytes(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"empty", []byte{}, "[]\n"},
		{"single", []byte{0x00}, "[00]\n"},
		{"multiple", []byte{0x00, 0x01, 0xFF}, "[00, 01, ff]\n"},
	}

	out := bytes.NewBuffer([]byte{})
	hex.Out = out

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			hex.PrintBytes(c.input)
			result := out.String()
			if result != c.expected {
				t.Errorf("Expected PrintBytes(%v) to be %v, got %v", c.input, c.expected, result)
			}
			out.Reset()
		})
	}
}
