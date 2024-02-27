// Example of printing keycodes with known actions
package main

import (
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/cmd/keycodes/annotation"
	"github.com/mishamyrt/nuga-lib/cmd/keycodes/keymap"
	"github.com/mishamyrt/nuga-lib/dump"
	"github.com/mishamyrt/nuga-lib/features/keys/layout"
	"github.com/mishamyrt/nuga-lib/hid"
)

func main() {
	err := nuga.Init()
	if err != nil {
		die("Error initializing: %v", err)
	}
	d, err := hid.Open()
	if err != nil {
		die("Error opening device: %v", err)
	}
	state, err := dump.Collect(d)
	if err != nil {
		die("Error collecting state: %v", err)
	}
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments")
		fmt.Println("Usage: nuga-keycodes <keymap>")
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "keymap":
		keymap.Print(state.Keys.Mac, layout.GetTemplate(state.Name))
	case "annotation":
		annotation.Print(state.Keys.Mac, layout.GetTemplate(state.Name))
	}

}

func die(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
	os.Exit(1)
}
