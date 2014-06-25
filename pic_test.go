package gotourexercises

import (
	"testing"
	"code.google.com/p/go-tour/pic"	
)

func TestPic(t *testing.T) {
	pic.Show(Pic)

	if  Pic == nil {
		t.Errorf("Something went wrong")
	}
}