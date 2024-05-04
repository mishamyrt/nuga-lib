// Package hid contains utils for working with human interface devices
package hid

import (
	"errors"
	"fmt"
)

// ErrNotFound is returned when the device wasn't found
var ErrNotFound = errors.New("keyboard is not found")

// ErrCountMismatch is returned when transmitted number of bytes is not expected
var ErrCountMismatch = errors.New("transmitted number of bytes is not expected")

// NewErrCountMismatch creates a byte count mismatch error.
func NewErrCountMismatch(expected, actual int) error {
	return fmt.Errorf("%w: expected %d, actual %d", ErrCountMismatch, expected, actual)
}
