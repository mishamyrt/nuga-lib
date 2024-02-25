// Example app for dumping and restoring device state
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/dump"
	"github.com/mishamyrt/nuga-lib/hid"
)

const usage = "Usage: nuga-dump [restore|dump] <file_path>"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Too few arguments")
		fmt.Println(usage)
		os.Exit(1)
	}
	command := os.Args[1]
	filePath := os.Args[2]
	err := nuga.Init()
	if err != nil {
		die("Error initializing: %v", err)
	}
	handle, err := hid.Open()
	if err != nil {
		die("Error opening device: %v", err)
	}
	switch command {
	case "restore":
		err = restoreDump(handle, filePath)
	case "dump":
		err = saveDump(handle, filePath)
	}
	if err != nil {
		die("Error: %v", err)
	}
}

func saveDump(d *hid.Device, path string) error {
	state, err := dump.Collect(d)
	if err != nil {
		return err
	}
	data, err := json.Marshal(&state)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func restoreDump(d *hid.Device, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var state dump.State
	err = json.Unmarshal(data, &state)
	if err != nil {
		return err
	}
	return dump.Restore(d, &state)
}

func die(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
	os.Exit(1)
}
