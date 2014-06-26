package gotourexercises

import (
	"testing"
)

func TestCbrt(t *testing.T) {
	expected := 2
	if found := Cbrt(8); found != 2 {
		t.Errorf("Expected %v, found %v", expected, found)
	}
}