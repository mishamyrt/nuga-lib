// Example of usage of nuga-lib for raw keys dump
package main

import (
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/cmd/keys/keymap"
	"github.com/mishamyrt/nuga-lib/cmd/keys/raw"
)

// DumpHandler is an interface for dumping and restoring keys
type DumpHandler interface {
	Dump(dev *nuga.Device, path string)
	Restore(dev *nuga.Device, path string)
}

const usage = "Usage: keys [raw] <dump|restore> <file_path>"

func main() {
	err := nuga.Init()
	if err != nil {
		fmt.Printf("HID initialization failed: %v\n", err)
		os.Exit(1)
	}
	dev, err := nuga.Open()
	if err != nil {
		fmt.Printf("Can't open device: %v\n", err)
		os.Exit(1)
	}

	cmd := os.Args[1]
	mode := os.Args[2]
	path := os.Args[3]

	var handler DumpHandler

	switch mode {
	case "raw":
		handler = raw.New()
	case "map":
		handler = keymap.New()
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
		fmt.Println(usage)
		os.Exit(1)
	}

	switch cmd {
	case "dump":
		handler.Dump(dev, path)
		return
	case "restore":
		handler.Restore(dev, path)
		return
	default:
		fmt.Printf("Unknown mode: %s\n", cmd)
		fmt.Println(usage)
		os.Exit(1)
	}
}
