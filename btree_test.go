package gotourexercises

import (
	"testing"
	"code.google.com/p/go-tour/tree"
	"fmt"
)

func TestSame(t *testing.T) {
    shouldBeTrue := Same(tree.New(1), tree.New(1))
    shouldBeFalse := Same(tree.New(1), tree.New(2))
	
	if !shouldBeTrue {
		t.Errorf("Two binary trees should be same")
	}
	
	if shouldBeFalse {
		t.Errorf("Two binary trees should not be same")
	}
}

func TestWalk(t *testing.T) {
    ch := make(chan int)
    go Walk(tree.New(3), ch)
    
	count := 0
	
    for leaf := range ch {
        fmt.Println(leaf)
		count++
    }
	
	if count != 10 {
		t.Errorf("Tree element count should be 10")
	}
}