// Example app for dumping and restoring device state
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/device"
	"github.com/mishamyrt/nuga-lib/dump"
	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/layout"
	"github.com/mishamyrt/nuga-lib/packages/cli"
)

var app = cli.App{
	Name: "nuga-dump",
	Help: "Utility for workings with keyboard dumps",
}

var (
	inputPath  string
	outputPath string

	cmdSave = cli.Command{
		Name:  "save",
		Help:  "Save device state to file",
		Usage: "<output_path>",
		Args:  cli.Args{&outputPath},
		Run: func(_ []string) {
			dev, err := nuga.Open()
			cli.Must("open device", err)
			state, err := dump.Collect(dev.Handle, dev.Name)
			cli.Must("collect device state", err)
			data, err := json.Marshal(&state)
			cli.Must("marshal device state", err)
			_, err = dev.Features.Light.GetEffects()
			cli.Must("get effects", err)
			cli.Must("write file", os.WriteFile(outputPath, data, 0644))
		},
	}

	cmdRestore = cli.Command{
		Name:  "restore",
		Help:  "Restore device state from file",
		Usage: "<input_path>",
		Args:  cli.Args{&inputPath},
		Run: func(_ []string) {
			d, err := hid.Open()
			cli.Must("open device", err)
			state, err := readStateDump(inputPath)
			cli.Must("read state file", err)
			cli.Must("restore device state", dump.Restore(d, state))
		},
	}

	cmdValidate = cli.Command{
		Name:  "validate",
		Help:  "Validate device state dump",
		Usage: "<input_path>",
		Args:  cli.Args{&inputPath},
		Run: func(_ []string) {
			state, err := readStateDump(inputPath)
			mustHandleDump("read", err)
			{
				_, err := light.ParseBacklightColors(state.Data.Lights.Colors)
				mustHandleField("lights.colors", err)
				_, err = light.ParseEffects(state.Data.Lights.Params)
				mustHandleField("lights.params", err)
				tpl := layout.GetBacklightTemplate(state.Model)
				if tpl != nil {
					_, err = light.ParseCustomEffect(state.Data.Lights.CustomEffect, tpl)
					mustHandleField("lights.custom_effect", err)
				}
			}
			{
				_, err := keys.ParseMacros(state.Data.Keys.Macros)
				mustHandleField("keys.macros", err)
				tpl := layout.GetKeystrokeTemplate(state.Model)
				if tpl != nil {
					_, err = keys.ParseKeyMap(state.Data.Keys.Mac, tpl)
					mustHandleField("keys.mac", err)
					_, err = keys.ParseKeyMap(state.Data.Keys.Win, tpl)
					mustHandleField("keys.win", err)
				} else {
					fmt.Println("Keystroke template is not defined, skipping keymap validation")
				}
			}
			fmt.Println("🟩 Dump is valid")
		},
	}

	cmdConvert = cli.Command{
		Name:  "convert",
		Help:  "Convert old JSON dump format to new nugafile",
		Usage: "<input_path> <output_path>",
		Args:  cli.Args{&inputPath, &outputPath},
		Run: func(_ []string) {
			data, err := os.ReadFile(inputPath)
			cli.Must("read file", err)

			var originalDump jsonDump
			err = json.Unmarshal(data, &originalDump)
			cli.Must("unmarshal file", err)

			var nugafile dump.State
			nugafile.Model = device.Model(originalDump.Name)
			nugafile.Data.Lights = &light.StateData{
				Colors:       originalDump.Lights.Colors[7:1031],
				Params:       originalDump.Lights.Params[15:138],
				CustomEffect: make([]byte, 1024),
			}
			nugafile.Data.Keys = &keys.StateData{
				Mac:    keys.UnpackKeyCodes(originalDump.Keys.Mac),
				Win:    keys.UnpackKeyCodes(originalDump.Keys.Win),
				Macros: make([]byte, 1024),
			}
			data, err = json.Marshal(originalDump)
			cli.Must("marshal file", err)
			err = os.WriteFile(outputPath, data, 0644)
			cli.Must("write file", err)
		},
	}
)

func main() {
	cli.Must("initialize connection", nuga.Init())
	app.RunWith(cmdSave, cmdRestore, cmdValidate, cmdConvert)
	cli.Must("close connection", nuga.Exit())
}

type jsonDump struct {
	Name     string `json:"name"`
	Firmware string `json:"firmware"`
	Lights   struct {
		Colors []byte `json:"colors"`
		Params []byte `json:"params"`
	}
	Keys struct {
		Mac []uint32 `json:"mac"`
		Win []uint32 `json:"win"`
	}
}

func readStateDump(path string) (*dump.State, error) {
	var state dump.State
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func mustHandleDump(m string, err error) {
	if err == nil {
		return
	}
	message := fmt.Sprintln("🟥 " + cli.Red("Dump is not valid"))
	if m != "" {
		message += fmt.Sprintln(m)
	}
	message += fmt.Sprintln(err)
	fmt.Println(message)
	os.Exit(1)
}

func mustHandleField(k string, err error) {
	if err != nil {
		m := fmt.Sprintf("Problem with field %s: %v\n", k, err)
		mustHandleDump(m, err)
	}
}
