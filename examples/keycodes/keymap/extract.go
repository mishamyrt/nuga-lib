package keymap

import (
	"strconv"
)

func extractFunctionKeyNumber(s string, prefixLength int) int {
	if num, err := strconv.Atoi(s[prefixLength:]); err == nil {
		return num
	}
	return 0
}

func extractFKeyNumber(s string) int {
	return extractFunctionKeyNumber(s, 1)
}

func extractFNKeyNumber(s string) int {
	return extractFunctionKeyNumber(s, 4)
}

func extractNumKeyNumber(s string) int {
	return extractFunctionKeyNumber(s, 3)
}
