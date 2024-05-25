package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib/features/keys"
	"github.com/mishamyrt/nuga-lib/internal/hex"
)

// {
// 	"model": "Halo75",
// 	"firmware": "1.0.5",
// 	"state": {
// 	  "lights": {
// 		"colors": "",
// 		"effects": "",
// 		"custom_effect": ""
// 	  },
// 	  "keys": {
// 		"mac": "",
// 		"win": "",
// 		"macros": ""
// 	  }
// 	}
//   }

type NewDump struct {
	Model    string `json:"model"`
	Firmware string `json:"firmware"`
	State    struct {
		Lights struct {
			Colors []byte `json:"colors"`
			Params []byte `json:"effects"`
			Custom []byte `json:"custom_effect"`
		} `json:"lights"`
		Keys struct {
			Mac    []byte `json:"mac"`
			Win    []byte `json:"win"`
			Macros []byte `json:"macros"`
		} `json:"keys"`
	} `json:"state"`
}

type OldDump struct {
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

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Too few arguments")
		fmt.Println("Usage: nuga-cc <from> <to>")
		os.Exit(1)
	}

	path := os.Args[1]
	data, err := os.ReadFile(path)
	must("read file", err)

	var oldDump OldDump
	err = json.Unmarshal(data, &oldDump)
	must("unmarshal file", err)

	var newDump NewDump
	newDump.Model = oldDump.Name
	newDump.Firmware = oldDump.Firmware
	newDump.State.Lights.Colors = oldDump.Lights.Colors[7:1031]
	newDump.State.Lights.Params = oldDump.Lights.Params[15:138]
	newDump.State.Lights.Custom = make([]byte, 1024)
	newDump.State.Keys.Mac = keys.UnpackKeyCodes(oldDump.Keys.Mac)
	newDump.State.Keys.Win = keys.UnpackKeyCodes(oldDump.Keys.Win)
	newDump.State.Keys.Macros = make([]byte, 1024)

	hex.PrintBytes(newDump.State.Lights.Params)
	// fmt.Println(len(newDump.State.Lights.Params))
	data, err = json.Marshal(newDump)
	must("marshal file", err)

	err = os.WriteFile(os.Args[2], data, 0644)
	must("write file", err)
}

func must(message string, err error) {
	if err != nil {
		fmt.Printf("app must %s, but got error: %v\n", message, err.Error())
		os.Exit(1)
	}
}
