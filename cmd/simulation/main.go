// Simulation template dumper
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/internal/slices"
	"github.com/mishamyrt/nuga-lib/keys"
	"github.com/mishamyrt/nuga-lib/light"
)

const usage = "Usage: nuga-simulation [read|dump] <file_path>"

func readTemplate(filePath string) (*nuga.SimulationTemplate, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var template nuga.SimulationTemplate
	err = json.Unmarshal(content, &template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func describe(t *nuga.SimulationTemplate) {
	fmt.Printf("Name: %v\n", t.Name)
	fmt.Printf("Firmware version: %v\n", t.Firmware)
}

func collectTemplate() (*nuga.SimulationTemplate, error) {
	err := nuga.Init()
	if err != nil {
		return nil, err
	}
	handle, err := hid.Open()
	if err != nil {
		return nil, err
	}
	lightsFeature := light.New(handle)
	keysFeature := keys.New(handle, handle.Info.Model)
	params, err := lightsFeature.GetRawEffects()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	colors, err := lightsFeature.GetRawColors()
	if err != nil {
		return nil, err
	}
	macCodes, err := keysFeature.GetMacCodes()
	if err != nil {
		return nil, err
	}
	winCode, err := keysFeature.GetWinCodes()
	if err != nil {
		return nil, err
	}
	return &nuga.SimulationTemplate{
		Name:     handle.Info.Model,
		Firmware: handle.Info.Firmware,
		Lights: &light.SimulationTemplate{
			Colors: slices.Cast[byte, int](colors),
			Params: slices.Cast[byte, int](params),
		},
		Keys: &keys.SimulationTemplate{Mac: macCodes, Win: winCode},
	}, nil
}

func dumpTemplate(filePath string) error {
	template, err := collectTemplate()
	if err != nil {
		return err
	}
	content, err := json.Marshal(template)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, []byte(content), 0644)
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
	if err != nil {
		fmt.Printf("HID initialization failed: %v", err)
		os.Exit(1)
	}
	switch command {
	case "read":
		template, err := readTemplate(filePath)
		if err == nil {
			describe(template)
			return
		}
	case "dump":
		err = dumpTemplate(filePath)
		if err == nil {
			fmt.Println("Successfully dumped")
			return
		}
	default:
		fmt.Printf("Unknown command '%v'\n", command)
		fmt.Println(usage)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error while handling command: %v\n", err)
	}
	err = nuga.Exit()
	if err != nil {
		fmt.Printf("HID exiting failed: %v", err)
		os.Exit(1)
	}
}
