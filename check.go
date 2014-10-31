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
func checkInc(inc ...int) {
	for _, n := range inc {
		if n < 1 {
			panic(fmt.Sprint("blas: invalid stride: %v", n))
		}
	}
}

//
func checkSize(rows, cols int) {
	//TODO
}

// check sizes for matrix-vector multiplication [Ny] = [rows x cols][Nx]
func checkMV(trans Transpose, rows, cols, Nx, Ny int){
if trans == Trans {rows, cols = cols, rows}
	if Nx != cols || Ny != rows{
			panic(fmt.Sprintf("blas: matrix-vector size mismatch for Y=A*X: size(A)=%vx%v, len(X)=%v, len(Y)=%v", rows, cols, Nx, Ny))
	}
}

// return and check sizes for matrix-vector multiplication
func checkSMV(trans Transpose, A[][]float32, X []float32, incX int, Y []float32, incY int) (rows, cols, lda int) {
	rows, cols, lda = SSize(A)
	checkMV(trans, rows, cols, len(X)/incX, len(Y)/incY)
	return
}

// return and check sizes for matrix-vector multiplication
func checkDMV(trans Transpose, A[][]float64, X []float64, incX int, Y []float64, incY int) (rows, cols, lda int) {
	rows, cols, lda = DSize(A)
	checkMV(trans, rows, cols, len(X)/incX, len(Y)/incY)
	return
}

// return and check sizes for matrix-vector multiplication
func checkCMV(trans Transpose, A[][]complex64, X []complex64, incX int, Y []complex64, incY int) (rows, cols, lda int) {
	rows, cols, lda = CSize(A)
	checkMV(trans, rows, cols, len(X)/incX, len(Y)/incY)
	return
}

// return and check sizes for matrix-vector multiplication
func checkZMV(trans Transpose, A[][]complex128, X []complex128, incX int, Y []complex128, incY int) (rows, cols, lda int) {
	rows, cols, lda = ZSize(A)
	checkMV(trans, rows, cols, len(X)/incX, len(Y)/incY)
	return
}

// check for square matrix
func checkSquare(rows, cols int){
	if rows != cols{
		panic(fmt.Sprintf("blas: need square matrix, have %v x %v", rows, cols))
	}
}
