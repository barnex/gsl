package blas

//#include "../cblas/gsl_gsl_cblas.h"
import "C"

type Order uint32 //Indicates whether a matrix is in Row Major or Column Major order.
const (
	RowMajor Order = C.CblasRowMajor // Row major order is the native order for C and Go programs.
	ColMajor Order = C.CblasColMajor // Column major order is native for Fortran.
)

type Transpose uint32 // Used to represent transpose operations on a matrix.
const (
	NoTrans   Transpose = C.CblasNoTrans   // NoTrans represents X.
	Trans     Transpose = C.CblasTrans     // Trans represents X^T.
	ConjTrans Transpose = C.CblasConjTrans // ConjTrans represents X^H (=conj(X^T))
)

type Uplo uint32 // Used to indicate which part of a symmetric matrix to use.
const (
	Upper Uplo = C.CblasUpper // Upper means use the upper triangle of the matrix.
	Lower Uplo = C.CblasLower // Lower means use the lower triangle of the matrix.
)

type Diag uint32 // Used to indicate whether a triangular matrix is unit-diagonal (diagonal elements are all equal to 1).
const (
	NonUnit Diag = C.CblasNonUnit // NonUnit means non unit-diagonal matrix.
	Unit    Diag = C.CblasUnit    //  Unit means unit-diagonal matrix.
)

type Side uint32 // Used to indicate the order of a matrix-matrix multiply.
const (
	Left  Side = C.CblasLeft  // Left means AB matrix-matrix multiply.
	Right Side = C.CblasRight // Right means BA matrix-matrix multiply.
)
