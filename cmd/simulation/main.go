// Simulation template dumper
package main

import (
	"fmt"
	"os"
)

const usage = "Usage: nuga-simulation [read|dump] <file_path>"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Too few arguments")
		fmt.Println(usage)
		os.Exit(1)
	}
	command := os.Args[1]
	filePath := os.Args[2]
	var err error
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
}
