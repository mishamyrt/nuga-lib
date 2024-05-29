package keys

import "errors"

// ErrNoTemplate is returned when template is not ready
var ErrNoTemplate = errors.New("connected keyboard template is not ready yet")

// ErrWrongKeyCode is returned when key code is wrong
var ErrWrongKeyCode = errors.New("wrong key code")

// ErrWrongDelay is returned when delay is wrong
var ErrWrongDelay = errors.New("wrong delay value. It must be more or equal 3")

// ErrWrongMacroHeader is returned when macro header is wrong
var ErrWrongMacroHeader = errors.New("wrong macro header. Must be 0x85")

// ErrWrongLength is returned when slice length is wrong
var ErrWrongLength = errors.New("slice should be of length 4")

// ErrKeyNotFound is returned when key not found
var ErrKeyNotFound = errors.New("key not found")
