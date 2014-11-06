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

// Matrix-vector multiplication plus vector with optional matrix transpose.
// When transA == NoTrans this computes:
// 	Y = alpha*A*X + beta*Y
// When transA == Trans this computes:
// 	Y = alpha*(A^T)*X + beta*Y
// Every incX'th element of X and incY'th element of Y is used.
// Matrices must be allocated with MakeFloat32Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func SGEMV(transA Transpose, alpha float32, A [][]float32, X []float32, incX int, beta float32, Y []float32, incY int) {
	rows, cols, lda := checkSMV(transA, A, X, incX, Y, incY)
	var A_ *float32 = &A[0][0]
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SGEMV(uint32(RowMajor), uint32(transA), rows, cols, alpha, A_, lda, X_, incX, beta, Y_, incY)
}

// Triangular matrix-vector multiplication plus vector with optional matrix transpose.
// With uplo == Upper or Lower, the upper or lower triangular part of A is used.
// diag specifies whether the matrix is unit triangular ().
// When transA == NoTrans this computes:
// 	X = A*X
// When transA == Trans this computes:
// 	X = (A^T)*X
// Every incX'th element of X is used.
// Matrices must be allocated with MakeFloat32Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func STRMV(uplo Uplo, transA Transpose, diag Diag, A [][]float32, X []float32, incX int) {
	rows, cols, lda := checkSMV(transA, A, X, incX, X, incX) // TODO: separate func without 2nd X, incX
	checkSquare(rows, cols)
	var A_ *float32 = &A[0][0]
	var X_ *float32 = &X[0]
	cblas.CBLAS_STRMV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), rows, A_, lda, X_, incX)
}

/*
func STRSV(uplo Uplo, transA Transpose, diag Diag, A [][]float32, X []float32, incX) {
	rows, cols, lda := checkSMV(transA, A, X, incX, Y, incY)
	var A_ *float32 = &A[0][0]
	var X_ *float32 = 0
	cblas.CBLAS_STRSV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), N_, A_, lda, X_, incX)
}
*/

// Matrix-vector multiplication plus vector with optional matrix transpose.
// When transA == NoTrans this computes:
// 	Y = alpha*A*X + beta*Y
// When transA == Trans this computes:
// 	Y = alpha*(A^T)*X + beta*Y
// Every incX'th element of X and incY'th element of Y is used.
// Matrices must be allocated with MakeFloat64Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func DGEMV(transA Transpose, alpha float64, A [][]float64, X []float64, incX int, beta float64, Y []float64, incY int) {
	rows, cols, lda := checkDMV(transA, A, X, incX, Y, incY)
	var A_ *float64 = &A[0][0]
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	cblas.CBLAS_DGEMV(uint32(RowMajor), uint32(transA), rows, cols, alpha, A_, lda, X_, incX, beta, Y_, incY)
}

// Triangular matrix-vector multiplication plus vector with optional matrix transpose.
// With uplo == Upper or Lower, the upper or lower triangular part of A is used.
// diag specifies whether the matrix is unit triangular ().
// When transA == NoTrans this computes:
// 	X = A*X
// When transA == Trans this computes:
// 	X = (A^T)*X
// Every incX'th element of X is used.
// Matrices must be allocated with MakeFloat64Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func DTRMV(uplo Uplo, transA Transpose, diag Diag, A [][]float64, X []float64, incX int) {
	rows, cols, lda := checkDMV(transA, A, X, incX, X, incX) // TODO: separate func without 2nd X, incX
	checkSquare(rows, cols)
	var A_ *float64 = &A[0][0]
	var X_ *float64 = &X[0]
	cblas.CBLAS_DTRMV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), rows, A_, lda, X_, incX)
}

/*
func DTRSV(uplo Uplo, transA Transpose, diag Diag, A [][]float64, X []float64) {



	var A_ *float64 = &A[0][0]

	var X_ *float64 = 0

	cblas.CBLAS_DTRSV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), N_, A_, lda, X_, incX)
}
*/

// Matrix-vector multiplication plus vector with optional matrix transpose.
// When transA == NoTrans this computes:
// 	Y = alpha*A*X + beta*Y
// When transA == Trans this computes:
// 	Y = alpha*(A^T)*X + beta*Y
// Every incX'th element of X and incY'th element of Y is used.
// Matrices must be allocated with MakeComplex64Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func CGEMV(transA Transpose, alpha complex64, A [][]complex64, X []complex64, incX int, beta complex64, Y []complex64, incY int) {
	rows, cols, lda := checkCMV(transA, A, X, incX, Y, incY)
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_CGEMV(uint32(RowMajor), uint32(transA), rows, cols, unsafe.Pointer(&alpha), A_, lda, X_, incX, unsafe.Pointer(&beta), Y_, incY)
}

