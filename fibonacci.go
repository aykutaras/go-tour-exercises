package gotourexercises

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    n, fib0, fib1 := 0, 0, 1
    return func() int {
        if n == 0 {
            n++
            return 0
        }
        
        if n == 1 {
            n++
            return 1
        } else {
        	fib := fib0 + fib1
            fib0 = fib1
            fib1 = fib
            return fib
        }
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}