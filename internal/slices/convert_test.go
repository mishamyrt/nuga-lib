package slices_test

import (
	"github.com/mishamyrt/nuga-lib/internal/slices"
	"reflect"
	"testing"
)

// TestCastIntToFloat32 tests the Cast function with int to float32 conversion
func TestCastIntToFloat32(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 3}
	expected := []float32{1.0, 2.0, 3.0}

	result := slices.Cast[int, float32](input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

// TestCastFloat64ToInt tests the Cast function with float64 to int conversion
func TestCastFloat64ToInt(t *testing.T) {
	t.Parallel()
	input := []float64{1.1, 2.2, 3.3}
	expected := []int{1, 2, 3}

	result := slices.Cast[float64, int](input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

// TestCastBytesToInt tests the Cast function with float64 to int conversion
func TestCastBytesToInt(t *testing.T) {
	t.Parallel()
	input := []byte{1, 2, 3}
	expected := []int{1, 2, 3}

	result := slices.Cast[byte, int](input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
