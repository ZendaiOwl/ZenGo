package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	var count int
	var character string
	var expected string
	repeated := Repeat(character, count)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
