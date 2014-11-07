//Package cblas provides an unsafe, low-level C interface used by package blas.
//Users should use package blas directly.
package internal

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

//#cgo LDFLAGS: -lm
//#cgo CFLAGS: -O0
//#include "gsl_gsl_cblas.h"
import "C"

import (
	"unsafe"
)

func CBLAS_SDSDOT(N int, alpha float32, X *float32, incX int, Y *float32, incY int) float32 {
	return (float32)(C.cblas_sdsdot((C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY)))
}

func CBLAS_DSDOT(N int, X *float32, incX int, Y *float32, incY int) float64 {
	return (float64)(C.cblas_dsdot((C.int)(N), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY)))
}

func CBLAS_SDOT(N int, X *float32, incX int, Y *float32, incY int) float32 {
	return (float32)(C.cblas_sdot((C.int)(N), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY)))
}

func CBLAS_DDOT(N int, X *float64, incX int, Y *float64, incY int) float64 {
	return (float64)(C.cblas_ddot((C.int)(N), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY)))
}

func CBLAS_CDOTU_SUB(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, dotu unsafe.Pointer) {
	C.cblas_cdotu_sub((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(dotu))
}

func CBLAS_CDOTC_SUB(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, dotc unsafe.Pointer) {
	C.cblas_cdotc_sub((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(dotc))
}

func CBLAS_ZDOTU_SUB(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, dotu unsafe.Pointer) {
	C.cblas_zdotu_sub((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(dotu))
}

func CBLAS_ZDOTC_SUB(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, dotc unsafe.Pointer) {
	C.cblas_zdotc_sub((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(dotc))
}

func CBLAS_SNRM2(N int, X *float32, incX int) float32 {
	return (float32)(C.cblas_snrm2((C.int)(N), (*C.float)(X), (C.int)(incX)))
}

func CBLAS_SASUM(N int, X *float32, incX int) float32 {
	return (float32)(C.cblas_sasum((C.int)(N), (*C.float)(X), (C.int)(incX)))
}

func CBLAS_DNRM2(N int, X *float64, incX int) float64 {
	return (float64)(C.cblas_dnrm2((C.int)(N), (*C.double)(X), (C.int)(incX)))
}

func CBLAS_DASUM(N int, X *float64, incX int) float64 {
	return (float64)(C.cblas_dasum((C.int)(N), (*C.double)(X), (C.int)(incX)))
}

func CBLAS_SCNRM2(N int, X unsafe.Pointer, incX int) float32 {
	return (float32)(C.cblas_scnrm2((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX)))
}

func CBLAS_SCASUM(N int, X unsafe.Pointer, incX int) float32 {
	return (float32)(C.cblas_scasum((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX)))
}

func CBLAS_DZNRM2(N int, X unsafe.Pointer, incX int) float64 {
	return (float64)(C.cblas_dznrm2((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX)))
}

func CBLAS_DZASUM(N int, X unsafe.Pointer, incX int) float64 {
	return (float64)(C.cblas_dzasum((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX)))
}

func CBLAS_ISAMAX(N int, X *float32, incX int) int {
	return (int)(C.cblas_isamax((C.int)(N), (*C.float)(X), (C.int)(incX)))
}

func CBLAS_IDAMAX(N int, X *float64, incX int) int {
	return (int)(C.cblas_idamax((C.int)(N), (*C.double)(X), (C.int)(incX)))
}

func CBLAS_ICAMAX(N int, X unsafe.Pointer, incX int) int {
	return (int)(C.cblas_icamax((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX)))
}

func CBLAS_IZAMAX(N int, X unsafe.Pointer, incX int) int {
	return (int)(C.cblas_izamax((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX)))
}

func CBLAS_SSWAP(N int, X *float32, incX int, Y *float32, incY int) {
	C.cblas_sswap((C.int)(N), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_SCOPY(N int, X *float32, incX int, Y *float32, incY int) {
	C.cblas_scopy((C.int)(N), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_SAXPY(N int, alpha float32, X *float32, incX int, Y *float32, incY int) {
	C.cblas_saxpy((C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_DSWAP(N int, X *float64, incX int, Y *float64, incY int) {
	C.cblas_dswap((C.int)(N), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_DCOPY(N int, X *float64, incX int, Y *float64, incY int) {
	C.cblas_dcopy((C.int)(N), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_DAXPY(N int, alpha float64, X *float64, incX int, Y *float64, incY int) {
	C.cblas_daxpy((C.int)(N), (C.double)(alpha), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_CSWAP(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int) {
	C.cblas_cswap((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_CCOPY(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int) {
	C.cblas_ccopy((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_CAXPY(N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int) {
	C.cblas_caxpy((C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZSWAP(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int) {
	C.cblas_zswap((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZCOPY(N int, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int) {
	C.cblas_zcopy((C.int)(N), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZAXPY(N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int) {
	C.cblas_zaxpy((C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_SROTG(a *float32, b *float32, c *float32, s *float32) {
	C.cblas_srotg((*C.float)(a), (*C.float)(b), (*C.float)(c), (*C.float)(s))
}

func CBLAS_SROTMG(d1 *float32, d2 *float32, b1 *float32, b2 float32, P *float32) {
	C.cblas_srotmg((*C.float)(d1), (*C.float)(d2), (*C.float)(b1), (C.float)(b2), (*C.float)(P))
}

func CBLAS_SROT(N int, X *float32, incX int, Y *float32, incY int, c float32, s float32) {
	C.cblas_srot((C.int)(N), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY), (C.float)(c), (C.float)(s))
}

func CBLAS_SROTM(N int, X *float32, incX int, Y *float32, incY int, P *float32) {
	C.cblas_srotm((C.int)(N), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY), (*C.float)(P))
}

func CBLAS_DROTG(a *float64, b *float64, c *float64, s *float64) {
	C.cblas_drotg((*C.double)(a), (*C.double)(b), (*C.double)(c), (*C.double)(s))
}

func CBLAS_DROTMG(d1 *float64, d2 *float64, b1 *float64, b2 float64, P *float64) {
	C.cblas_drotmg((*C.double)(d1), (*C.double)(d2), (*C.double)(b1), (C.double)(b2), (*C.double)(P))
}

func CBLAS_DROT(N int, X *float64, incX int, Y *float64, incY int, c float64, s float64) {
	C.cblas_drot((C.int)(N), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY), (C.double)(c), (C.double)(s))
}

func CBLAS_DROTM(N int, X *float64, incX int, Y *float64, incY int, P *float64) {
	C.cblas_drotm((C.int)(N), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY), (*C.double)(P))
}

func CBLAS_SSCAL(N int, alpha float32, X *float32, incX int) {
	C.cblas_sscal((C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX))
}

func CBLAS_DSCAL(N int, alpha float64, X *float64, incX int) {
	C.cblas_dscal((C.int)(N), (C.double)(alpha), (*C.double)(X), (C.int)(incX))
}

func CBLAS_CSCAL(N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int) {
	C.cblas_cscal((C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZSCAL(N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int) {
	C.cblas_zscal((C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_CSSCAL(N int, alpha float32, X unsafe.Pointer, incX int) {
	C.cblas_csscal((C.int)(N), (C.float)(alpha), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZDSCAL(N int, alpha float64, X unsafe.Pointer, incX int) {
	C.cblas_zdscal((C.int)(N), (C.double)(alpha), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_SGEMV(order uint32, TransA uint32, M int, N int, alpha float32, A *float32, lda int, X *float32, incX int, beta float32, Y *float32, incY int) {
	C.cblas_sgemv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX), (C.float)(beta), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_SGBMV(order uint32, TransA uint32, M int, N int, KL int, KU int, alpha float32, A *float32, lda int, X *float32, incX int, beta float32, Y *float32, incY int) {
	C.cblas_sgbmv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (C.int)(KL), (C.int)(KU), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX), (C.float)(beta), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_STRMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A *float32, lda int, X *float32, incX int) {
	C.cblas_strmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX))
}

func CBLAS_STBMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A *float32, lda int, X *float32, incX int) {
	C.cblas_stbmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX))
}

func CBLAS_STPMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap *float32, X *float32, incX int) {
	C.cblas_stpmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.float)(Ap), (*C.float)(X), (C.int)(incX))
}

func CBLAS_STRSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A *float32, lda int, X *float32, incX int) {
	C.cblas_strsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX))
}

func CBLAS_STBSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A *float32, lda int, X *float32, incX int) {
	C.cblas_stbsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX))
}

func CBLAS_STPSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap *float32, X *float32, incX int) {
	C.cblas_stpsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.float)(Ap), (*C.float)(X), (C.int)(incX))
}

func CBLAS_DGEMV(order uint32, TransA uint32, M int, N int, alpha float64, A *float64, lda int, X *float64, incX int, beta float64, Y *float64, incY int) {
	C.cblas_dgemv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX), (C.double)(beta), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_DGBMV(order uint32, TransA uint32, M int, N int, KL int, KU int, alpha float64, A *float64, lda int, X *float64, incX int, beta float64, Y *float64, incY int) {
	C.cblas_dgbmv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (C.int)(KL), (C.int)(KU), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX), (C.double)(beta), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_DTRMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A *float64, lda int, X *float64, incX int) {
	C.cblas_dtrmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX))
}

func CBLAS_DTBMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A *float64, lda int, X *float64, incX int) {
	C.cblas_dtbmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX))
}

func CBLAS_DTPMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap *float64, X *float64, incX int) {
	C.cblas_dtpmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.double)(Ap), (*C.double)(X), (C.int)(incX))
}

func CBLAS_DTRSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A *float64, lda int, X *float64, incX int) {
	C.cblas_dtrsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX))
}

func CBLAS_DTBSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A *float64, lda int, X *float64, incX int) {
	C.cblas_dtbsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX))
}

func CBLAS_DTPSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap *float64, X *float64, incX int) {
	C.cblas_dtpsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (*C.double)(Ap), (*C.double)(X), (C.int)(incX))
}

func CBLAS_CGEMV(order uint32, TransA uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_cgemv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_CGBMV(order uint32, TransA uint32, M int, N int, KL int, KU int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_cgbmv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (C.int)(KL), (C.int)(KU), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_CTRMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ctrmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_CTBMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ctbmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_CTPMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap unsafe.Pointer, X unsafe.Pointer, incX int) {
	C.cblas_ctpmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(Ap), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_CTRSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ctrsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_CTBSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ctbsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_CTPSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap unsafe.Pointer, X unsafe.Pointer, incX int) {
	C.cblas_ctpsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(Ap), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZGEMV(order uint32, TransA uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_zgemv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZGBMV(order uint32, TransA uint32, M int, N int, KL int, KU int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_zgbmv((uint32)(order), (uint32)(TransA), (C.int)(M), (C.int)(N), (C.int)(KL), (C.int)(KU), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZTRMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ztrmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZTBMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ztbmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZTPMV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap unsafe.Pointer, X unsafe.Pointer, incX int) {
	C.cblas_ztpmv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(Ap), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZTRSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ztrsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZTBSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, K int, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int) {
	C.cblas_ztbsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (C.int)(K), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_ZTPSV(order uint32, Uplo uint32, TransA uint32, Diag uint32, N int, Ap unsafe.Pointer, X unsafe.Pointer, incX int) {
	C.cblas_ztpsv((uint32)(order), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(N), (unsafe.Pointer)(Ap), (unsafe.Pointer)(X), (C.int)(incX))
}

func CBLAS_SSYMV(order uint32, Uplo uint32, N int, alpha float32, A *float32, lda int, X *float32, incX int, beta float32, Y *float32, incY int) {
	C.cblas_ssymv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX), (C.float)(beta), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_SSBMV(order uint32, Uplo uint32, N int, K int, alpha float32, A *float32, lda int, X *float32, incX int, beta float32, Y *float32, incY int) {
	C.cblas_ssbmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.int)(K), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(X), (C.int)(incX), (C.float)(beta), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_SSPMV(order uint32, Uplo uint32, N int, alpha float32, Ap *float32, X *float32, incX int, beta float32, Y *float32, incY int) {
	C.cblas_sspmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (*C.float)(Ap), (*C.float)(X), (C.int)(incX), (C.float)(beta), (*C.float)(Y), (C.int)(incY))
}

func CBLAS_SGER(order uint32, M int, N int, alpha float32, X *float32, incX int, Y *float32, incY int, A *float32, lda int) {
	C.cblas_sger((uint32)(order), (C.int)(M), (C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY), (*C.float)(A), (C.int)(lda))
}

func CBLAS_SSYR(order uint32, Uplo uint32, N int, alpha float32, X *float32, incX int, A *float32, lda int) {
	C.cblas_ssyr((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX), (*C.float)(A), (C.int)(lda))
}

func CBLAS_SSPR(order uint32, Uplo uint32, N int, alpha float32, X *float32, incX int, Ap *float32) {
	C.cblas_sspr((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX), (*C.float)(Ap))
}

func CBLAS_SSYR2(order uint32, Uplo uint32, N int, alpha float32, X *float32, incX int, Y *float32, incY int, A *float32, lda int) {
	C.cblas_ssyr2((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY), (*C.float)(A), (C.int)(lda))
}

func CBLAS_SSPR2(order uint32, Uplo uint32, N int, alpha float32, X *float32, incX int, Y *float32, incY int, A *float32) {
	C.cblas_sspr2((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (*C.float)(X), (C.int)(incX), (*C.float)(Y), (C.int)(incY), (*C.float)(A))
}

func CBLAS_DSYMV(order uint32, Uplo uint32, N int, alpha float64, A *float64, lda int, X *float64, incX int, beta float64, Y *float64, incY int) {
	C.cblas_dsymv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX), (C.double)(beta), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_DSBMV(order uint32, Uplo uint32, N int, K int, alpha float64, A *float64, lda int, X *float64, incX int, beta float64, Y *float64, incY int) {
	C.cblas_dsbmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.int)(K), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(X), (C.int)(incX), (C.double)(beta), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_DSPMV(order uint32, Uplo uint32, N int, alpha float64, Ap *float64, X *float64, incX int, beta float64, Y *float64, incY int) {
	C.cblas_dspmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (*C.double)(Ap), (*C.double)(X), (C.int)(incX), (C.double)(beta), (*C.double)(Y), (C.int)(incY))
}

func CBLAS_DGER(order uint32, M int, N int, alpha float64, X *float64, incX int, Y *float64, incY int, A *float64, lda int) {
	C.cblas_dger((uint32)(order), (C.int)(M), (C.int)(N), (C.double)(alpha), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY), (*C.double)(A), (C.int)(lda))
}

func CBLAS_DSYR(order uint32, Uplo uint32, N int, alpha float64, X *float64, incX int, A *float64, lda int) {
	C.cblas_dsyr((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (*C.double)(X), (C.int)(incX), (*C.double)(A), (C.int)(lda))
}

func CBLAS_DSPR(order uint32, Uplo uint32, N int, alpha float64, X *float64, incX int, Ap *float64) {
	C.cblas_dspr((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (*C.double)(X), (C.int)(incX), (*C.double)(Ap))
}

func CBLAS_DSYR2(order uint32, Uplo uint32, N int, alpha float64, X *float64, incX int, Y *float64, incY int, A *float64, lda int) {
	C.cblas_dsyr2((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY), (*C.double)(A), (C.int)(lda))
}

func CBLAS_DSPR2(order uint32, Uplo uint32, N int, alpha float64, X *float64, incX int, Y *float64, incY int, A *float64) {
	C.cblas_dspr2((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (*C.double)(X), (C.int)(incX), (*C.double)(Y), (C.int)(incY), (*C.double)(A))
}

func CBLAS_CHEMV(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_chemv((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_CHBMV(order uint32, Uplo uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_chbmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_CHPMV(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, Ap unsafe.Pointer, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_chpmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(Ap), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_CGERU(order uint32, M int, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, A unsafe.Pointer, lda int) {
	C.cblas_cgeru((uint32)(order), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_CGERC(order uint32, M int, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, A unsafe.Pointer, lda int) {
	C.cblas_cgerc((uint32)(order), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_CHER(order uint32, Uplo uint32, N int, alpha float32, X unsafe.Pointer, incX int, A unsafe.Pointer, lda int) {
	C.cblas_cher((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_CHPR(order uint32, Uplo uint32, N int, alpha float32, X unsafe.Pointer, incX int, A unsafe.Pointer) {
	C.cblas_chpr((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.float)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(A))
}

func CBLAS_CHER2(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, A unsafe.Pointer, lda int) {
	C.cblas_cher2((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_CHPR2(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, Ap unsafe.Pointer) {
	C.cblas_chpr2((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(Ap))
}

func CBLAS_ZHEMV(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_zhemv((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZHBMV(order uint32, Uplo uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_zhbmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZHPMV(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, Ap unsafe.Pointer, X unsafe.Pointer, incX int, beta unsafe.Pointer, Y unsafe.Pointer, incY int) {
	C.cblas_zhpmv((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(Ap), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(beta), (unsafe.Pointer)(Y), (C.int)(incY))
}

func CBLAS_ZGERU(order uint32, M int, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, A unsafe.Pointer, lda int) {
	C.cblas_zgeru((uint32)(order), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_ZGERC(order uint32, M int, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, A unsafe.Pointer, lda int) {
	C.cblas_zgerc((uint32)(order), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_ZHER(order uint32, Uplo uint32, N int, alpha float64, X unsafe.Pointer, incX int, A unsafe.Pointer, lda int) {
	C.cblas_zher((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_ZHPR(order uint32, Uplo uint32, N int, alpha float64, X unsafe.Pointer, incX int, A unsafe.Pointer) {
	C.cblas_zhpr((uint32)(order), (uint32)(Uplo), (C.int)(N), (C.double)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(A))
}

func CBLAS_ZHER2(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, A unsafe.Pointer, lda int) {
	C.cblas_zher2((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(A), (C.int)(lda))
}

func CBLAS_ZHPR2(order uint32, Uplo uint32, N int, alpha unsafe.Pointer, X unsafe.Pointer, incX int, Y unsafe.Pointer, incY int, Ap unsafe.Pointer) {
	C.cblas_zhpr2((uint32)(order), (uint32)(Uplo), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(X), (C.int)(incX), (unsafe.Pointer)(Y), (C.int)(incY), (unsafe.Pointer)(Ap))
}

func CBLAS_SGEMM(Order uint32, TransA uint32, TransB uint32, M int, N int, K int, alpha float32, A *float32, lda int, B *float32, ldb int, beta float32, C *float32, ldc int) {
	C.cblas_sgemm((uint32)(Order), (uint32)(TransA), (uint32)(TransB), (C.int)(M), (C.int)(N), (C.int)(K), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(B), (C.int)(ldb), (C.float)(beta), (*C.float)(C), (C.int)(ldc))
}

func CBLAS_SSYMM(Order uint32, Side uint32, Uplo uint32, M int, N int, alpha float32, A *float32, lda int, B *float32, ldb int, beta float32, C *float32, ldc int) {
	C.cblas_ssymm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (C.int)(M), (C.int)(N), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(B), (C.int)(ldb), (C.float)(beta), (*C.float)(C), (C.int)(ldc))
}

func CBLAS_SSYRK(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha float32, A *float32, lda int, beta float32, C *float32, ldc int) {
	C.cblas_ssyrk((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (C.float)(beta), (*C.float)(C), (C.int)(ldc))
}

func CBLAS_SSYR2K(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha float32, A *float32, lda int, B *float32, ldb int, beta float32, C *float32, ldc int) {
	C.cblas_ssyr2k((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(B), (C.int)(ldb), (C.float)(beta), (*C.float)(C), (C.int)(ldc))
}

func CBLAS_STRMM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha float32, A *float32, lda int, B *float32, ldb int) {
	C.cblas_strmm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(B), (C.int)(ldb))
}

func CBLAS_STRSM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha float32, A *float32, lda int, B *float32, ldb int) {
	C.cblas_strsm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (C.float)(alpha), (*C.float)(A), (C.int)(lda), (*C.float)(B), (C.int)(ldb))
}

func CBLAS_DGEMM(Order uint32, TransA uint32, TransB uint32, M int, N int, K int, alpha float64, A *float64, lda int, B *float64, ldb int, beta float64, C *float64, ldc int) {
	C.cblas_dgemm((uint32)(Order), (uint32)(TransA), (uint32)(TransB), (C.int)(M), (C.int)(N), (C.int)(K), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(B), (C.int)(ldb), (C.double)(beta), (*C.double)(C), (C.int)(ldc))
}

func CBLAS_DSYMM(Order uint32, Side uint32, Uplo uint32, M int, N int, alpha float64, A *float64, lda int, B *float64, ldb int, beta float64, C *float64, ldc int) {
	C.cblas_dsymm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (C.int)(M), (C.int)(N), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(B), (C.int)(ldb), (C.double)(beta), (*C.double)(C), (C.int)(ldc))
}

func CBLAS_DSYRK(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha float64, A *float64, lda int, beta float64, C *float64, ldc int) {
	C.cblas_dsyrk((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (C.double)(beta), (*C.double)(C), (C.int)(ldc))
}

func CBLAS_DSYR2K(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha float64, A *float64, lda int, B *float64, ldb int, beta float64, C *float64, ldc int) {
	C.cblas_dsyr2k((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(B), (C.int)(ldb), (C.double)(beta), (*C.double)(C), (C.int)(ldc))
}

func CBLAS_DTRMM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha float64, A *float64, lda int, B *float64, ldb int) {
	C.cblas_dtrmm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(B), (C.int)(ldb))
}

func CBLAS_DTRSM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha float64, A *float64, lda int, B *float64, ldb int) {
	C.cblas_dtrsm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (C.double)(alpha), (*C.double)(A), (C.int)(lda), (*C.double)(B), (C.int)(ldb))
}

func CBLAS_CGEMM(Order uint32, TransA uint32, TransB uint32, M int, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_cgemm((uint32)(Order), (uint32)(TransA), (uint32)(TransB), (C.int)(M), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_CSYMM(Order uint32, Side uint32, Uplo uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_csymm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_CSYRK(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_csyrk((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_CSYR2K(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_csyr2k((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_CTRMM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int) {
	C.cblas_ctrmm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb))
}

func CBLAS_CTRSM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int) {
	C.cblas_ctrsm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb))
}

func CBLAS_ZGEMM(Order uint32, TransA uint32, TransB uint32, M int, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_zgemm((uint32)(Order), (uint32)(TransA), (uint32)(TransB), (C.int)(M), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_ZSYMM(Order uint32, Side uint32, Uplo uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_zsymm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_ZSYRK(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_zsyrk((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_ZSYR2K(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_zsyr2k((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_ZTRMM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int) {
	C.cblas_ztrmm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb))
}

func CBLAS_ZTRSM(Order uint32, Side uint32, Uplo uint32, TransA uint32, Diag uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int) {
	C.cblas_ztrsm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (uint32)(TransA), (uint32)(Diag), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb))
}

func CBLAS_CHEMM(Order uint32, Side uint32, Uplo uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_chemm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_CHERK(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha float32, A unsafe.Pointer, lda int, beta float32, C unsafe.Pointer, ldc int) {
	C.cblas_cherk((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (C.float)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (C.float)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_CHER2K(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta float32, C unsafe.Pointer, ldc int) {
	C.cblas_cher2k((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (C.float)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_ZHEMM(Order uint32, Side uint32, Uplo uint32, M int, N int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta unsafe.Pointer, C unsafe.Pointer, ldc int) {
	C.cblas_zhemm((uint32)(Order), (uint32)(Side), (uint32)(Uplo), (C.int)(M), (C.int)(N), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (unsafe.Pointer)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_ZHERK(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha float64, A unsafe.Pointer, lda int, beta float64, C unsafe.Pointer, ldc int) {
	C.cblas_zherk((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (C.double)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (C.double)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}

func CBLAS_ZHER2K(Order uint32, Uplo uint32, Trans uint32, N int, K int, alpha unsafe.Pointer, A unsafe.Pointer, lda int, B unsafe.Pointer, ldb int, beta float64, C unsafe.Pointer, ldc int) {
	C.cblas_zher2k((uint32)(Order), (uint32)(Uplo), (uint32)(Trans), (C.int)(N), (C.int)(K), (unsafe.Pointer)(alpha), (unsafe.Pointer)(A), (C.int)(lda), (unsafe.Pointer)(B), (C.int)(ldb), (C.double)(beta), (unsafe.Pointer)(C), (C.int)(ldc))
}
