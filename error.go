package gotourexercises

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func SqrtError(f float64) (float64, error) {
    if f < 0 {
        return f, ErrNegativeSqrt(f)
    }
    
    const delta = 1e-10
    z, zOld := f, f
    
    for {
        zOld = z
        z = z - (z*z - f)/(2*z)
        
        if math.Abs(zOld - z) < delta {
            break
        }
    }
    
    return z, nil
}