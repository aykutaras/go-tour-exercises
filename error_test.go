package gotourexercises

import "testing"

func TestSqrtError(t *testing.T) {
	_, err := SqrtError(-2)
	
	if err == nil {
		t.Errorf("Error should be raised")
	}
}