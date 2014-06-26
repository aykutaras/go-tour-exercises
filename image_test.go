package gotourexercises

import (
	"testing"
	"code.google.com/p/go-tour/pic"	
)

func TestImage(t *testing.T) {
    m := Image{255,255}
    pic.ShowImage(m)

	if  false {
		t.Errorf("Something went wrong")
	}
}