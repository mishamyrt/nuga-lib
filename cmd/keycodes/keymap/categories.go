// Package keymap provides helpers for working with keymaps
package keymap

import "strings"

type category int

const (
	categoryLetter    category = iota
	categoryNum       category = iota
	categoryF         category = iota
	categoryFnF       category = iota
	categorizeUnknown category = iota
)

func categorize(item string) category {
	if len(item) == 1 && item[0] >= 'a' && item[0] <= 'z' {
		return categoryLetter
	}
	if strings.HasPrefix(item, "num") && len(item) <= 5 {
		return categoryNum
	}
	if strings.HasPrefix(item, "fn_f") {
		return categoryFnF
	}
	if strings.HasPrefix(item, "f") {
		return categoryF
	}
	return categorizeUnknown
}
