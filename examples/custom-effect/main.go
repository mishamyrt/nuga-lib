// custom-tpl-builder
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/examples/custom-effect/gostdin"
	"github.com/mishamyrt/nuga-lib/examples/custom-effect/wizard"
	"github.com/mishamyrt/nuga-lib/features/light"
)

func printDescription() {
	fmt.Println("Custom backlight template wizard")
	fmt.Println("The utility will sequentially set the color for every key.")
	fmt.Println("If the button is illuminated - write its name.")
	fmt.Println("If not, just press Enter.")
}

func recordMap(dev *nuga.Device) {
	wizard, err := wizard.New(dev)
	if err != nil {
		die("Error creating wizard: %v", err)
	}

	closeChan := make(chan bool)
	err = wizard.Start()
	if err != nil {
		fmt.Printf("Error starting wizard: %v", err)
		os.Exit(1)
	}

	gostdin.OnSignal(syscall.SIGINT, func() {
		fmt.Println("\nReverting state...")
		err := wizard.Restore()
		if err != nil {
			die("Error restoring state: %v", err)
		}
		data, err := json.Marshal(wizard.Keys)
		if err != nil {
			fmt.Printf("Error marshalling: %v", err)
		} else {
			fmt.Println(string(data))
		}
		close(closeChan)
		os.Exit(1)
	})

	printDescription()

	gostdin.OnInput("Key:", func(s string) {
		err := wizard.HandleKey(s)
		if err != nil {
			fmt.Println("\nKey name is invalid")
			return
		}
		err = wizard.Next()
		if err != nil {
			fmt.Printf("\nError setting next state: %v", err)
		}
	}, closeChan)
}

func saveDump(d *nuga.Device, path string) {
	state, err := d.Features.Light.GetCustomEffect()
	if err != nil {
		die("Error getting state: %v", err)
	}
	data, err := json.Marshal(&state)
	if err != nil {
		die("Error marshalling: %v", err)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		die("Error writing file: %v", err)
	}
}

func restoreDump(d *nuga.Device, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		die("Error reading file: %v", err)
	}
	var state light.CustomEffectMap
	err = json.Unmarshal(data, &state)
	if err != nil {
		die("Error unmarshalling: %v", err)
	}
	err = d.Features.Light.SetCustomEffect(&state)
	if err != nil {
		die("Error restoring state: %v", err)
	}
}

func main() {
	err := nuga.Init()
	if err != nil {
		die("Error initializing nuga: %v", err)
	}
	dev, err := nuga.Open()
	if err != nil {
		die("Error opening device: %v", err)
	}
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments")
		fmt.Println("Usage: nuga-custom-effect <record|restore|dump> [<path>]")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "record":
		recordMap(dev)
	case "restore":
		restoreDump(dev, os.Args[2])
	case "dump":
		saveDump(dev, os.Args[2])
	}
}

func die(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
	os.Exit(1)
}
