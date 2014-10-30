package blas

import (
	"fmt"
)

// check if X and Y have the some number of elements (considering their stride),
// and return that number.
func checkIncS(X []float32, incX int, Y []float32, incY int) int {
	checkInc(incX, incY)
	N := len(X) / incX
	if len(Y)/incY != N {
		panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}

// check if X and Y have the some number of elements (considering their stride),
// and return that number.
func checkIncD(X []float64, incX int, Y []float64, incY int) int {
	checkInc(incX, incY)
	N := len(X) / incX
	if len(Y)/incY != N {
		panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}

// check if X and Y have the some number of elements (considering their stride),
// and return that number.
func checkIncC(X []complex64, incX int, Y []complex64, incY int) int {
	checkInc(incX, incY)
	N := len(X) / incX
	if len(Y)/incY != N {
		panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}

// check if X and Y have the some number of elements (considering their stride),
// and return that number.
func checkIncZ(X []complex128, incX int, Y []complex128, incY int) int {
	checkInc(incX, incY)
	N := len(X) / incX
	if len(Y)/incY != N {
		panic(fmt.Sprintf("blas: unmatched vector sizes: len(X)=%v, incX=%v, len(Y)=%v, incY=%v", len(X), incX, len(Y), incY))
	}
	return N
}

// check for positive strides.
func checkInc(inc ...int){
	for _,n:=range inc{
		if n<1{panic(fmt.Sprint("blas: invalid stride: %v", n))}
	}
}

func checkSMV(A [][]float32,X[]float32,incX int,Y[]float32,incY int) (M, N, lda int){
	N = checkIncS(X,incX,Y,incY)
	m,n := len(A),len(A[0])
	if n!=N{
		// TODO: Transpose
		panic(fmt.Sprint("blas: unmatched matrix*vector sizes: [%v x %v]*[%v]", m, n, N))
	}
	return m, N, m
}


