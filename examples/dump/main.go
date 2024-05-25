// Example app for dumping and restoring device state
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/dump"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/internal/hex"
)

const usage = "Usage: nuga-dump [restore|dump|describe] <file_path>"

func saveDump(d *hid.Device, path string) {
	state, err := dump.Collect(d)
	must("collect device state", err)
	data, err := json.Marshal(&state)
	must("marshal device state", err)
	err = os.WriteFile(path, data, 0644)
	must("write file", err)
}

func restoreDump(d *hid.Device, path string) {
	data, err := os.ReadFile(path)
	must("read file", err)
	var state dump.State
	err = json.Unmarshal(data, &state)
	must("unmarshal file", err)
	err = dump.Restore(d, &state)
	must("restore device state", err)
}

func describeDumb(path string) {
	data, err := os.ReadFile(path)
	must("read file", err)
	var state dump.State
	err = json.Unmarshal(data, &state)
	must("unmarshal file", err)
	fmt.Printf("Model: %v\nFirmware: %v\n", state.Model, state.Firmware)
	fmt.Println("Lights ")
	fmt.Println("- Colors")
	hex.PrintBytes(state.Data.Lights.Colors)
	fmt.Println("- Params")
	hex.PrintBytes(state.Data.Lights.Params)
	fmt.Println("- Custom effect")
	hex.PrintBytes(state.Data.Lights.CustomEffect)
	fmt.Println("Keys")
	fmt.Println("- Mac")
	// hex.PrintBytes(state.State.Keys.Mac)
	fmt.Println("- Win")
	fmt.Println(len(keys.PackKeyCodes(state.Data.Keys.Win)))
	// hex.PrintBytes(state.Data.Keys.Win)
	fmt.Println("- Macros")
	// fmt.Println(len(state.Data.Keys.Macros))
	// hex.PrintBytes(state.State.Keys.Macros)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Too few arguments")
		fmt.Println(usage)
		os.Exit(1)
	}
	command := os.Args[1]
	filePath := os.Args[2]
	err := nuga.Init()
	must("initialize connection", err)
	handle, err := hid.Open()
	must("open device", err)
	switch command {
	case "restore":
		restoreDump(handle, filePath)
	case "dump":
		saveDump(handle, filePath)
	case "describe":
		describeDumb(filePath)
	}
}

func must(message string, err error) {
	if err != nil {
		fmt.Printf("app must %s, but got error: %v\n", message, err.Error())
		os.Exit(1)
	}
}
