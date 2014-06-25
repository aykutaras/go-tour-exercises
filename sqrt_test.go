package gotourexercises

import (
	"testing"
	"math"
)

func TestSqrt(t *testing.T) {
	const in = 2
	mathSqrt := math.Sqrt(in)
	if x := Sqrt(in); x != mathSqrt {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, mathSqrt)
	}
}