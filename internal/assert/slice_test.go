package assert_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/internal/assert"
)

func TestSliceValue(t *testing.T) {
	t.Parallel()
	err := assert.SliceValue([]int{1, 2, 3}, 2, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = assert.SliceValue([]int{1, 2, 3}, 2, 4)
	if err == nil {
		t.Error("Expected error, but got none")
	}
}

func TestSliceIncludes(t *testing.T) {
	t.Parallel()
	err := assert.SliceIncludes([]int{1, 2, 3}, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = assert.SliceIncludes([]int{1, 2, 3}, 4)
	if err == nil {
		t.Error("Expected error, but got none")
	}
}
