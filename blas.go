package blas

import (
	"unsafe"

	"github.com/barnex/blas/cblas"
)

// Computes the dot product of vectors X and Y plus an initial value alpha.
// Every incX'th and incY'th element is used.
func SDSDOT(alpha float32, X []float32, incX int, Y []float32, incY int) float32 {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	return cblas.CBLAS_SDSDOT(N, alpha, X_, incX, Y_, incY)
}

// Computes the dot product of vectors X and Y.
// Every incX'th and incY'th element is used.
func DSDOT(X []float32, incX int, Y []float32, incY int) float64 {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	return cblas.CBLAS_DSDOT(N, X_, incX, Y_, incY)
}

// Computes the dot product of vectors X and Y.
// Every incX'th and incY'th element is used.
func SDOT(X []float32, incX int, Y []float32, incY int) float32 {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	return cblas.CBLAS_SDOT(N, X_, incX, Y_, incY)
}

// Computes the dot product of vectors X and Y.
// Every incX'th and incY'th element is used.
func DDOT(X []float64, incX int, Y []float64, incY int) float64 {
	var N int = checkIncD(X, incX, Y, incY)
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	return cblas.CBLAS_DDOT(N, X_, incX, Y_, incY)
}

// Computes the dot product of vectors X and Y.
// Every incX'th and incY'th element is used.
func CDOTU(X []complex64, incX int, Y []complex64, incY int) complex64 {
	var result complex64
	var N int = checkIncC(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CDOTU_SUB(N, X_, incX, Y_, incY, unsafe.Pointer(&result))
	return result
}

// Calculates the dot product of the complex conjugate of X with Y.
// Every incX'th and incY'th element is used.
func CDOTC(X []complex64, incX int, Y []complex64, incY int) complex64 {
	var result complex64
	var N int = checkIncC(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CDOTC_SUB(N, X_, incX, Y_, incY, unsafe.Pointer(&result))
	return result
}

// Computes the dot product of vectors X and Y.
// Every incX'th and incY'th element is used.
func ZDOTU(X []complex128, incX int, Y []complex128, incY int) complex128 {
	var result complex128
	var N int = checkIncZ(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.ZDOTU_SUB(N, X_, incX, Y_, incY, unsafe.Pointer(&result))
	return result
}

// Calculates the dot product of the complex conjugate of X with Y.
// Every incX'th and incY'th element is used.
func ZDOTC(X []complex128, incX int, Y []complex128, incY int) complex128 {
	var result complex128
	var N int = checkIncZ(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.ZDOTC_SUB(N, X_, incX, Y_, incY, unsafe.Pointer(&result))
	return result
}

// Computes the L2 norm (Euclidian length) of vector X.
// Every incX'th element is used.
func SNRM2(X []float32, incX int) float32 {
	var N int = len(X) / incX
	var X_ *float32 = &X[0]
	return cblas.CBLAS_SNRM2(N, X_, incX)
}

// Computes the sum of the absolute values of the elements in vector X.
// Every incX'th element is used.
func SASUM(X []float32, incX int) float32 {
	var N int = len(X) / incX
	var X_ *float32 = &X[0]
	return cblas.CBLAS_SASUM(N, X_, incX)
}

// Computes the L2 norm (Euclidian length) of vector X.
// Every incX'th element is used.
func DNRM2(X []float64, incX int) float64 {
	var N int = len(X) / incX
	var X_ *float64 = &X[0]
	return cblas.CBLAS_DNRM2(N, X_, incX)
}

// Computes the sum of the absolute values of the elements in vector X.
// Every incX'th element is used.
func DASUM(X []float64, incX int) float64 {
	var N int = len(X) / incX
	var X_ *float64 = &X[0]
	return cblas.CBLAS_DASUM(N, X_, incX)
}

// Computes the unitary norm of vector X.
// Every incX'th element is used.
func SCNRM2(X []complex64, incX int) float32 {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	return cblas.CBLAS_SCNRM2(N, X_, incX)
}

// Computes the sum of the absolute values of real and imaginary parts of elements in vector X.
// Every incX'th element is used.
func SCASUM(X []complex64, incX int) float32 {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	return cblas.CBLAS_SCASUM(N, X_, incX)
}

// Computes the unitary norm of vector X.
// Every incX'th element is used.
func DZNRM2(X []complex128, incX int) float64 {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	return cblas.CBLAS_DZNRM2(N, X_, incX)
}

// Computes the sum of the absolute values of real and imaginary parts of elements in vector X.
// Every incX'th element is used.
func DZASUM(X []complex128, incX int) float64 {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	return cblas.CBLAS_DZASUM(N, X_, incX)
}

// Returns the index of the element with the largest absolute value in vector X.
// Every incX'th element is used.
func ISAMAX(X []float32, incX int) int {
	var N int = len(X) / incX
	var X_ *float32 = &X[0]
	return cblas.CBLAS_ISAMAX(N, X_, incX)
}

// Returns the index of the element with the largest absolute value in vector X.
// Every incX'th element is used.
func IDAMAX(X []float64, incX int) int {
	var N int = len(X) / incX
	var X_ *float64 = &X[0]
	return cblas.CBLAS_IDAMAX(N, X_, incX)
}

// Returns the index of the element with the largest absolute value in vector X.
// Every incX'th element is used.
func ICAMAX(X []complex64, incX int) int {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	return cblas.CBLAS_ICAMAX(N, X_, incX)
}

// Returns the index of the element with the largest absolute value in vector X.
// Every incX'th element is used.
func IZAMAX(X []complex128, incX int) int {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	return cblas.CBLAS_IZAMAX(N, X_, incX)
}

// Exchanges the elements of vectors X and Y.
// Every incX'th and incY'th element is used.
func SSWAP(X []float32, incX int, Y []float32, incY int) {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SSWAP(N, X_, incX, Y_, incY)
}

// Copies X to Y.
// Every incX'th and incY'th element is used.
func SCOPY(X []float32, incX int, Y []float32, incY int) {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SCOPY(N, X_, incX, Y_, incY)
}

// Replaces Y by (alpha*X) + Y.
// Every incX'th and incY'th element is used.
func SAXPY(alpha float32, X []float32, incX int, Y []float32, incY int) {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SAXPY(N, alpha, X_, incX, Y_, incY)
}

// Exchanges the elements of vectors X and Y.
// Every incX'th and incY'th element is used.
func DSWAP(X []float64, incX int, Y []float64, incY int) {
	var N int = checkIncD(X, incX, Y, incY)
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	cblas.CBLAS_DSWAP(N, X_, incX, Y_, incY)
}

// Copies X to Y.
// Every incX'th and incY'th element is used.
func DCOPY(X []float64, incX int, Y []float64, incY int) {
	var N int = checkIncD(X, incX, Y, incY)
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	cblas.CBLAS_DCOPY(N, X_, incX, Y_, incY)
}

// Replaces Y by (alpha*X) + Y.
// Every incX'th and incY'th element is used.
func DAXPY(alpha float64, X []float64, incX int, Y []float64, incY int) {
	var N int = checkIncD(X, incX, Y, incY)
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	cblas.CBLAS_DAXPY(N, alpha, X_, incX, Y_, incY)
}

// Exchanges the elements of vectors X and Y.
// Every incX'th and incY'th element is used.
func CSWAP(X []complex64, incX int, Y []complex64, incY int) {
	var N int = checkIncC(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_CSWAP(N, X_, incX, Y_, incY)
}

// Copies X to Y.
// Every incX'th and incY'th element is used.
func CCOPY(X []complex64, incX int, Y []complex64, incY int) {
	var N int = checkIncC(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_CCOPY(N, X_, incX, Y_, incY)
}

// Replaces Y by (alpha*X) + Y.
// Every incX'th and incY'th element is used.
func CAXPY(alpha complex64, X []complex64, incX int, Y []complex64, incY int) {
	var N int = checkIncC(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_CAXPY(N, unsafe.Pointer(&alpha), X_, incX, Y_, incY)
}

// Exchanges the elements of vectors X and Y.
// Every incX'th and incY'th element is used.
func ZSWAP(X []complex128, incX int, Y []complex128, incY int) {
	var N int = checkIncZ(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_ZSWAP(N, X_, incX, Y_, incY)
}

// Copies X to Y.
// Every incX'th and incY'th element is used.
func ZCOPY(X []complex128, incX int, Y []complex128, incY int) {
	var N int = checkIncZ(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_ZCOPY(N, X_, incX, Y_, incY)
}

// Replaces Y by (alpha*X) + Y.
// Every incX'th and incY'th element is used.
func ZAXPY(alpha complex128, X []complex128, incX int, Y []complex128, incY int) {
	var N int = checkIncZ(X, incX, Y, incY)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_ZAXPY(N, unsafe.Pointer(&alpha), X_, incX, Y_, incY)
}

// Constructs a Givens rotation matrix.
// Computes the values of c and s so that:
//
//    [ c  s ] [ a ]    [ r ]
//    [ -s c ] [ b ] =  [ 0 ].
//
// Returns r, c and s.
func SROTG(a float32, b float32) (r, c, s float32) {
	var a_ float32 = a
	var b_ float32 = b
	var c_ float32 = 0
	var s_ float32 = 0
	cblas.CBLAS_SROTG(&a_, &b_, &c_, &s_)
	return a_, c_, s_
}

// Constructs a modified Givens rotation matrix.
// The input scalars d1, d2, x1 and y1 define a 2-vector [a1 a2]' such that
//
//    [ b1 ]   [ d1^{1/2}  0      ] [ x1 ]
//    [ b2 ] = [   0     d2^{1/2} ] [ y1 ].
//
// This subroutine determines the modified Givens rotation matrix H that
// transforms y1 and thus a2 to zero.
// A representation of this matrix is returned:
//
// P[0]: defines the form of matrix H:
// 	-2.0: matrix H contains the identity matrix.
// 	-1.0: matrix H is identical to matrix SH (defined by the remaining values in the vector).
// 	 0.0: H[1,2] and H[2,1] are obtained from matrix SH. The remaining values are both 1.0.
// 	 1.0: H[1,1] and H[2,2] are obtained from matrix SH. H[1,2] is 1.0. H[2,1] is -1.0.
//
// The other elements contain SH:
//
// 	P[1]	contains SH[1,1].
// 	P[2]	contains SH[2,1].
// 	P[3]	contains SH[1,2].
// 	P[4]	contains SH[2,2].
//
func SROTMG(d1 float32, d2 float32, b1 float32, b2 float32) [5]float32 {
	var d1_ float32 = d1
	var d2_ float32 = d2
	var b1_ float32 = b1
	var P [5]float32
	cblas.CBLAS_SROTMG(&d1_, &d2_, &b1_, b2, &P[0])
	return P
}

// Applies a plane rotation to the two vectors X and Y,
// This routine computes:
//
//    [ x_i ]   [ c s ] [ x_i ]
//    [ y_i ] = [-s c ] [ y_i ]
//
// for all i with strides incX and incY.
func SROT(X []float32, incX int, Y []float32, incY int, c float32, s float32) {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SROT(N, X_, incX, Y_, incY, c, s)
}

// Applies the modified-Givens rotation of the vectors X and Y:
//
//      [ x_i ]   [ h_11 h_12 ] [ x_i ]
//      [ y_i ] = [ h_21 h_22 ] [ y_i ]
// for all i.
//
// H  stored in P according to the encoding used by SROTMG.
func SROTM(X []float32, incX int, Y []float32, incY int, P [5]float32) {
	var N int = checkIncS(X, incX, Y, incY)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SROTM(N, X_, incX, Y_, incY, &P[0])
}

// Constructs a Givens rotation matrix.
// Computes the values of c and s so that:
//
//    [ c  s ] [ a ]    [ r ]
//    [ -s c ] [ b ] =  [ 0 ].
//
// Returns r, c and s.
func DROTG(a float64, b float64) (r, c, s float64) {
	var a_ float64 = a
	var b_ float64 = b
	var c_ float64 = 0
	var s_ float64 = 0
	cblas.CBLAS_DROTG(&a_, &b_, &c_, &s_)
	return a_, c_, s_
}

// Constructs a modified Givens rotation matrix.
// The input scalars d1, d2, x1 and y1 define a 2-vector [a1 a2]' such that
//
//    [ b1 ]   [ d1^{1/2}  0      ] [ x1 ]
//    [ b2 ] = [   0     d2^{1/2} ] [ y1 ].
//
// This subroutine determines the modified Givens rotation matrix H that
// transforms y1 and thus a2 to zero.
// A representation of this matrix is returned:
//
// P[0]: defines the form of matrix H:
// 	-2.0: matrix H contains the identity matrix.
// 	-1.0: matrix H is identical to matrix SH (defined by the remaining values in the vector).
// 	 0.0: H[1,2] and H[2,1] are obtained from matrix SH. The remaining values are both 1.0.
// 	 1.0: H[1,1] and H[2,2] are obtained from matrix SH. H[1,2] is 1.0. H[2,1] is -1.0.
//
// The other elements contain SH:
//
// 	P[1]	contains SH[1,1].
// 	P[2]	contains SH[2,1].
// 	P[3]	contains SH[1,2].
// 	P[4]	contains SH[2,2].
//
func DROTMG(d1 float64, d2 float64, b1 float64, b2 float64) [5]float64 {
	var d1_ float64 = d1
	var d2_ float64 = d2
	var b1_ float64 = b1
	var P [5]float64
	cblas.CBLAS_DROTMG(&d1_, &d2_, &b1_, b2, &P[0])
	return P
}

// Applies a plane rotation to the two vectors X and Y.
// This routine computes:
//
//    [ x_i ]   [ c s ] [ x_i ]
//    [ y_i ] = [-s c ] [ y_i ]
// for all i with strides incX, incY.
func DROT(X []float64, incX int, Y []float64, incY int, c float64, s float64) {
	var N int = checkIncD(X, incX, Y, incY)
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	cblas.CBLAS_DROT(N, X_, incX, Y_, incY, c, s)
}

// Applies the modified-Givens rotation of the vectors X and Y:
//
//      [ x_i ]   [ h_11 h_12 ] [ x_i ]
//      [ y_i ] = [ h_21 h_22 ] [ y_i ]
//
// for all i with strides incX, incY
//
// H  stored in P according to the encoding used by DROTMG.
func DROTM(X []float64, incX int, Y []float64, incY int, P [5]float64) {
	var N int = checkIncD(X, incX, Y, incY)
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	cblas.CBLAS_DROTM(N, X_, incX, Y_, incY, &P[0])
}

// Multiply X by alpha.
// Every incX'th element is used.
func SSCAL(alpha float32, X []float32, incX int) {
	var N int = len(X) / incX
	var X_ *float32 = &X[0]
	cblas.CBLAS_SSCAL(N, alpha, X_, incX)
}

// Multiply X by alpha.
// Every incX'th element is used.
func DSCAL(alpha float64, X []float64, incX int) {
	var N int = len(X) / incX
	var X_ *float64 = &X[0]
	cblas.CBLAS_DSCAL(N, alpha, X_, incX)
}

// Multiply X by alpha.
// Every incX'th element is used.
func CSCAL(alpha complex64, X []complex64, incX int) {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	cblas.CBLAS_CSCAL(N, unsafe.Pointer(&alpha), X_, incX)
}

// Multiply X by alpha.
// Every incX'th element is used.
func ZSCAL(alpha complex128, X []complex128, incX int) {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	cblas.CBLAS_ZSCAL(N, unsafe.Pointer(&alpha), X_, incX)
}

// Multiply X by alpha.
// Every incX'th element is used.
func CSSCAL(alpha float32, X []complex64, incX int) {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	cblas.CBLAS_CSSCAL(N, alpha, X_, incX)
}

// Multiply X by alpha.
// Every incX'th element is used.
func ZDSCAL(alpha float64, X []complex128, incX int) {
	var N int = len(X) / incX
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	cblas.CBLAS_ZDSCAL(N, alpha, X_, incX)
}

// Matrix-vector multiplication
// 	Y = alpha*AX + beta*Y
func SGEMV(transA Transpose, alpha float32, A [][]float32, X []float32, incX int, beta float32, Y []float32, incY int) {
	rows, cols, lda := SSize(A)
	checkSMV(transA, A, X, incX, Y, incY)
	var A_ *float32 = &A[0][0]
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SGEMV(uint32(RowMajor), uint32(transA), rows, cols, alpha, A_, lda, X_, incX, beta, Y_, incY)
}

/*
func STRMV(Uplo Uplo, transA Transpose, Diag Diag, A [][]float32, X []float32) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_STRMV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func STRSV(Uplo Uplo, transA Transpose, Diag Diag, A [][]float32, X []float32) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_STRSV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func DGEMV(transA Transpose, alpha float64, A [][]float64, X []float64, beta float64, Y []float64) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var beta_ float64 = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DGEMV(RowMajor, uint32(transA), M_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func DTRMV(Uplo Uplo, transA Transpose, Diag Diag, A [][]float64, X []float64) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_DTRMV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func DTRSV(Uplo Uplo, transA Transpose, Diag Diag, A [][]float64, X []float64) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_DTRSV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func CGEMV(transA Transpose, alpha complex64, A [][]complex64, X []complex64, beta complex64, Y []complex64) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var beta_ unsafe.Pointer = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_CGEMV(RowMajor, uint32(transA), M_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func CTRMV(Uplo Uplo, transA Transpose, Diag Diag, A [][]complex64, X []complex64) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_CTRMV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func CTRSV(Uplo Uplo, transA Transpose, Diag Diag, A [][]complex64, X []complex64) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_CTRSV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func ZGEMV(transA Transpose, alpha complex128, A [][]complex128, X []complex128, beta complex128, Y []complex128) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var beta_ unsafe.Pointer = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_ZGEMV(RowMajor, uint32(transA), M_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func ZTRMV(Uplo Uplo, transA Transpose, Diag Diag, A [][]complex128, X []complex128) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_ZTRMV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func ZTRSV(Uplo Uplo, transA Transpose, Diag Diag, A [][]complex128, X []complex128) {
	var Uplo_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_ZTRSV(RowMajor, Uplo_, uint32(transA), Diag_, N_, A_, lda_, X_, incX_)
}

func SSYMV(Uplo Uplo, alpha float32, A [][]float32, X []float32, beta float32, Y []float32) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var beta_ float32 = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_SSYMV(RowMajor, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func SGER(alpha float32, X []float32, Y []float32, A [][]float32) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	cblas.CBLAS_SGER(RowMajor, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func SSYR(Uplo Uplo, alpha float32, X []float32, A [][]float32) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	cblas.CBLAS_SSYR(RowMajor, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func SSYR2(Uplo Uplo, alpha float32, X []float32, Y []float32, A [][]float32) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	cblas.CBLAS_SSYR2(RowMajor, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func DSYMV(Uplo Uplo, alpha float64, A [][]float64, X []float64, beta float64, Y []float64) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var beta_ float64 = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DSYMV(RowMajor, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func DGER(alpha float64, X []float64, Y []float64, A [][]float64) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	cblas.CBLAS_DGER(RowMajor, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func DSYR(Uplo Uplo, alpha float64, X []float64, A [][]float64) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	cblas.CBLAS_DSYR(RowMajor, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func DSYR2(Uplo Uplo, alpha float64, X []float64, Y []float64, A [][]float64) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	cblas.CBLAS_DSYR2(RowMajor, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func CHEMV(Uplo Uplo, alpha complex64, A [][]complex64, X []complex64, beta complex64, Y []complex64) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var beta_ unsafe.Pointer = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_CHEMV(RowMajor, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func CGERU(alpha complex64, X []complex64, Y []complex64, A [][]complex64) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CGERU(RowMajor, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func CGERC(alpha complex64, X []complex64, Y []complex64, A [][]complex64) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CGERC(RowMajor, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func CHER(Uplo Uplo, alpha float32, X []complex64, A [][]complex64) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CHER(RowMajor, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func CHER2(Uplo Uplo, alpha complex64, X []complex64, Y []complex64, A [][]complex64) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CHER2(RowMajor, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func ZHEMV(Uplo Uplo, alpha complex128, A [][]complex128, X []complex128, beta complex128, Y []complex128) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var beta_ unsafe.Pointer = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_ZHEMV(RowMajor, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func ZGERU(alpha complex128, X []complex128, Y []complex128, A [][]complex128) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZGERU(RowMajor, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func ZGERC(alpha complex128, X []complex128, Y []complex128, A [][]complex128) {
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZGERC(RowMajor, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func ZHER(Uplo Uplo, alpha float64, X []complex128, A [][]complex128) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZHER(RowMajor, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func ZHER2(Uplo Uplo, alpha complex128, X []complex128, Y []complex128, A [][]complex128) {
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZHER2(RowMajor, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func SGEMM(transA Transpose, transB Transpose, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {

	var M_ int = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SGEMM(RowMajor, uint32(transA), uint32(transB), M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func SSYMM(Side Side, Uplo Uplo, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SSYMM(RowMajor, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func SSYRK(Uplo Uplo, trans Transpose, alpha float32, A [][]float32, beta float32, C [][]float32) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SSYRK(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func SSYR2K(Uplo Uplo, trans Transpose, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SSYR2K(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func STRMM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha float32, A [][]float32, B [][]float32) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	cblas.CBLAS_STRMM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func STRSM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha float32, A [][]float32, B [][]float32) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = &A[0][0]
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	cblas.CBLAS_STRSM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func DGEMM(transA Transpose, transB Transpose, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {


	var M_ int = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DGEMM(RowMajor, uint32(transA), uint32(transB), M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func DSYMM(Side Side, Uplo Uplo, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DSYMM(RowMajor, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func DSYRK(Uplo Uplo, trans Transpose, alpha float64, A [][]float64, beta float64, C [][]float64) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DSYRK(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func DSYR2K(Uplo Uplo, trans Transpose, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DSYR2K(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func DTRMM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha float64, A [][]float64, B [][]float64) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	cblas.CBLAS_DTRMM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func DTRSM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha float64, A [][]float64, B [][]float64) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = &A[0][0]
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	cblas.CBLAS_DTRSM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func CGEMM(transA Transpose, transB Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {


	var M_ int = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CGEMM(RowMajor, uint32(transA), uint32(transB), M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CSYMM(Side Side, Uplo Uplo, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CSYMM(RowMajor, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CSYRK(Uplo Uplo, trans Transpose, alpha complex64, A [][]complex64, beta complex64, C [][]complex64) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CSYRK(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func CSYR2K(Uplo Uplo, trans Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CSYR2K(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CTRMM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha complex64, A [][]complex64, B [][]complex64) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_CTRMM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func CTRSM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha complex64, A [][]complex64, B [][]complex64) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_CTRSM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func ZGEMM(transA Transpose, transB Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {


	var M_ int = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZGEMM(RowMajor, uint32(transA), uint32(transB), M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZSYMM(Side Side, Uplo Uplo, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZSYMM(RowMajor, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZSYRK(Uplo Uplo, trans Transpose, alpha complex128, A [][]complex128, beta complex128, C [][]complex128) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZSYRK(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func ZSYR2K(Uplo Uplo, trans Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZSYR2K(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZTRMM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha complex128, A [][]complex128, B [][]complex128) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_ZTRMM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func ZTRSM(Side Side, Uplo Uplo, transA Transpose, Diag Diag, alpha complex128, A [][]complex128, B [][]complex128) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0

	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_ZTRSM(RowMajor, Side_, Uplo_, uint32(transA), Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func CHEMM(Side Side, Uplo Uplo, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CHEMM(RowMajor, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CHERK(Uplo Uplo, trans Transpose, alpha float32, A [][]complex64, beta float32, C [][]complex64) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ float32 = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CHERK(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func CHER2K(Uplo Uplo, trans Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta float32, C [][]complex64) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ float32 = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CHER2K(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZHEMM(Side Side, Uplo Uplo, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZHEMM(RowMajor, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZHERK(Uplo Uplo, trans Transpose, alpha float64, A [][]complex128, beta float64, C [][]complex128) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ float64 = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZHERK(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func ZHER2K(Uplo Uplo, trans Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta float64, C [][]complex128) {
	var Uplo_ uint32 = 0
	var uint32(trans) uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	var beta_ float64 = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZHER2K(RowMajor, Uplo_, uint32(trans), N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}
*/
