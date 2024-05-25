package assert

import "errors"

// ErrUnexpectedValue is returned when an unexpected value is found at a position
var ErrUnexpectedValue = errors.New("unexpected value")

// ErrNotFound is returned when an expected value is not found
var ErrNotFound = errors.New("not found")
