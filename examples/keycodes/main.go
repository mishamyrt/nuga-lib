// Example of printing keycodes with known actions
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib/dump"
	"github.com/mishamyrt/nuga-lib/examples/keycodes/annotation"
	"github.com/mishamyrt/nuga-lib/examples/keycodes/keymap"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/layout"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Too few arguments")
		fmt.Println("Usage: nuga-keycodes <keymap|annotation> <dump_path>")
		os.Exit(1)
	}
	cmd := os.Args[1]
	path := os.Args[2]
	data, err := os.ReadFile(path)
	if err != nil {
		die("Error reading file: %v", err)
	}
	var state dump.State
	err = json.Unmarshal(data, &state)
	if err != nil {
		die("Error unmarshalling: %v", err)
	}
	macCodes := keys.PackKeyCodes(state.Data.Keys.Mac)
	winCodes := keys.PackKeyCodes(state.Data.Keys.Win)
	tpl := layout.GetKeystrokeTemplate(state.Model)
	switch cmd {
	case "keymap":
		keymap.Print(macCodes, tpl, false)
	case "annotation":
		fmt.Println("Mac:")
		annotation.Print(macCodes, tpl)
		fmt.Println("Win:")
		annotation.Print(winCodes, tpl)
	}

}

func die(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
	os.Exit(1)
}
