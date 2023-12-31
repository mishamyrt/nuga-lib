package main

import (
	"fmt"
	"nuga/pkg/features"
)

func describe(t *features.SimulationTemplate) {
	fmt.Printf("Name: %v\n", t.Name)
	fmt.Printf("Firmware version: %v\n", t.Firmware)
}
