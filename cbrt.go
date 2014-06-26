package gotourexercises

import "math/cmplx"

func Cbrt(x complex128) complex128 {
    z := x
    
    for i:=0; i<20; i++ {
        zCube := cmplx.Pow(z, 3)
    	zSquare := cmplx.Pow(z, 2)
    	z = z - (zCube - x)/(3*zSquare)
    }
    
    return z
}