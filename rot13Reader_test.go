package gotourexercises

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRot13Reader(t *testing.T) {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	
	written, _ := io.Copy(os.Stdout, &r)
	
	if written != 21 {
		t.Errorf("Expected %v, found %v", 21, written)
	}
}