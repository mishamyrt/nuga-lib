package keymap

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mishamyrt/nuga-lib/features/keys/layout"
)

// Custom comparison logic for sorting
func sortKeys(keys []string) {
	sort.SliceStable(keys, func(i, j int) bool {
		valueI := keys[i]
		valueJ := keys[j]
		catI := categorize(keys[i])
		catJ := categorize(keys[j])

		if catI != catJ {
			return catI < catJ
		}

		switch catI {
		case categoryNum:
			return extractNumKeyNumber(valueI) < extractNumKeyNumber(valueJ)
		case categoryF:
			return extractFKeyNumber(valueI) < extractFKeyNumber(valueJ)
		case categoryFnF:
			return extractFNKeyNumber(valueI) < extractFNKeyNumber(valueJ)
		default:
			return valueI < valueJ
		}
	})
}

// Print keymap
func Print(values []uint32, tpl *layout.Template) {
	keys := make([]string, 0, len(*tpl))
	for key := range *tpl {
		keys = append(keys, string(key))
	}
	sortKeys(keys)
	lines := make([]string, 0, len(*tpl))
	for _, key := range keys {
		position := tpl.GetPosition(layout.KeyName(key))
		line := fmt.Sprintf("%-15s %-5d 0x%08x", key, position+1, values[position])
		lines = append(lines, line)
	}

	fmt.Println(strings.Join(lines, "\n"))
}