// Triangular matrix-vector multiplication plus vector with optional matrix transpose.
// With uplo == Upper or Lower, the upper or lower triangular part of A is used.
// diag specifies whether the matrix is unit triangular ().
// When transA == NoTrans this computes:
// 	X = A*X
// When transA == Trans this computes:
// 	X = (A^T)*X
// Every incX'th element of X is used.
// Matrices must be allocated with MakeComplex64Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func CTRMV(uplo Uplo, transA Transpose, diag Diag, A [][]complex64, X []complex64, incX int) {
	rows, cols, lda := checkCMV(transA, A, X, incX, X, incX)
	checkSquare(rows, cols)
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	cblas.CBLAS_CTRMV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), rows, A_, lda, X_, incX)
}

/*

func CTRSV(uplo Uplo, transA Transpose, diag Diag, A [][]complex64, X []complex64) {



	var A_ unsafe.Pointer = 0

	var X_ unsafe.Pointer = 0

	cblas.CBLAS_CTRSV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), N_, A_, lda, X_, incX)
}
*/

// Matrix-vector multiplication plus vector with optional matrix transpose.
// When transA == NoTrans this computes:
// 	Y = alpha*A*X + beta*Y
// When transA == Trans this computes:
// 	Y = alpha*(A^T)*X + beta*Y
// Every incX'th element of X and incY'th element of Y is used.
// Matrices must be allocated with MakeComplex128Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func ZGEMV(transA Transpose, alpha complex128, A [][]complex128, X []complex128, incX int, beta complex128, Y []complex128, incY int) {
	rows, cols, lda := checkZMV(transA, A, X, incX, Y, incY)
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_ZGEMV(uint32(RowMajor), uint32(transA), rows, cols, unsafe.Pointer(&alpha), A_, lda, X_, incX, unsafe.Pointer(&beta), Y_, incY)
}

// Triangular matrix-vector multiplication plus vector with optional matrix transpose.
// With uplo == Upper or Lower, the upper or lower triangular part of A is used.
// diag specifies whether the matrix is unit triangular ().
// When transA == NoTrans this computes:
// 	X = A*X
// When transA == Trans this computes:
// 	X = (A^T)*X
// Every incX'th element of X is used.
// Matrices must be allocated with MakeComplex128Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
func ZTRMV(uplo Uplo, transA Transpose, diag Diag, A [][]complex128, X []complex128, incX int) {
	rows, cols, lda := checkZMV(transA, A, X, incX, X, incX)
	checkSquare(rows, cols)
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	cblas.CBLAS_ZTRMV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), rows, A_, lda, X_, incX)
}

/*

func ZTRSV(uplo Uplo, transA Transpose, diag Diag, A [][]complex128, X []complex128) {



	var A_ unsafe.Pointer = 0

	var X_ unsafe.Pointer = 0

	cblas.CBLAS_ZTRSV(uint32(RowMajor), uint32(uplo), uint32(transA), uint32(diag), N_, A_, lda, X_, incX)
}
*/

// Symmetric matrix-vector multiplication:
// 	Y = alpha*A*X + beta*Y
// Matrices must be allocated with MakeFloat32Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
// A is a symmetric matrix, where uplo (Upper or Lower) determines which part of A is used.
// Every incX'th element of X and incY'th element of Y is used.
func SSYMV(uplo Uplo, alpha float32, A [][]float32, X []float32, incX int, beta float32, Y []float32, incY int) {
	rows, cols, lda := checkSMV(NoTrans, A, X, incX, Y, incY)
	checkSquare(rows, cols)
	var A_ *float32 = &A[0][0]
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	cblas.CBLAS_SSYMV(uint32(RowMajor), uint32(uplo), rows, alpha, A_, lda, X_, incX, beta, Y_, incY)
}

// Computes
// 	A = alpha*X*(Y^T) + A
// Matrices must be allocated with MakeFloat32Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
// Every incX'th element of X and incY'th element of Y is used.
func SGER(alpha float32, X []float32, incX int, Y []float32, incY int, A [][]float32) {
	rows, cols, lda := SSize(A)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, N)
	var X_ *float32 = &X[0]
	var Y_ *float32 = &Y[0]
	var A_ *float32 = &A[0][0]
	cblas.CBLAS_SGER(uint32(RowMajor), M, N, alpha, X_, incX, Y_, incY, A_, lda)
}

