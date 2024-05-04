package hid_test

import (
	"testing"

	"github.com/mishamyrt/nuga-lib/hid"
)

func TestNewErrCountMismatch(t *testing.T) {
	err := hid.NewErrCountMismatch(1, 2)
	if err == nil {
		t.Errorf("NewErrCountMismatch did not return an error")
	}
}
