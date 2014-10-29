package blas

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

import (
	"github.com/barnex/blas/cblas"
)

func SDSDOT(alpha float32, X []float32, incX int, Y []float32, incY int) float32 {
	checkSizeS(X, Y)
	var N_ int = len(X)
	var X_ *float32 = &X[0] 
	var Y_ *float32 = &Y[0]
	return cblas.CBLAS_SDSDOT(N_, alpha, X_, incX, Y_, incY)
}

/*
func DSDOT(X []float32, Y []float32) float64 {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_DSDOT(N_, X_, incX_, Y_, incY_)
}

func SDOT(X []float32, Y []float32) float32 {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_SDOT(N_, X_, incX_, Y_, incY_)
}

func DDOT(X []float64, Y []float64) float64 {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DDOT(N_, X_, incX_, Y_, incY_)
}

func CDOTU(X []complex64, Y []complex64, dotu *complex64) {

}

func CDOTC(X []complex64, Y []complex64, dotc *complex64) {

}

func ZDOTU(X []complex128, Y []complex128, dotu *complex128) {

}

func ZDOTC(X []complex128, Y []complex128, dotc *complex128) {

}

func SNRM2(X []float32) float32 {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_SNRM2(N_, X_, incX_)
}

func SASUM(X []float32) float32 {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_SASUM(N_, X_, incX_)
}

func DNRM2(X []float64) float64 {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_DNRM2(N_, X_, incX_)
}

func DASUM(X []float64) float64 {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_DASUM(N_, X_, incX_)
}

func SCNRM2(X []complex64) float32 {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_SCNRM2(N_, X_, incX_)
}

func SCASUM(X []complex64) float32 {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_SCASUM(N_, X_, incX_)
}

func DZNRM2(X []complex128) float64 {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_DZNRM2(N_, X_, incX_)
}

func DZASUM(X []complex128) float64 {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_DZASUM(N_, X_, incX_)
}

func ISAMAX(X []float32) int {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_ISAMAX(N_, X_, incX_)
}

func IDAMAX(X []float64) int {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_IDAMAX(N_, X_, incX_)
}

func ICAMAX(X []complex64) int {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_ICAMAX(N_, X_, incX_)
}

func IZAMAX(X []complex128) int {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_IZAMAX(N_, X_, incX_)
}

func SSWAP(X []float32, Y []float32) {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_SSWAP(N_, X_, incX_, Y_, incY_)
}

func SCOPY(X []float32, Y []float32) {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_SCOPY(N_, X_, incX_, Y_, incY_)
}

func SAXPY(alpha float32, X []float32, Y []float32) {
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_SAXPY(N_, alpha_, X_, incX_, Y_, incY_)
}

func DSWAP(X []float64, Y []float64) {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DSWAP(N_, X_, incX_, Y_, incY_)
}

func DCOPY(X []float64, Y []float64) {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DCOPY(N_, X_, incX_, Y_, incY_)
}

func DAXPY(alpha float64, X []float64, Y []float64) {
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DAXPY(N_, alpha_, X_, incX_, Y_, incY_)
}

func CSWAP(X []complex64, Y []complex64) {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_CSWAP(N_, X_, incX_, Y_, incY_)
}

func CCOPY(X []complex64, Y []complex64) {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_CCOPY(N_, X_, incX_, Y_, incY_)
}

func CAXPY(alpha complex64, X []complex64, Y []complex64) {
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_CAXPY(N_, alpha_, X_, incX_, Y_, incY_)
}

func ZSWAP(X []complex128, Y []complex128) {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_ZSWAP(N_, X_, incX_, Y_, incY_)
}

func ZCOPY(X []complex128, Y []complex128) {
	var N_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_ZCOPY(N_, X_, incX_, Y_, incY_)
}

func ZAXPY(alpha complex128, X []complex128, Y []complex128) {
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	cblas.CBLAS_ZAXPY(N_, alpha_, X_, incX_, Y_, incY_)
}

func SROTG(a *float32, b *float32, c *float32, s *float32) {
	var a_ *float32 = 0
	var b_ *float32 = 0
	var c_ *float32 = 0
	var s_ *float32 = 0
	cblas.CBLAS_SROTG(a_, b_, c_, s_)
}

func SROTMG(d1 *float32, d2 *float32, b1 *float32, b2 float32, P *float32) {
	var d1_ *float32 = 0
	var d2_ *float32 = 0
	var b1_ *float32 = 0
	var b2_ float32 = 0
	var P_ *float32 = 0
	cblas.CBLAS_SROTMG(d1_, d2_, b1_, b2_, P_)
}

func SROT(X []float32, Y []float32, c float32, s float32) {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	var c_ float32 = 0
	var s_ float32 = 0
	cblas.CBLAS_SROT(N_, X_, incX_, Y_, incY_, c_, s_)
}

func SROTM(X []float32, Y []float32, P *float32) {
	var N_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	var P_ *float32 = 0
	cblas.CBLAS_SROTM(N_, X_, incX_, Y_, incY_, P_)
}

func DROTG(a *float64, b *float64, c *float64, s *float64) {
	var a_ *float64 = 0
	var b_ *float64 = 0
	var c_ *float64 = 0
	var s_ *float64 = 0
	cblas.CBLAS_DROTG(a_, b_, c_, s_)
}

func DROTMG(d1 *float64, d2 *float64, b1 *float64, b2 float64, P *float64) {
	var d1_ *float64 = 0
	var d2_ *float64 = 0
	var b1_ *float64 = 0
	var b2_ float64 = 0
	var P_ *float64 = 0
	cblas.CBLAS_DROTMG(d1_, d2_, b1_, b2_, P_)
}

func DROT(X []float64, Y []float64, c float64, s float64) {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	var c_ float64 = 0
	var s_ float64 = 0
	cblas.CBLAS_DROT(N_, X_, incX_, Y_, incY_, c_, s_)
}

func DROTM(X []float64, Y []float64, P *float64) {
	var N_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	var P_ *float64 = 0
	cblas.CBLAS_DROTM(N_, X_, incX_, Y_, incY_, P_)
}

func SSCAL(alpha float32, X []float32) {
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_SSCAL(N_, alpha_, X_, incX_)
}

func DSCAL(alpha float64, X []float64) {
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_DSCAL(N_, alpha_, X_, incX_)
}

func CSCAL(alpha complex64, X []complex64) {
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_CSCAL(N_, alpha_, X_, incX_)
}

func ZSCAL(alpha complex128, X []complex128) {
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_ZSCAL(N_, alpha_, X_, incX_)
}

func CSSCAL(alpha float32, X []complex64) {
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_CSSCAL(N_, alpha_, X_, incX_)
}

func ZDSCAL(alpha float64, X []complex128) {
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_ZDSCAL(N_, alpha_, X_, incX_)
}

func SGEMV(TransA Transpose, alpha float32, A [][]float32, X []float32, beta float32, Y []float32) {
	var order_ uint32 = 0
	var TransA_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var beta_ float32 = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_SGEMV(order_, TransA_, M_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func STRMV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]float32, X []float32) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_STRMV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func STRSV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]float32, X []float32) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	cblas.CBLAS_STRSV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func DGEMV(TransA Transpose, alpha float64, A [][]float64, X []float64, beta float64, Y []float64) {
	var order_ uint32 = 0
	var TransA_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var beta_ float64 = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DGEMV(order_, TransA_, M_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func DTRMV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]float64, X []float64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_DTRMV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func DTRSV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]float64, X []float64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	cblas.CBLAS_DTRSV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func CGEMV(TransA Transpose, alpha complex64, A [][]complex64, X []complex64, beta complex64, Y []complex64) {
	var order_ uint32 = 0
	var TransA_ uint32 = 0
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
	cblas.CBLAS_CGEMV(order_, TransA_, M_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func CTRMV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]complex64, X []complex64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_CTRMV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func CTRSV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]complex64, X []complex64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_CTRSV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func ZGEMV(TransA Transpose, alpha complex128, A [][]complex128, X []complex128, beta complex128, Y []complex128) {
	var order_ uint32 = 0
	var TransA_ uint32 = 0
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
	cblas.CBLAS_ZGEMV(order_, TransA_, M_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func ZTRMV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]complex128, X []complex128) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_ZTRMV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func ZTRSV(Uplo Uplo, TransA Transpose, Diag Diag, A [][]complex128, X []complex128) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var N_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	cblas.CBLAS_ZTRSV(order_, Uplo_, TransA_, Diag_, N_, A_, lda_, X_, incX_)
}

func SSYMV(Uplo Uplo, alpha float32, A [][]float32, X []float32, beta float32, Y []float32) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var beta_ float32 = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	cblas.CBLAS_SSYMV(order_, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func SGER(alpha float32, X []float32, Y []float32, A [][]float32) {
	var order_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	var A_ *float32 = 0
	var lda_ int = 0
	cblas.CBLAS_SGER(order_, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func SSYR(Uplo Uplo, alpha float32, X []float32, A [][]float32) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var A_ *float32 = 0
	var lda_ int = 0
	cblas.CBLAS_SSYR(order_, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func SSYR2(Uplo Uplo, alpha float32, X []float32, Y []float32, A [][]float32) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ *float32 = 0
	var incX_ int = 0
	var Y_ *float32 = 0
	var incY_ int = 0
	var A_ *float32 = 0
	var lda_ int = 0
	cblas.CBLAS_SSYR2(order_, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func DSYMV(Uplo Uplo, alpha float64, A [][]float64, X []float64, beta float64, Y []float64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var beta_ float64 = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	cblas.CBLAS_DSYMV(order_, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func DGER(alpha float64, X []float64, Y []float64, A [][]float64) {
	var order_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	var A_ *float64 = 0
	var lda_ int = 0
	cblas.CBLAS_DGER(order_, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func DSYR(Uplo Uplo, alpha float64, X []float64, A [][]float64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var A_ *float64 = 0
	var lda_ int = 0
	cblas.CBLAS_DSYR(order_, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func DSYR2(Uplo Uplo, alpha float64, X []float64, Y []float64, A [][]float64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ *float64 = 0
	var incX_ int = 0
	var Y_ *float64 = 0
	var incY_ int = 0
	var A_ *float64 = 0
	var lda_ int = 0
	cblas.CBLAS_DSYR2(order_, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func CHEMV(Uplo Uplo, alpha complex64, A [][]complex64, X []complex64, beta complex64, Y []complex64) {
	var order_ uint32 = 0
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
	cblas.CBLAS_CHEMV(order_, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func CGERU(alpha complex64, X []complex64, Y []complex64, A [][]complex64) {
	var order_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CGERU(order_, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func CGERC(alpha complex64, X []complex64, Y []complex64, A [][]complex64) {
	var order_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CGERC(order_, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func CHER(Uplo Uplo, alpha float32, X []complex64, A [][]complex64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CHER(order_, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func CHER2(Uplo Uplo, alpha complex64, X []complex64, Y []complex64, A [][]complex64) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_CHER2(order_, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func ZHEMV(Uplo Uplo, alpha complex128, A [][]complex128, X []complex128, beta complex128, Y []complex128) {
	var order_ uint32 = 0
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
	cblas.CBLAS_ZHEMV(order_, Uplo_, N_, alpha_, A_, lda_, X_, incX_, beta_, Y_, incY_)
}

func ZGERU(alpha complex128, X []complex128, Y []complex128, A [][]complex128) {
	var order_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZGERU(order_, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func ZGERC(alpha complex128, X []complex128, Y []complex128, A [][]complex128) {
	var order_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZGERC(order_, M_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func ZHER(Uplo Uplo, alpha float64, X []complex128, A [][]complex128) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZHER(order_, Uplo_, N_, alpha_, X_, incX_, A_, lda_)
}

func ZHER2(Uplo Uplo, alpha complex128, X []complex128, Y []complex128, A [][]complex128) {
	var order_ uint32 = 0
	var Uplo_ uint32 = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var X_ unsafe.Pointer = 0
	var incX_ int = 0
	var Y_ unsafe.Pointer = 0
	var incY_ int = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	cblas.CBLAS_ZHER2(order_, Uplo_, N_, alpha_, X_, incX_, Y_, incY_, A_, lda_)
}

func SGEMM(TransA Transpose, TransB Transpose, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {
	var Order_ uint32 = 0
	var TransA_ uint32 = 0
	var TransB_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SGEMM(Order_, TransA_, TransB_, M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func SSYMM(Side Side, Uplo Uplo, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SSYMM(Order_, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func SSYRK(Uplo Uplo, Trans Transpose, alpha float32, A [][]float32, beta float32, C [][]float32) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SSYRK(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func SSYR2K(Uplo Uplo, Trans Transpose, alpha float32, A [][]float32, B [][]float32, beta float32, C [][]float32) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	var beta_ float32 = 0
	var C_ *float32 = 0
	var ldc_ int = 0
	cblas.CBLAS_SSYR2K(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func STRMM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha float32, A [][]float32, B [][]float32) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	cblas.CBLAS_STRMM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func STRSM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha float32, A [][]float32, B [][]float32) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float32 = 0
	var A_ *float32 = 0
	var lda_ int = 0
	var B_ *float32 = 0
	var ldb_ int = 0
	cblas.CBLAS_STRSM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func DGEMM(TransA Transpose, TransB Transpose, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {
	var Order_ uint32 = 0
	var TransA_ uint32 = 0
	var TransB_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DGEMM(Order_, TransA_, TransB_, M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func DSYMM(Side Side, Uplo Uplo, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DSYMM(Order_, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func DSYRK(Uplo Uplo, Trans Transpose, alpha float64, A [][]float64, beta float64, C [][]float64) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DSYRK(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func DSYR2K(Uplo Uplo, Trans Transpose, alpha float64, A [][]float64, B [][]float64, beta float64, C [][]float64) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	var beta_ float64 = 0
	var C_ *float64 = 0
	var ldc_ int = 0
	cblas.CBLAS_DSYR2K(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func DTRMM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha float64, A [][]float64, B [][]float64) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	cblas.CBLAS_DTRMM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func DTRSM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha float64, A [][]float64, B [][]float64) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ float64 = 0
	var A_ *float64 = 0
	var lda_ int = 0
	var B_ *float64 = 0
	var ldb_ int = 0
	cblas.CBLAS_DTRSM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func CGEMM(TransA Transpose, TransB Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	var Order_ uint32 = 0
	var TransA_ uint32 = 0
	var TransB_ uint32 = 0
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
	cblas.CBLAS_CGEMM(Order_, TransA_, TransB_, M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CSYMM(Side Side, Uplo Uplo, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	var Order_ uint32 = 0
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
	cblas.CBLAS_CSYMM(Order_, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CSYRK(Uplo Uplo, Trans Transpose, alpha complex64, A [][]complex64, beta complex64, C [][]complex64) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CSYRK(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func CSYR2K(Uplo Uplo, Trans Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
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
	cblas.CBLAS_CSYR2K(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CTRMM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha complex64, A [][]complex64, B [][]complex64) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_CTRMM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func CTRSM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha complex64, A [][]complex64, B [][]complex64) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_CTRSM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func ZGEMM(TransA Transpose, TransB Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	var Order_ uint32 = 0
	var TransA_ uint32 = 0
	var TransB_ uint32 = 0
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
	cblas.CBLAS_ZGEMM(Order_, TransA_, TransB_, M_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZSYMM(Side Side, Uplo Uplo, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	var Order_ uint32 = 0
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
	cblas.CBLAS_ZSYMM(Order_, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZSYRK(Uplo Uplo, Trans Transpose, alpha complex128, A [][]complex128, beta complex128, C [][]complex128) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ unsafe.Pointer = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZSYRK(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func ZSYR2K(Uplo Uplo, Trans Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
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
	cblas.CBLAS_ZSYR2K(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZTRMM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha complex128, A [][]complex128, B [][]complex128) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_ZTRMM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func ZTRSM(Side Side, Uplo Uplo, TransA Transpose, Diag Diag, alpha complex128, A [][]complex128, B [][]complex128) {
	var Order_ uint32 = 0
	var Side_ uint32 = 0
	var Uplo_ uint32 = 0
	var TransA_ uint32 = 0
	var Diag_ uint32 = 0
	var M_ int = 0
	var N_ int = 0
	var alpha_ unsafe.Pointer = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var B_ unsafe.Pointer = 0
	var ldb_ int = 0
	cblas.CBLAS_ZTRSM(Order_, Side_, Uplo_, TransA_, Diag_, M_, N_, alpha_, A_, lda_, B_, ldb_)
}

func CHEMM(Side Side, Uplo Uplo, alpha complex64, A [][]complex64, B [][]complex64, beta complex64, C [][]complex64) {
	var Order_ uint32 = 0
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
	cblas.CBLAS_CHEMM(Order_, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func CHERK(Uplo Uplo, Trans Transpose, alpha float32, A [][]complex64, beta float32, C [][]complex64) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float32 = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ float32 = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_CHERK(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func CHER2K(Uplo Uplo, Trans Transpose, alpha complex64, A [][]complex64, B [][]complex64, beta float32, C [][]complex64) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
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
	cblas.CBLAS_CHER2K(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZHEMM(Side Side, Uplo Uplo, alpha complex128, A [][]complex128, B [][]complex128, beta complex128, C [][]complex128) {
	var Order_ uint32 = 0
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
	cblas.CBLAS_ZHEMM(Order_, Side_, Uplo_, M_, N_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}

func ZHERK(Uplo Uplo, Trans Transpose, alpha float64, A [][]complex128, beta float64, C [][]complex128) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
	var N_ int = 0
	var K_ int = 0
	var alpha_ float64 = 0
	var A_ unsafe.Pointer = 0
	var lda_ int = 0
	var beta_ float64 = 0
	var C_ unsafe.Pointer = 0
	var ldc_ int = 0
	cblas.CBLAS_ZHERK(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, beta_, C_, ldc_)
}

func ZHER2K(Uplo Uplo, Trans Transpose, alpha complex128, A [][]complex128, B [][]complex128, beta float64, C [][]complex128) {
	var Order_ uint32 = 0
	var Uplo_ uint32 = 0
	var Trans_ uint32 = 0
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
	cblas.CBLAS_ZHER2K(Order_, Uplo_, Trans_, N_, K_, alpha_, A_, lda_, B_, ldb_, beta_, C_, ldc_)
}
*/