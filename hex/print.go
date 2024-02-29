// Package hex provides functions for formatting hex strings
package hex

import (
	"fmt"
	"io"
	"os"
)

// Out is target for printing
var Out io.Writer = os.Stdout

// PrintBytes prints byte array in hex format
func PrintBytes(b []byte) {
	if len(b) == 0 {
		fmt.Fprint(Out, "[]\n")
		return
	}
	result := FromByte(b[0])
	if len(b) > 1 {
		rest := b[1:]
		for _, v := range rest {
			result += ", " + FromByte(v)
		}
	}
	fmt.Fprintf(Out, "[%s]\n", result)
}
