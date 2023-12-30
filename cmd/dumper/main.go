package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nuga/internal/slices"
	"nuga/pkg/features"
	"nuga/pkg/features/light"
	"nuga/pkg/hid"
	"os"
)

func collectTemplate() features.SimulationTemplate {
	err := hid.Init()
	if err != nil {
		log.Panicf("Error while initializing HID: %v", err)
	}
	handle, err := hid.Open()
	if err != nil {
		log.Panicf("Error while opening device: %v", err)
	}
	f := light.New(handle)
	params, err := f.GetRawEffects()
	if err != nil {
		log.Panicf("Couldn't read params: %v", err)
	}
	info, err := handle.GetInfo()
	if err != nil {
		log.Panicf("Couldn't read product name: %v", err)
	}
	colors, err := f.GetRawColors()
	if err != nil {
		log.Panicf("Couldn't read colors: %v", err)
	}
	return features.SimulationTemplate{
		Name:     info.Name,
		Firmware: info.Firmware,
		Lights: &light.SimulationTemplate{
			Colors: slices.Cast[byte, int](colors),
			Params: slices.Cast[byte, int](params),
		},
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Output path not provided")
		fmt.Println("Usage: dumper <output_file>")
		os.Exit(1)
	}
	outFile := os.Args[1]
	template := collectTemplate()
	content, err := json.Marshal(template)
	if err != nil {
		log.Panicf("Couldn't format dump: %v", err)
	}
	err = os.WriteFile(outFile, []byte(content), 0644)
	if err != nil {
		log.Panicf("Couldn't write file: %v", err)
	}
}
