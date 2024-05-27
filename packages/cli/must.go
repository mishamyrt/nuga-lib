package cli

import (
	"fmt"
)

// Must prints error message and exits the program if err is not nil
func Must(message string, err error) {
	if err != nil {
		fmt.Printf("Tried to %s, but got an error\n", message)
		panic(err)
	}
}