// Computes
// 	A = A + alpha*X*(X^T)
// A is a symmetric matrix, where uplo (Upper or Lower) determines which part of A is used.
func SSYR(uplo Uplo, alpha float32, X []float32, incX int, A [][]float32) {
	rows, cols, lda := SSize(A)
	checkSquare(rows, cols)
	N := len(X) / incX
	checkGER(rows, cols, N, N)
	var X_ *float32 = &X[0]
	var A_ *float32 = &A[0][0]
	cblas.CBLAS_SSYR(uint32(RowMajor), uint32(uplo), N, alpha, X_, incX, A_, lda)
}

/*
func SSYR2(uplo Uplo, alpha float32, X []float32, incX int, Y []float32, A [][]float32) {
	var X_ *float32 = 0
	var Y_ *float32 = 0
	var A_ *float32 = &A[0][0]
	cblas.CBLAS_SSYR2(uint32(RowMajor), uint32(uplo), N_, alpha, X_, incX, Y_, incY, A_, lda)
}

*/
// Symmetric matrix-vector multiplication:
// 	Y = alpha*A*X + beta*Y
// A is a symmetric matrix, where uplo (Upper or Lower) determines which part of A is used.
// Every incX'th element of X and incY'th element of Y is used.
func DSYMV(uplo Uplo, alpha float64, A [][]float64, X []float64, incX int, beta float64, Y []float64, incY int) {
	rows, cols, lda := checkDMV(NoTrans, A, X, incX, Y, incY)
	checkSquare(rows, cols)
	var A_ *float64 = &A[0][0]
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	cblas.CBLAS_DSYMV(uint32(RowMajor), uint32(uplo), rows, alpha, A_, lda, X_, incX, beta, Y_, incY)
}

// Computes
// 	A = alpha*X*(Y^T) + A
// Matrices must be allocated with MakeFloat64Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
// Every incX'th element of X and incY'th element of Y is used.
func DGER(alpha float64, X []float64, incX int, Y []float64, incY int, A [][]float64) {
	rows, cols, lda := DSize(A)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, N)
	var X_ *float64 = &X[0]
	var Y_ *float64 = &Y[0]
	var A_ *float64 = &A[0][0]
	cblas.CBLAS_DGER(uint32(RowMajor), M, N, alpha, X_, incX, Y_, incY, A_, lda)
}

// Computes
// 	A = A + alpha*X*(X^T)
// A is a symmetric matrix, where uplo (Upper or Lower) determines which part of A is used.
func DSYR(uplo Uplo, alpha float64, X []float64, incX int, A [][]float64) {
	rows, cols, lda := DSize(A)
	checkSquare(rows, cols)
	N := len(X) / incX
	checkGER(rows, cols, N, N)
	var X_ *float64 = &X[0]
	var A_ *float64 = &A[0][0]
	cblas.CBLAS_DSYR(uint32(RowMajor), uint32(uplo), N, alpha, X_, incX, A_, lda)
}

/*


func DSYR2(uplo Uplo, alpha float64, X []float64, Y []float64, A [][]float64) {



	var X_ *float64 = 0

	var Y_ *float64 = 0

	var A_ *float64 = &A[0][0]

	cblas.CBLAS_DSYR2(uint32(RowMajor), uint32(uplo), N_, alpha, X_, incX, Y_, incY, A_, lda)
}
*/
// Hermitian matrix-vector product:
// 	Y = alpha*X + beta*Y
// Matrices must be allocated with MakeComplex64Matrix to ensure contiguous underlying storage,
// With uplo == Upper or Lower, the upper or lower triangular part of A is used.
// Every incX'th and incY'th element is used.
func CHEMV(uplo Uplo, alpha complex64, A [][]complex64, X []complex64, incX int, beta complex64, Y []complex64, incY int) {
	rows, cols, lda := checkCMV(NoTrans, A, X, incX, Y, incY)
	checkSquare(rows, cols)
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_CHEMV(uint32(RowMajor), uint32(uplo), rows, unsafe.Pointer(&alpha), A_, lda, X_, incX, unsafe.Pointer(&beta), Y_, incY)
}

