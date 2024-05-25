package light

import "errors"

// ErrLightsOff is returned when the user tries to perform an action with the light off.
var ErrLightsOff = errors.New("this action cannot be performed when the lights are off")

// ErrNotSupported is returned when the user tries to change a property that is not supported by the effect.
var ErrNotSupported = errors.New("current light mode does not support changing this property")

// ErrOutOfRange is returned when the value transmitted by the user is outside the permitted limit.
var ErrOutOfRange = errors.New("passed value is outside the permitted limits")

// ErrNoTemplate is returned when template is not ready
var ErrNoTemplate = errors.New("keyboard is not supports custom colors yet")

// ErrInvalidColorFormat is returned when hex color string format is unexpected
var ErrInvalidColorFormat = errors.New("invalid hex color string format")
