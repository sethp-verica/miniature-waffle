//go:build moretests

package math

import "testing"

func TestSubtract(t *testing.T) {
	if got, want := Subtract(3, 2), 1; got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}
}
