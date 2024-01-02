// Keyboard describe example app
package main

import (
	"fmt"
	"github.com/mishamyrt/nuga-lib"
)

func collectFeatures(c *nuga.Capability) []string {
	names := make([]string, 0)
	if c.Has(nuga.BacklightCapability) {
		names = append(names, "backlight")
	}
	if c.Has(nuga.HalolightCapability) {
		names = append(names, "halolight")
	}
	if c.Has(nuga.SidelightCapability) {
		names = append(names, "sidelight")
	}
	return names
}

func describe(d *nuga.Device, supports []string) {
	fmt.Printf("Model: %v\n", d.Name)
	fmt.Printf("Firmware: %v\n", d.Firmware)
	fmt.Printf("Path: %v\n", d.Path)
	fmt.Println("Supported features:")
	for _, name := range supports {
		fmt.Printf("- %v\n", name)
	}
}

func main() {
	dev, err := nuga.Open()
	if err != nil {
		fmt.Printf("Can't open device: %v", err)
	}
	supports := collectFeatures(dev.Capabilities)
	describe(dev, supports)
}
