package gotourexercises

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := fibonacci()
	tenthFibonacci := 0
    for i := 0; i < 10; i++ {
		tenthFibonacci = f()
        fmt.Println(tenthFibonacci)
    }
	
	if tenthFibonacci != 34 {
		t.Errorf("Expected number %v, found %v", 34, tenthFibonacci)
	}
}