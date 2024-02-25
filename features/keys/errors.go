package keys

import "errors"

// ErrNoTemplate is returned when template is not ready
var ErrNoTemplate = errors.New("connected keyboard template is not ready yet")
