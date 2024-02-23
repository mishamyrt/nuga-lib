// Package keymap provides helpers for working with keys layout
package keymap

import "github.com/mishamyrt/nuga-lib/keys/layout"

// Dump represents keys layout dump
type Dump struct {
	Mac *layout.KeyMap `json:"mac"`
	Win *layout.KeyMap `json:"win"`
}
