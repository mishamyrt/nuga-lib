package bit_test

import (
	"github.com/mishamyrt/nuga-lib/internal/bit"
	"testing"
)

const (
	First  bit.Mask = 1 << iota
	Second bit.Mask = 2 << iota
	Third  bit.Mask = 4 << iota
	Fourth bit.Mask = 8 << iota
)

func TestIsSet(t *testing.T) {
	t.Parallel()
	var mask bit.Mask
	if mask.IsSet(First) || mask.IsSet(Second) || mask.IsSet(Third) {
		t.Error("Expected mask to be empty")
	}
	mask = First | Second | Third
	if !mask.IsSet(First) || !mask.IsSet(Second) || !mask.IsSet(Third) {
		t.Error("Expected mask to be filled with First, Second and Third")
	}
	if mask.IsSet(Fourth) {
		t.Error("Expected Fourth not to be set to mask")
	}
}

func TestIsSetAny(t *testing.T) {
	t.Parallel()
	var mask bit.Mask
	if mask.IsAnySet(First, Second, Third) {
		t.Error("Expected mask to be truthy")
	}
	mask = First
	if !mask.IsAnySet(First, Second, Third) {
		t.Error("Expected result to be truthy")
	}
	mask = First | Second
	if !mask.IsAnySet(First, Second, Third) {
		t.Error("Expected result to be truthy")
	}
	mask = Third
	if !mask.IsAnySet(Third) {
		t.Error("Expected result to be falsy")
	}
}
