// Keyboard describe example app
package main

import (
	"fmt"
	"nuga/internal/bit"
	"nuga/pkg/device"
)

func collectSupports(supports bit.Mask) []string {
	names := make([]string, 0)
	if supports.IsSet(device.SupportsBacklight) {
		names = append(names, "backlight")
	}
	if supports.IsSet(device.SupportsHalolight) {
		names = append(names, "halolight")
	}
	if supports.IsSet(device.SupportsSidelight) {
		names = append(names, "sidelight")
	}
	return names
}

func describe(d *device.Device, supports []string) {
	fmt.Printf("Model: %v\n", d.Name)
	fmt.Printf("Firmware: %v\n", d.Firmware)
	fmt.Printf("Path: %v\n", d.Path)
	fmt.Println("Supported features:")
	for _, name := range supports {
		fmt.Printf("- %v\n", name)
	}
}

func main() {
	dev, err := device.Open()
	if err != nil {
		fmt.Printf("Can't open device: %v", err)
	}
	supports := collectSupports(dev.Supports)
	describe(dev, supports)
}
