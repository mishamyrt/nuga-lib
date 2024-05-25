// Package assert contains assertion functions
package assert

import (
	"fmt"

	"github.com/pkg/errors"
)

// SliceValue returns error if value at position is not equal to target value
func SliceValue[T comparable](haystack []T, position int, needle T) error {
	if haystack[position] == needle {
		return nil
	}
	return errors.Wrap(ErrUnexpectedValue, fmt.Sprintf("%v at %v", needle, position))
}

// SliceIncludes returns error if target value is not found in slice
func SliceIncludes[T comparable](haystack []T, needle T) error {
	for _, v := range haystack {
		if v == needle {
			return nil
		}
	}
	return errors.Wrap(ErrNotFound, fmt.Sprintf("%v", needle))
}
