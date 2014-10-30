package blas

import (
	"fmt"
)

func  checkIncS(X []float32, incX int, Y []float32, incY int) int{
	N := len(X)/incX
	if len(Y)/incY != N{
			panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}


func  checkIncD(X []float64, incX int, Y []float64, incY int) int{
	N := len(X)/incX
	if len(Y)/incY != N{
			panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}

func  checkIncC(X []complex64, incX int, Y []complex64, incY int) int{
	N := len(X)/incX
	if len(Y)/incY != N{
			panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}


func  checkIncZ(X []complex128, incX int, Y []complex128, incY int) int{
	N := len(X)/incX
	if len(Y)/incY != N{
			panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}
