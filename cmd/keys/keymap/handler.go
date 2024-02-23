// Package keymap provides helpers for working with keys layout
package keymap

import (
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/cmd/keys/jsonfile"
)

// Handler handles keys layout
type Handler struct {
}

// New creates keys layout handler
func New() *Handler {
	return &Handler{}
}

// Dump collects keys from device and dumps it to file
func (h *Handler) Dump(dev *nuga.Device, path string) {
	macKeys, err := dev.Features.Keys.GetMac()
	if err != nil {
		fmt.Printf("Can't read mac layout: %v", err)
		os.Exit(1)
	}
	winKeys, err := dev.Features.Keys.GetWin()
	if err != nil {
		fmt.Printf("Can't read win layout: %v", err)
		os.Exit(1)
	}
	err = jsonfile.Write(path, &Dump{
		Mac: macKeys,
		Win: winKeys,
	})
	if err != nil {
		fmt.Printf("Can't write dump: %v", err)
		os.Exit(1)
	}
}

// Restore restores keys from file
func (h *Handler) Restore(dev *nuga.Device, path string) {
	dump, err := jsonfile.Read[Dump](path)
	if err != nil {
		fmt.Printf("Can't read dump: %v", err)
		os.Exit(1)
	}
	err = dev.Features.Keys.SetMac(dump.Mac)
	if err != nil {
		fmt.Printf("Can't set mac layout: %v", err)
		os.Exit(1)
	}
	err = dev.Features.Keys.SetWin(dump.Win)
	if err != nil {
		fmt.Printf("Can't set win layout: %v", err)
		os.Exit(1)
	}
}
