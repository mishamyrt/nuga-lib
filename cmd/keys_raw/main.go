// Example of usage of nuga-lib for raw keys dump
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mishamyrt/nuga-lib"
)

type keysRawDump struct {
	Mac []uint32 `json:"mac"`
	Win []uint32 `json:"win"`
}

func readDump(path string) *keysRawDump {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Can't read dump: %v", err)
		os.Exit(1)
	}
	var dump keysRawDump
	err = json.Unmarshal(data, &dump)
	if err != nil {
		fmt.Printf("Can't unmarshal dump: %v", err)
		os.Exit(1)
	}
	return &dump
}

func writeDump(path string, dump *keysRawDump) {
	data, err := json.Marshal(&dump)
	if err != nil {
		fmt.Printf("Can't marshal json dump: %v", err)
		os.Exit(1)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Printf("Can't write dump: %v", err)
		os.Exit(1)
	}
}

func restoreKeys(dev *nuga.Device, path string) error {
	dump := readDump(path)
	fmt.Println("Restoring mac keys...")
	err := dev.Features.Keys.SetMac(dump.Mac)
	if err != nil {
		return err
	}
	fmt.Println("Restoring win keys...")
	err = dev.Features.Keys.SetWin(dump.Win)
	if err != nil {
		return err
	}
	fmt.Println("Keys is restored")
	return nil
}

func dumpKeys(dev *nuga.Device, path string) error {
	fmt.Println("Reading mac keys...")
	macValues, err := dev.Features.Keys.GetMac()
	if err != nil {
		return err
	}
	fmt.Println("Reading win keys...")
	winValues, err := dev.Features.Keys.GetWin()
	if err != nil {
		return err
	}
	fmt.Println("Saving dump...")
	writeDump(path, &keysRawDump{
		Mac: macValues,
		Win: winValues,
	})
	fmt.Println("Keys is dumped")
	return nil
}

func main() {
	err := nuga.Init()
	if err != nil {
		fmt.Printf("HID initialization failed: %v", err)
		os.Exit(1)
	}
	dev, err := nuga.Open()
	if err != nil {
		fmt.Printf("Can't open device: %v", err)
		os.Exit(1)
	}

	cmd := os.Args[1]
	path := os.Args[2]
	switch cmd {
	case "dump":
		err = dumpKeys(dev, path)
		if err != nil {
			fmt.Printf("Error while dumping keys: %v\n", err)
			os.Exit(1)
		}
		return
	case "restore":
		err = restoreKeys(dev, path)
		if err != nil {
			fmt.Printf("Error while restoring keys: %v\n", err)
			os.Exit(1)
		}
		return
	default:
		fmt.Println("Unknown command")
		fmt.Println("Usage: keys_raw [dump|restore] <file_path>")
		os.Exit(1)
	}
}