// Computes the rank-1 update:
// 	A = alpha*X*(Y^T) + A
// Matrices must be allocated with MakeComplex64Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
// Every incX'th element of X and incY'th element of Y is used.
func CGERU(alpha complex64, X []complex64, incX int, Y []complex64, incY int, A [][]complex64) {
	rows, cols, lda := CSize(A)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, N)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_CGERU(uint32(RowMajor), M, N, unsafe.Pointer(&alpha), X_, incX, Y_, incY, A_, lda)
}

// Computes the conjugate rank-1 update:
// 	A = alpha*X*conj(Y^T) + A
// Matrices must be allocated with MakeComplex64Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
// Every incX'th element of X and incY'th element of Y is used.
func CGERC(alpha complex64, X []complex64, incX int, Y []complex64, incY int, A [][]complex64) {
	rows, cols, lda := CSize(A)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, N)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_CGERC(uint32(RowMajor), M, N, unsafe.Pointer(&alpha), X_, incX, Y_, incY, A_, lda)
}

// Computes the hermitian rank-1 update:
// 	A = alpha*X*conj(X^T) + A
// A is a hermitian matrix. Only its upper half or lower half need to be stored,
// specified by uplo=Upper/Lower.
func CHER(uplo Uplo, alpha float32, X []complex64, incX int, A [][]complex64) {
	rows, cols, lda := CSize(A)
	checkSquare(rows, cols)
	var M int = len(X) / incX
	checkGER(rows, cols, M, M)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_CHER(uint32(RowMajor), uint32(uplo), rows, alpha, X_, incX, A_, lda)
}

// Computes the hermitian rank-2 update:
// A = alpha X conj(Y^T) + conj(alpha) Y conj(X^T) + A
// A is a hermitian matrix. Only its upper half or lower half need to be stored,
// specified by uplo=Upper/Lower.
func CHER2(uplo Uplo, alpha complex64, X []complex64, incX int, Y []complex64, incY int, A [][]complex64) {
	rows, cols, lda := CSize(A)
	checkSquare(rows, cols)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, M)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_CHER2(uint32(RowMajor), uint32(uplo), N, unsafe.Pointer(&alpha), X_, incX, Y_, incY, A_, lda)
}

// Hermitian matrix-vector product:
// 	Y = alpha*X + beta*Y
// Matrices must be allocated with MakeComplex128Matrix to ensure contiguous underlying storage,
// With uplo == Upper or Lower, the upper or lower triangular part of A is used.
// Every incX'th and incY'th element is used.
func ZHEMV(uplo Uplo, alpha complex128, A [][]complex128, X []complex128, incX int, beta complex128, Y []complex128, incY int) {
	rows, cols, lda := checkZMV(NoTrans, A, X, incX, Y, incY)
	checkSquare(rows, cols)
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	cblas.CBLAS_ZHEMV(uint32(RowMajor), uint32(uplo), rows, unsafe.Pointer(&alpha), A_, lda, X_, incX, unsafe.Pointer(&beta), Y_, incY)
}

// Computes the rank-1 update:
// 	A = alpha*X*(Y^T) + A
// Matrices must be allocated with MakeComplex128Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
// Every incX'th element of X and incY'th element of Y is used.
func ZGERU(alpha complex128, X []complex128, incX int, Y []complex128, incY int, A [][]complex128) {
	rows, cols, lda := ZSize(A)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, N)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_ZGERU(uint32(RowMajor), M, N, unsafe.Pointer(&alpha), X_, incX, Y_, incY, A_, lda)
}

// Computes the conjugate rank-1 update
// 	A = alpha*X*conj(Y^T) + A
// Matrices must be allocated with MakeComplex128Matrix to ensure contiguous underlying storage,
// otherwise this function may panic or return unexpected results.
// Every incX'th element of X and incY'th element of Y is used.
func ZGERC(alpha complex128, X []complex128, incX int, Y []complex128, incY int, A [][]complex128) {
	rows, cols, lda := ZSize(A)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, N)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&Y[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_ZGERC(uint32(RowMajor), M, N, unsafe.Pointer(&alpha), X_, incX, Y_, incY, A_, lda)
}

// Computes the hermitian rank-1 update:
// 	A = alpha*X*conj(X^T) + A
// A is a hermitian matrix. Only its upper half or lower half need to be stored,
// specified by uplo=Upper/Lower.
func ZHER(uplo Uplo, alpha float64, X []complex128, incX int, A [][]complex128) {
	rows, cols, lda := ZSize(A)
	checkSquare(rows, cols)
	var M int = len(X) / incX
	checkGER(rows, cols, M, M)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_ZHER(uint32(RowMajor), uint32(uplo), rows, alpha, X_, incX, A_, lda)
}

