// Example app for working with keyboard light effects
package main

import (
	"fmt"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/packages/cli"
)

var app = cli.App{
	Name: "nuga-effects",
	Help: "Utility for workings with keyboard light effects",
}

var (
	cmdDescribe = cli.Command{
		Name: "describe",
		Help: "Print current keyboard lights effects",
		Run: func(_ []string) {
			dev, err := nuga.Open()
			cli.Must("open device", err)
			effects, err := dev.Features.Light.GetEffects()
			cli.Must("get effects", err)
			fmt.Println("Backlight")
			fmt.Printf("  Mode: %v\n", effects.Backlight.Mode.Name)
			fmt.Printf("  Color: %v\n", effects.Backlight.CurrentParams().Color)
			fmt.Printf("  Brightness: %v\n", effects.Backlight.CurrentParams().Brightness)
			fmt.Printf("  Speed: %v\n", effects.Backlight.CurrentParams().Speed)
			fmt.Println("Halo")
			fmt.Printf("  Mode: %v\n", effects.Halo.Mode.Name)
			fmt.Printf("  Color: %v\n", effects.Halo.Color)
			fmt.Printf("  Brightness: %v\n", effects.Halo.Brightness)
			fmt.Printf("  Speed: %v\n", effects.Halo.Speed)
			fmt.Println("Sidelight")
			fmt.Printf("  Mode: %v\n", effects.Sidelight.Mode.Name)
			fmt.Printf("  Color: %v\n", effects.Sidelight.Color)
			fmt.Printf("  Brightness: %v\n", effects.Sidelight.Brightness)
			fmt.Printf("  Speed: %v\n", effects.Sidelight.Speed)
		},
	}
)

func init() {
	err := nuga.Init()
	cli.Must("initialize CLI", err)
}

func main() {
	cli.Must("initialize connection", nuga.Init())
	app.RunWith(cmdDescribe)
	cli.Must("close connection", nuga.Exit())
}
