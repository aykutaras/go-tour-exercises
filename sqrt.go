package gotourexercises

import "math"

func Sqrt(x float64) float64 {
    const delta = 1e-10
    z, zOld := x, x
    
    for {
        zOld = z
        z = z - (z*z - x)/(2*z)
        
        if math.Abs(zOld - z) < delta {
            break
        }
    }
    
    return z
}