package main

import (
	"fmt"
	"github.com/mishamyrt/nuga-lib"
)

func describe(t *nuga.SimulationTemplate) {
	fmt.Printf("Name: %v\n", t.Name)
	fmt.Printf("Firmware version: %v\n", t.Firmware)
}
