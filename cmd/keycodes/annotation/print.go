// Package annotation provides helpers for working with keycodes
package annotation

import (
	"fmt"
	"strings"

	"github.com/mishamyrt/nuga-lib/features/keys/layout"
)

// Print keycodes with known actions
func Print(values []uint32, tpl *layout.Template) {
	lines := make([]string, 0, len(*tpl))
	for i, value := range values {
		keyName := layout.FindKeyName(value)
		displayName := ""
		if keyName != "none" {
			displayName = string(keyName)
		}
		line := fmt.Sprintf("%-5d 0x%08x   %s", i, value, displayName)
		lines = append(lines, line)
	}
	fmt.Println(strings.Join(lines, "\n"))
}
