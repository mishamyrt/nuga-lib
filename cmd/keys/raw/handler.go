// Package raw provides helpers for working with raw keys layout
package raw

import (
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/cmd/keys/jsonfile"
)

// Handler handles raw keys layout
type Handler struct{}

// New creates raw keys layout handler
func New() *Handler {
	return &Handler{}
}

// Dump collects raw keys from device and dumps it to file
func (h *Handler) Dump(dev *nuga.Device, path string) {
	fmt.Println("Collecting keys...")
	dump, err := collectKeys(dev)
	if err != nil {
		fmt.Printf("Can't collect keys: %v", err)
		os.Exit(1)
	}
	fmt.Println("Writing dump...")
	err = jsonfile.Write(path, dump)
	if err != nil {
		fmt.Printf("Can't write dump: %v", err)
		os.Exit(1)
	}
	fmt.Println("Keys successfully dumped")
}

// Restore restores raw keys from file
func (h *Handler) Restore(dev *nuga.Device, path string) {
	fmt.Println("Reading dump...")
	dump, err := jsonfile.Read[keysDump](path)
	if err != nil {
		fmt.Printf("Can't read dump: %v", err)
		os.Exit(1)
	}

	fmt.Println("Restoring keys...")
	err = applyDump(dev, dump)
	if err != nil {
		fmt.Printf("Can't restore keys: %v", err)
		os.Exit(1)
	}
	fmt.Println("Keys is restored")
}

func applyDump(dev *nuga.Device, dump *keysDump) error {
	err := dev.Features.Keys.SetMacCodes(dump.Mac)
	if err != nil {
		return err
	}
	err = dev.Features.Keys.SetWinCodes(dump.Win)
	if err != nil {
		return err
	}
	return nil
}

func collectKeys(dev *nuga.Device) (*keysDump, error) {
	macValues, err := dev.Features.Keys.GetMacCodes()
	if err != nil {
		return nil, err
	}
	winValues, err := dev.Features.Keys.GetWinCodes()
	if err != nil {
		return nil, err
	}
	return &keysDump{
		Mac: macValues,
		Win: winValues,
	}, nil
}
