package bit_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/internal/bit"
)

func TestCombine16(t *testing.T) {
	t.Parallel()
	result := bit.Combine16(0x12, 0x34)
	if result != 0x1234 {
		t.Errorf("Expected 0x1234, got 0x%x", result)
	}
}

func TestSplit16(t *testing.T) {
	t.Parallel()
	high, low := bit.Split16(0x1234)
	if high != 0x12 || low != 0x34 {
		t.Errorf("Expected 0x12, 0x34, got 0x%x, 0x%x", high, low)
	}
}