// Computes the hermitian rank-2 update:
// A = alpha X conj(Y^T) + conj(alpha) Y conj(X^T) + A
// A is a hermitian matrix. Only its upper half or lower half need to be stored,
// specified by uplo=Upper/Lower.
func ZHER2(uplo Uplo, alpha complex128, X []complex128, incX int, Y []complex128, incY int, A [][]complex128) {
	rows, cols, lda := ZSize(A)
	checkSquare(rows, cols)
	var M int = len(X) / incX
	var N int = len(Y) / incY
	checkGER(rows, cols, M, M)
	var X_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var Y_ unsafe.Pointer = unsafe.Pointer(&X[0])
	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	cblas.CBLAS_ZHER2(uint32(RowMajor), uint32(uplo), N, unsafe.Pointer(&alpha), X_, incX, Y_, incY, A_, lda)
}

// General matrix-matrix multiplication:
// 	C = alpha*A*B + beta*C
// where A and B can optionally be transposed by specifying transA, transB = Trans/NoTrans
func SGEMM(transA Transpose, transB Transpose, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {
	rowsA, colsA, lda := SSize(A)
	rowsB, colsB, ldb := SSize(B)
	rowsC, colsC, ldc := SSize(C)
	checkMM(transA, transB, rowsA, colsA, rowsB, colsB, rowsC, colsC)

	var A_ *float32 = &A[0][0]
	var B_ *float32 = &B[0][0]
	var C_ *float32 = &C[0][0]

	cblas.CBLAS_SGEMM(uint32(RowMajor), uint32(transA), uint32(transB), rowsC, colsC, colsA, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

// Computes the matrix-matrix product
// 	C = alpha*A*B + beta*C  (for Side==Left)
// 	C = alpha*B*A + beta*C  (for Side==Right)
// where the matrix A is symmetric. Only its upper half or lower half is used, specified by uplo=Upper/Lower.
func SSYMM(side Side, uplo Uplo, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {
	rowsA, colsA, lda := SSize(A)
	rowsB, colsB, ldb := SSize(B)
	rowsC, colsC, ldc := SSize(C)
	checkSquare(rowsA, colsA)
	if side == Left {
		checkMM(NoTrans, NoTrans, rowsA, colsA, rowsB, colsB, rowsC, colsC)
	} else {
		checkMM(NoTrans, NoTrans, rowsB, colsB, rowsA, colsA, rowsC, colsC)
	}

	var A_ *float32 = &A[0][0]
	var B_ *float32 = &B[0][0]
	var C_ *float32 = &C[0][0]

	cblas.CBLAS_SSYMM(uint32(RowMajor), uint32(side), uint32(uplo), rowsC, colsC, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

// Computes a rank-k update of the symmetric matrix C:
// 	C = alpha A*(A^T) + beta C (for trans==NoTrans)
//  C = alpha (A^T)*A + beta C (for trans==Trans)
// Since the matrix C is symmetric only its upper half or lower half need to be stored, specified by uplo=Upper/Lower.
func SSYRK(uplo Uplo, trans Transpose, alpha float32, A [][]float32, beta float32, C [][]float32) {
	rowsA, colsA, lda := SSize(A)
	rowsC, colsC, ldc := SSize(C)
	checkSquare(rowsC, colsC)
	if trans == NoTrans {
		checkMM(NoTrans, Trans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	} else {
		checkMM(Trans, NoTrans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	}

	var A_ *float32 = &A[0][0]
	var C_ *float32 = &C[0][0]
	K := 0
	if trans == NoTrans {
		K = colsA
	} else {
		K = rowsA
	}
	cblas.CBLAS_SSYRK(uint32(RowMajor), uint32(uplo), uint32(trans), rowsC, K, alpha, A_, lda, beta, C_, ldc)
}

/*
func SSYR2K(uplo Uplo, trans Transpose, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ *float32 = &A[0][0]

	var B_ *float32 = 0
	var ldb int = 0
	var beta float32 = 0
	var C_ *float32 = 0
	var ldc int = 0
	cblas.CBLAS_SSYR2K(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

func STRMM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float32, A [][]float32, B [][]float32) {




	var M_ int = 0


	var A_ *float32 = &A[0][0]

	var B_ *float32 = 0
	var ldb int = 0
	cblas.CBLAS_STRMM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}

func STRSM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float32, A [][]float32, B [][]float32) {




	var M_ int = 0


	var A_ *float32 = &A[0][0]

	var B_ *float32 = 0
	var ldb int = 0
	cblas.CBLAS_STRSM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}



*/
// General matrix-matrix multiplication:
// 	C = alpha*A*B + beta*C
// where A and B can optionally be transposed by specifying transA, transB = Trans/NoTrans
func DGEMM(transA Transpose, transB Transpose, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {
	rowsA, colsA, lda := DSize(A)
	rowsB, colsB, ldb := DSize(B)
	rowsC, colsC, ldc := DSize(C)
	checkMM(transA, transB, rowsA, colsA, rowsB, colsB, rowsC, colsC)

	var A_ *float64 = &A[0][0]
	var B_ *float64 = &B[0][0]
	var C_ *float64 = &C[0][0]

	cblas.CBLAS_DGEMM(uint32(RowMajor), uint32(transA), uint32(transB), rowsC, colsC, colsA, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

// Computes the matrix-matrix product
// 	C = alpha*A*B + beta*C  (for Side==Left)
// 	C = alpha*B*A + beta*C  (for Side==Right)
// where the matrix A is symmetric. Only its upper half or lower half is used, specified by uplo=Upper/Lower.
func DSYMM(side Side, uplo Uplo, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {
	rowsA, colsA, lda := DSize(A)
	rowsB, colsB, ldb := DSize(B)
	rowsC, colsC, ldc := DSize(C)
	checkSquare(rowsA, colsA)
	if side == Left {
		checkMM(NoTrans, NoTrans, rowsA, colsA, rowsB, colsB, rowsC, colsC)
	} else {
		checkMM(NoTrans, NoTrans, rowsB, colsB, rowsA, colsA, rowsC, colsC)
	}

	var A_ *float64 = &A[0][0]
	var B_ *float64 = &B[0][0]
	var C_ *float64 = &C[0][0]

	cblas.CBLAS_DSYMM(uint32(RowMajor), uint32(side), uint32(uplo), rowsC, colsC, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

// Computes a rank-k update of the symmetric matrix C:
// 	C = alpha A*(A^T) + beta C (for trans==NoTrans)
//  C = alpha (A^T)*A + beta C (for trans==Trans)
// Since the matrix C is symmetric only its upper half or lower half need to be stored, specified by uplo=Upper/Lower.
func DSYRK(uplo Uplo, trans Transpose, alpha float64, A [][]float64, beta float64, C [][]float64) {
	rowsA, colsA, lda := DSize(A)
	rowsC, colsC, ldc := DSize(C)
	checkSquare(rowsC, colsC)
	if trans == NoTrans {
		checkMM(NoTrans, Trans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	} else {
		checkMM(Trans, NoTrans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	}

	var A_ *float64 = &A[0][0]
	var C_ *float64 = &C[0][0]
	K := 0
	if trans == NoTrans {
		K = colsA
	} else {
		K = rowsA
	}
	cblas.CBLAS_DSYRK(uint32(RowMajor), uint32(uplo), uint32(trans), rowsC, K, alpha, A_, lda, beta, C_, ldc)
}

/*

func DSYR2K(uplo Uplo, trans Transpose, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ *float64 = &A[0][0]

	var B_ *float64 = 0
	var ldb int = 0
	var beta float64 = 0
	var C_ *float64 = 0
	var ldc int = 0
	cblas.CBLAS_DSYR2K(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

func DTRMM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float64, A [][]float64, B [][]float64) {




	var M_ int = 0


	var A_ *float64 = &A[0][0]

	var B_ *float64 = 0
	var ldb int = 0
	cblas.CBLAS_DTRMM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}

func DTRSM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha float64, A [][]float64, B [][]float64) {




	var M_ int = 0


	var A_ *float64 = &A[0][0]

	var B_ *float64 = 0
	var ldb int = 0
	cblas.CBLAS_DTRSM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}
*/

// General matrix-matrix multiplication:
// 	C = alpha*A*B + beta*C
// where A and B can optionally be transposed by specifying transA, transB = Trans/NoTrans
func CGEMM(transA Transpose, transB Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	rowsA, colsA, lda := CSize(A)
	rowsB, colsB, ldb := CSize(B)
	rowsC, colsC, ldc := CSize(C)
	checkMM(transA, transB, rowsA, colsA, rowsB, colsB, rowsC, colsC)

	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var B_ unsafe.Pointer = unsafe.Pointer(&B[0][0])
	var C_ unsafe.Pointer = unsafe.Pointer(&C[0][0])

	cblas.CBLAS_CGEMM(uint32(RowMajor), uint32(transA), uint32(transB), rowsC, colsC, colsA, unsafe.Pointer(&alpha), A_, lda, B_, ldb, unsafe.Pointer(&beta), C_, ldc)
}

// Computes the matrix-matrix product
// 	C = alpha*A*B + beta*C  (for Side==Left)
// 	C = alpha*B*A + beta*C  (for Side==Right)
// where the matrix A is symmetric. Only its upper half or lower half is used, specified by uplo=Upper/Lower.
func CSYMM(side Side, uplo Uplo, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	rowsA, colsA, lda := CSize(A)
	rowsB, colsB, ldb := CSize(B)
	rowsC, colsC, ldc := CSize(C)
	checkSquare(rowsA, colsA)
	if side == Left {
		checkMM(NoTrans, NoTrans, rowsA, colsA, rowsB, colsB, rowsC, colsC)
	} else {
		checkMM(NoTrans, NoTrans, rowsB, colsB, rowsA, colsA, rowsC, colsC)
	}

	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var B_ unsafe.Pointer = unsafe.Pointer(&B[0][0])
	var C_ unsafe.Pointer = unsafe.Pointer(&C[0][0])

	cblas.CBLAS_CSYMM(uint32(RowMajor), uint32(side), uint32(uplo), rowsC, colsC, unsafe.Pointer(&alpha), A_, lda, B_, ldb, unsafe.Pointer(&beta), C_, ldc)
}

// Computes a rank-k update of the symmetric matrix C:
// 	C = alpha A*(A^T) + beta C (for trans==NoTrans)
//  C = alpha (A^T)*A + beta C (for trans==Trans)
// Since the matrix C is symmetric only its upper half or lower half need to be stored, specified by uplo=Upper/Lower.
func CSYRK(uplo Uplo, trans Transpose, alpha complex64, A [][]complex64, beta complex64, C [][]complex64) {
	rowsA, colsA, lda := CSize(A)
	rowsC, colsC, ldc := CSize(C)
	checkSquare(rowsC, colsC)
	if trans == NoTrans {
		checkMM(NoTrans, Trans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	} else {
		checkMM(Trans, NoTrans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	}

	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var C_ unsafe.Pointer = unsafe.Pointer(&C[0][0])
	K := 0
	if trans == NoTrans {
		K = colsA
	} else {
		K = rowsA
	}
	cblas.CBLAS_CSYRK(uint32(RowMajor), uint32(uplo), uint32(trans), rowsC, K, unsafe.Pointer(&alpha), A_, lda, unsafe.Pointer(&beta), C_, ldc)
}

/*



func CSYR2K(uplo Uplo, trans Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	var beta unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_CSYR2K(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

func CTRMM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex64, A [][]complex64, B [][]complex64) {




	var M_ int = 0


	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	cblas.CBLAS_CTRMM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}

func CTRSM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex64, A [][]complex64, B [][]complex64) {




	var M_ int = 0


	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	cblas.CBLAS_CTRSM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}
*/
// General matrix-matrix multiplication:
// 	C = alpha*A*B + beta*C
// where A and B can optionally be transposed by specifying transA, transB = Trans/NoTrans
func ZGEMM(transA Transpose, transB Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	rowsA, colsA, lda := ZSize(A)
	rowsB, colsB, ldb := ZSize(B)
	rowsC, colsC, ldc := ZSize(C)
	checkMM(transA, transB, rowsA, colsA, rowsB, colsB, rowsC, colsC)

	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var B_ unsafe.Pointer = unsafe.Pointer(&B[0][0])
	var C_ unsafe.Pointer = unsafe.Pointer(&C[0][0])

	cblas.CBLAS_ZGEMM(uint32(RowMajor), uint32(transA), uint32(transB), rowsC, colsC, colsA, unsafe.Pointer(&alpha), A_, lda, B_, ldb, unsafe.Pointer(&beta), C_, ldc)
}

// Computes the matrix-matrix product
// 	C = alpha*A*B + beta*C  (for Side==Left)
// 	C = alpha*B*A + beta*C  (for Side==Right)
// where the matrix A is symmetric. Only its upper half or lower half is used, specified by uplo=Upper/Lower.
func ZSYMM(side Side, uplo Uplo, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	rowsA, colsA, lda := ZSize(A)
	rowsB, colsB, ldb := ZSize(B)
	rowsC, colsC, ldc := ZSize(C)
	checkSquare(rowsA, colsA)
	if side == Left {
		checkMM(NoTrans, NoTrans, rowsA, colsA, rowsB, colsB, rowsC, colsC)
	} else {
		checkMM(NoTrans, NoTrans, rowsB, colsB, rowsA, colsA, rowsC, colsC)
	}

	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var B_ unsafe.Pointer = unsafe.Pointer(&B[0][0])
	var C_ unsafe.Pointer = unsafe.Pointer(&C[0][0])

	cblas.CBLAS_ZSYMM(uint32(RowMajor), uint32(side), uint32(uplo), rowsC, colsC, unsafe.Pointer(&alpha), A_, lda, B_, ldb, unsafe.Pointer(&beta), C_, ldc)
}

// Computes a rank-k update of the symmetric matrix C:
// 	C = alpha A*(A^T) + beta C (for trans==NoTrans)
//  C = alpha (A^T)*A + beta C (for trans==Trans)
// Since the matrix C is symmetric only its upper half or lower half need to be stored, specified by uplo=Upper/Lower.
func ZSYRK(uplo Uplo, trans Transpose, alpha complex128, A [][]complex128, beta complex128, C [][]complex128) {
	rowsA, colsA, lda := ZSize(A)
	rowsC, colsC, ldc := ZSize(C)
	checkSquare(rowsC, colsC)
	if trans == NoTrans {
		checkMM(NoTrans, Trans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	} else {
		checkMM(Trans, NoTrans, rowsA, colsA, rowsA, colsA, rowsC, colsC)
	}

	var A_ unsafe.Pointer = unsafe.Pointer(&A[0][0])
	var C_ unsafe.Pointer = unsafe.Pointer(&C[0][0])
	K := 0
	if trans == NoTrans {
		K = colsA
	} else {
		K = rowsA
	}
	cblas.CBLAS_ZSYRK(uint32(RowMajor), uint32(uplo), uint32(trans), rowsC, K, unsafe.Pointer(&alpha), A_, lda, unsafe.Pointer(&beta), C_, ldc)
}

/*


func ZSYR2K(uplo Uplo, trans Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	var beta unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_ZSYR2K(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

func ZTRMM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex128, A [][]complex128, B [][]complex128) {




	var M_ int = 0


	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	cblas.CBLAS_ZTRMM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}

func ZTRSM(side Side, uplo Uplo, transA Transpose, diag Diag, alpha complex128, A [][]complex128, B [][]complex128) {




	var M_ int = 0


	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	cblas.CBLAS_ZTRSM(uint32(RowMajor), uint32(side), uint32(uplo), uint32(transA), uint32(diag), M_, N_, alpha, A_, lda, B_, ldb)
}

func CHEMM(side Side, uplo Uplo, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {


	var M_ int = 0


	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	var beta unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_CHEMM(uint32(RowMajor), uint32(side), uint32(uplo), M_, N_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

func CHERK(uplo Uplo, trans Transpose, alpha float32, A [][]complex64, beta float32, C [][]complex64) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ unsafe.Pointer = 0

	var beta float32 = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_CHERK(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, beta, C_, ldc)
}

func CHER2K(uplo Uplo, trans Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta float32, C [][]complex64) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	var beta float32 = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_CHER2K(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

func ZHEMM(side Side, uplo Uplo, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {


	var M_ int = 0


	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	var beta unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_ZHEMM(uint32(RowMajor), uint32(side), uint32(uplo), M_, N_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}

func ZHERK(uplo Uplo, trans Transpose, alpha float64, A [][]complex128, beta float64, C [][]complex128) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ unsafe.Pointer = 0

	var beta float64 = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_ZHERK(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, beta, C_, ldc)
}

func ZHER2K(uplo Uplo, trans Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta float64, C [][]complex128) {

	var uint32(trans) uint32 = 0

	var K_ int = 0

	var A_ unsafe.Pointer = 0

	var B_ unsafe.Pointer = 0
	var ldb int = 0
	var beta float64 = 0
	var C_ unsafe.Pointer = 0
	var ldc int = 0
	cblas.CBLAS_ZHER2K(uint32(RowMajor), uint32(uplo), uint32(trans), N_, K_, alpha, A_, lda, B_, ldb, beta, C_, ldc)
}
*/
