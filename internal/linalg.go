package internal

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

//#include "gsl_gsl_linalg.h"
import "C"

import (
	"unsafe"
)

func Matmult(A [][]float64, B [][]float64, C [][]float64) int {
	return (int)(C.gsl_linalg_matmult((*C.gsl_matrix)(A), (*C.gsl_matrix)(B), (*C.gsl_matrix)(C)))
}

func MatmultMod(A [][]float64, modA Transpose, B [][]float64, modB Transpose, C [][]float64) int {
	return (int)(C.gsl_linalg_matmult_mod((*C.gsl_matrix)(A), (C.gsl_linalg_matrix_mod_t)(modA), (*C.gsl_matrix)(B), (C.gsl_linalg_matrix_mod_t)(modB), (*C.gsl_matrix)(C)))
}

func ExponentialSs(A [][]float64, eA [][]float64, mode Mode) int {
	return (int)(C.gsl_linalg_exponential_ss((*C.gsl_matrix)(A), (*C.gsl_matrix)(eA), (C.gsl_mode_t)(mode)))
}

func HouseholderTransform(v []float64) float64 {
	return (float64)(C.gsl_linalg_householder_transform((*C.gsl_vector)(v)))
}

func ComplexHouseholderTransform(v []complex128) complex128 {
	return (complex128)(C.gsl_linalg_complex_householder_transform((*C.gsl_vector_complex)(v)))
}

func HouseholderHm(tau float64, v []float64, A [][]float64) int {
	return (int)(C.gsl_linalg_householder_hm((C.double)(tau), (*C.gsl_vector)(v), (*C.gsl_matrix)(A)))
}

func HouseholderMh(tau float64, v []float64, A [][]float64) int {
	return (int)(C.gsl_linalg_householder_mh((C.double)(tau), (*C.gsl_vector)(v), (*C.gsl_matrix)(A)))
}

func HouseholderHv(tau float64, v []float64, w []float64) int {
	return (int)(C.gsl_linalg_householder_hv((C.double)(tau), (*C.gsl_vector)(v), (*C.gsl_vector)(w)))
}

func HouseholderHm1(tau float64, A [][]float64) int {
	return (int)(C.gsl_linalg_householder_hm1((C.double)(tau), (*C.gsl_matrix)(A)))
}

func ComplexHouseholderHm(tau complex128, v []complex128, A [][]complex128) int {
	return (int)(C.gsl_linalg_complex_householder_hm((C.gsl_complex)(tau), (*C.gsl_vector_complex)(v), (*C.gsl_matrix_complex)(A)))
}

func ComplexHouseholderMh(tau complex128, v []complex128, A [][]complex128) int {
	return (int)(C.gsl_linalg_complex_householder_mh((C.gsl_complex)(tau), (*C.gsl_vector_complex)(v), (*C.gsl_matrix_complex)(A)))
}

func ComplexHouseholderHv(tau complex128, v []complex128, w []complex128) int {
	return (int)(C.gsl_linalg_complex_householder_hv((C.gsl_complex)(tau), (*C.gsl_vector_complex)(v), (*C.gsl_vector_complex)(w)))
}

func HessenbergDecomp(A [][]float64, tau []float64) int {
	return (int)(C.gsl_linalg_hessenberg_decomp((*C.gsl_matrix)(A), (*C.gsl_vector)(tau)))
}

func HessenbergUnpack(H [][]float64, tau []float64, U [][]float64) int {
	return (int)(C.gsl_linalg_hessenberg_unpack((*C.gsl_matrix)(H), (*C.gsl_vector)(tau), (*C.gsl_matrix)(U)))
}

func HessenbergUnpackAccum(H [][]float64, tau []float64, U [][]float64) int {
	return (int)(C.gsl_linalg_hessenberg_unpack_accum((*C.gsl_matrix)(H), (*C.gsl_vector)(tau), (*C.gsl_matrix)(U)))
}

func HessenbergSetZero(H [][]float64) int {
	return (int)(C.gsl_linalg_hessenberg_set_zero((*C.gsl_matrix)(H)))
}

func HessenbergSubmatrix(M [][]float64, A [][]float64, top int, tau []float64) int {
	return (int)(C.gsl_linalg_hessenberg_submatrix((*C.gsl_matrix)(M), (*C.gsl_matrix)(A), (C.size_t)(top), (*C.gsl_vector)(tau)))
}

func Hessenberg(A [][]float64, tau []float64) int {
	return (int)(C.gsl_linalg_hessenberg((*C.gsl_matrix)(A), (*C.gsl_vector)(tau)))
}

func HesstriDecomp(A [][]float64, B [][]float64, U [][]float64, V [][]float64, work []float64) int {
	return (int)(C.gsl_linalg_hesstri_decomp((*C.gsl_matrix)(A), (*C.gsl_matrix)(B), (*C.gsl_matrix)(U), (*C.gsl_matrix)(V), (*C.gsl_vector)(work)))
}

func SVDecomp(A [][]float64, V [][]float64, S []float64, work []float64) int {
	return (int)(C.gsl_linalg_SV_decomp((*C.gsl_matrix)(A), (*C.gsl_matrix)(V), (*C.gsl_vector)(S), (*C.gsl_vector)(work)))
}

func SVDecompMod(A [][]float64, X [][]float64, V [][]float64, S []float64, work []float64) int {
	return (int)(C.gsl_linalg_SV_decomp_mod((*C.gsl_matrix)(A), (*C.gsl_matrix)(X), (*C.gsl_matrix)(V), (*C.gsl_vector)(S), (*C.gsl_vector)(work)))
}

func SVDecompJacobi(A [][]float64, Q [][]float64, S []float64) int {
	return (int)(C.gsl_linalg_SV_decomp_jacobi((*C.gsl_matrix)(A), (*C.gsl_matrix)(Q), (*C.gsl_vector)(S)))
}

func SVSolve(U [][]float64, Q [][]float64, S []float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_SV_solve((*C.gsl_matrix)(U), (*C.gsl_matrix)(Q), (*C.gsl_vector)(S), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func SVLeverage(U [][]float64, h []float64) int {
	return (int)(C.gsl_linalg_SV_leverage((*C.gsl_matrix)(U), (*C.gsl_vector)(h)))
}

func LUDecomp(A [][]float64, p *Permutation, signum *int) int {
	return (int)(C.gsl_linalg_LU_decomp((*C.gsl_matrix)(A), (*C.gsl_permutation)(p), (*C.int)(signum)))
}

func LUSolve(LU [][]float64, p *Permutation, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_LU_solve((*C.gsl_matrix)(LU), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func LUSvx(LU [][]float64, p *Permutation, x []float64) int {
	return (int)(C.gsl_linalg_LU_svx((*C.gsl_matrix)(LU), (*C.gsl_permutation)(p), (*C.gsl_vector)(x)))
}

func LURefine(A [][]float64, LU [][]float64, p *Permutation, b []float64, x []float64, residual []float64) int {
	return (int)(C.gsl_linalg_LU_refine((*C.gsl_matrix)(A), (*C.gsl_matrix)(LU), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x), (*C.gsl_vector)(residual)))
}

func LUInvert(LU [][]float64, p *Permutation, inverse [][]float64) int {
	return (int)(C.gsl_linalg_LU_invert((*C.gsl_matrix)(LU), (*C.gsl_permutation)(p), (*C.gsl_matrix)(inverse)))
}

func LUDet(LU [][]float64, signum int) float64 {
	return (float64)(C.gsl_linalg_LU_det((*C.gsl_matrix)(LU), (C.int)(signum)))
}

func LULndet(LU [][]float64) float64 {
	return (float64)(C.gsl_linalg_LU_lndet((*C.gsl_matrix)(LU)))
}

func LUSgndet(lu [][]float64, signum int) int {
	return (int)(C.gsl_linalg_LU_sgndet((*C.gsl_matrix)(lu), (C.int)(signum)))
}

func ComplexLUDecomp(A [][]complex128, p *Permutation, signum *int) int {
	return (int)(C.gsl_linalg_complex_LU_decomp((*C.gsl_matrix_complex)(A), (*C.gsl_permutation)(p), (*C.int)(signum)))
}

func ComplexLUSolve(LU [][]complex128, p *Permutation, b []complex128, x []complex128) int {
	return (int)(C.gsl_linalg_complex_LU_solve((*C.gsl_matrix_complex)(LU), (*C.gsl_permutation)(p), (*C.gsl_vector_complex)(b), (*C.gsl_vector_complex)(x)))
}

func ComplexLUSvx(LU [][]complex128, p *Permutation, x []complex128) int {
	return (int)(C.gsl_linalg_complex_LU_svx((*C.gsl_matrix_complex)(LU), (*C.gsl_permutation)(p), (*C.gsl_vector_complex)(x)))
}

func ComplexLURefine(A [][]complex128, LU [][]complex128, p *Permutation, b []complex128, x []complex128, residual []complex128) int {
	return (int)(C.gsl_linalg_complex_LU_refine((*C.gsl_matrix_complex)(A), (*C.gsl_matrix_complex)(LU), (*C.gsl_permutation)(p), (*C.gsl_vector_complex)(b), (*C.gsl_vector_complex)(x), (*C.gsl_vector_complex)(residual)))
}

func ComplexLUInvert(LU [][]complex128, p *Permutation, inverse [][]complex128) int {
	return (int)(C.gsl_linalg_complex_LU_invert((*C.gsl_matrix_complex)(LU), (*C.gsl_permutation)(p), (*C.gsl_matrix_complex)(inverse)))
}

func ComplexLUDet(LU [][]complex128, signum int) complex128 {
	return (complex128)(C.gsl_linalg_complex_LU_det((*C.gsl_matrix_complex)(LU), (C.int)(signum)))
}

func ComplexLULndet(LU [][]complex128) float64 {
	return (float64)(C.gsl_linalg_complex_LU_lndet((*C.gsl_matrix_complex)(LU)))
}

func ComplexLUSgndet(LU [][]complex128, signum int) complex128 {
	return (complex128)(C.gsl_linalg_complex_LU_sgndet((*C.gsl_matrix_complex)(LU), (C.int)(signum)))
}

func QRDecomp(A [][]float64, tau []float64) int {
	return (int)(C.gsl_linalg_QR_decomp((*C.gsl_matrix)(A), (*C.gsl_vector)(tau)))
}

func QRSolve(QR [][]float64, tau []float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_QR_solve((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func QRSvx(QR [][]float64, tau []float64, x []float64) int {
	return (int)(C.gsl_linalg_QR_svx((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_vector)(x)))
}

func QRLssolve(QR [][]float64, tau []float64, b []float64, x []float64, residual []float64) int {
	return (int)(C.gsl_linalg_QR_lssolve((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_vector)(b), (*C.gsl_vector)(x), (*C.gsl_vector)(residual)))
}

func QRQRsolve(Q [][]float64, R [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_QR_QRsolve((*C.gsl_matrix)(Q), (*C.gsl_matrix)(R), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func QRRsolve(QR [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_QR_Rsolve((*C.gsl_matrix)(QR), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func QRRsvx(QR [][]float64, x []float64) int {
	return (int)(C.gsl_linalg_QR_Rsvx((*C.gsl_matrix)(QR), (*C.gsl_vector)(x)))
}

func QRUpdate(Q [][]float64, R [][]float64, w []float64, v []float64) int {
	return (int)(C.gsl_linalg_QR_update((*C.gsl_matrix)(Q), (*C.gsl_matrix)(R), (*C.gsl_vector)(w), (*C.gsl_vector)(v)))
}

func QRQTvec(QR [][]float64, tau []float64, v []float64) int {
	return (int)(C.gsl_linalg_QR_QTvec((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_vector)(v)))
}

func QRQvec(QR [][]float64, tau []float64, v []float64) int {
	return (int)(C.gsl_linalg_QR_Qvec((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_vector)(v)))
}

func QRQTmat(QR [][]float64, tau []float64, A [][]float64) int {
	return (int)(C.gsl_linalg_QR_QTmat((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_matrix)(A)))
}

func QRUnpack(QR [][]float64, tau []float64, Q [][]float64, R [][]float64) int {
	return (int)(C.gsl_linalg_QR_unpack((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_matrix)(Q), (*C.gsl_matrix)(R)))
}

func RSolve(R [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_R_solve((*C.gsl_matrix)(R), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func RSvx(R [][]float64, x []float64) int {
	return (int)(C.gsl_linalg_R_svx((*C.gsl_matrix)(R), (*C.gsl_vector)(x)))
}

func QRPTDecomp(A [][]float64, tau []float64, p *Permutation, signum *int, norm []float64) int {
	return (int)(C.gsl_linalg_QRPT_decomp((*C.gsl_matrix)(A), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.int)(signum), (*C.gsl_vector)(norm)))
}

func QRPTDecomp2(A [][]float64, q [][]float64, r [][]float64, tau []float64, p *Permutation, signum *int, norm []float64) int {
	return (int)(C.gsl_linalg_QRPT_decomp2((*C.gsl_matrix)(A), (*C.gsl_matrix)(q), (*C.gsl_matrix)(r), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.int)(signum), (*C.gsl_vector)(norm)))
}

func QRPTSolve(QR [][]float64, tau []float64, p *Permutation, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_QRPT_solve((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func QRPTSvx(QR [][]float64, tau []float64, p *Permutation, x []float64) int {
	return (int)(C.gsl_linalg_QRPT_svx((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.gsl_vector)(x)))
}

func QRPTQRsolve(Q [][]float64, R [][]float64, p *Permutation, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_QRPT_QRsolve((*C.gsl_matrix)(Q), (*C.gsl_matrix)(R), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func QRPTRsolve(QR [][]float64, p *Permutation, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_QRPT_Rsolve((*C.gsl_matrix)(QR), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func QRPTRsvx(QR [][]float64, p *Permutation, x []float64) int {
	return (int)(C.gsl_linalg_QRPT_Rsvx((*C.gsl_matrix)(QR), (*C.gsl_permutation)(p), (*C.gsl_vector)(x)))
}

func QRPTUpdate(Q [][]float64, R [][]float64, p *Permutation, u []float64, v []float64) int {
	return (int)(C.gsl_linalg_QRPT_update((*C.gsl_matrix)(Q), (*C.gsl_matrix)(R), (*C.gsl_permutation)(p), (*C.gsl_vector)(u), (*C.gsl_vector)(v)))
}

func LQDecomp(A [][]float64, tau []float64) int {
	return (int)(C.gsl_linalg_LQ_decomp((*C.gsl_matrix)(A), (*C.gsl_vector)(tau)))
}

func LQSolveT(LQ [][]float64, tau []float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_LQ_solve_T((*C.gsl_matrix)(LQ), (*C.gsl_vector)(tau), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func LQSvxT(LQ [][]float64, tau []float64, x []float64) int {
	return (int)(C.gsl_linalg_LQ_svx_T((*C.gsl_matrix)(LQ), (*C.gsl_vector)(tau), (*C.gsl_vector)(x)))
}

func LQLssolveT(LQ [][]float64, tau []float64, b []float64, x []float64, residual []float64) int {
	return (int)(C.gsl_linalg_LQ_lssolve_T((*C.gsl_matrix)(LQ), (*C.gsl_vector)(tau), (*C.gsl_vector)(b), (*C.gsl_vector)(x), (*C.gsl_vector)(residual)))
}

func LQLsolveT(LQ [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_LQ_Lsolve_T((*C.gsl_matrix)(LQ), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func LQLsvxT(LQ [][]float64, x []float64) int {
	return (int)(C.gsl_linalg_LQ_Lsvx_T((*C.gsl_matrix)(LQ), (*C.gsl_vector)(x)))
}

func LSolveT(L [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_L_solve_T((*C.gsl_matrix)(L), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func LQVecQ(LQ [][]float64, tau []float64, v []float64) int {
	return (int)(C.gsl_linalg_LQ_vecQ((*C.gsl_matrix)(LQ), (*C.gsl_vector)(tau), (*C.gsl_vector)(v)))
}

func LQVecQT(LQ [][]float64, tau []float64, v []float64) int {
	return (int)(C.gsl_linalg_LQ_vecQT((*C.gsl_matrix)(LQ), (*C.gsl_vector)(tau), (*C.gsl_vector)(v)))
}

func LQUnpack(LQ [][]float64, tau []float64, Q [][]float64, L [][]float64) int {
	return (int)(C.gsl_linalg_LQ_unpack((*C.gsl_matrix)(LQ), (*C.gsl_vector)(tau), (*C.gsl_matrix)(Q), (*C.gsl_matrix)(L)))
}

func LQUpdate(Q [][]float64, R [][]float64, v []float64, w []float64) int {
	return (int)(C.gsl_linalg_LQ_update((*C.gsl_matrix)(Q), (*C.gsl_matrix)(R), (*C.gsl_vector)(v), (*C.gsl_vector)(w)))
}

func LQLQsolve(Q [][]float64, L [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_LQ_LQsolve((*C.gsl_matrix)(Q), (*C.gsl_matrix)(L), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func PTLQDecomp(A [][]float64, tau []float64, p *Permutation, signum *int, norm []float64) int {
	return (int)(C.gsl_linalg_PTLQ_decomp((*C.gsl_matrix)(A), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.int)(signum), (*C.gsl_vector)(norm)))
}

func PTLQDecomp2(A [][]float64, q [][]float64, r [][]float64, tau []float64, p *Permutation, signum *int, norm []float64) int {
	return (int)(C.gsl_linalg_PTLQ_decomp2((*C.gsl_matrix)(A), (*C.gsl_matrix)(q), (*C.gsl_matrix)(r), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.int)(signum), (*C.gsl_vector)(norm)))
}

func PTLQSolveT(QR [][]float64, tau []float64, p *Permutation, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_PTLQ_solve_T((*C.gsl_matrix)(QR), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func PTLQSvxT(LQ [][]float64, tau []float64, p *Permutation, x []float64) int {
	return (int)(C.gsl_linalg_PTLQ_svx_T((*C.gsl_matrix)(LQ), (*C.gsl_vector)(tau), (*C.gsl_permutation)(p), (*C.gsl_vector)(x)))
}

func PTLQLQsolveT(Q [][]float64, L [][]float64, p *Permutation, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_PTLQ_LQsolve_T((*C.gsl_matrix)(Q), (*C.gsl_matrix)(L), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func PTLQLsolveT(LQ [][]float64, p *Permutation, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_PTLQ_Lsolve_T((*C.gsl_matrix)(LQ), (*C.gsl_permutation)(p), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func PTLQLsvxT(LQ [][]float64, p *Permutation, x []float64) int {
	return (int)(C.gsl_linalg_PTLQ_Lsvx_T((*C.gsl_matrix)(LQ), (*C.gsl_permutation)(p), (*C.gsl_vector)(x)))
}

func PTLQUpdate(Q [][]float64, L [][]float64, p *Permutation, v []float64, w []float64) int {
	return (int)(C.gsl_linalg_PTLQ_update((*C.gsl_matrix)(Q), (*C.gsl_matrix)(L), (*C.gsl_permutation)(p), (*C.gsl_vector)(v), (*C.gsl_vector)(w)))
}

func CholeskyDecomp(A [][]float64) int {
	return (int)(C.gsl_linalg_cholesky_decomp((*C.gsl_matrix)(A)))
}

func CholeskySolve(cholesky [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_cholesky_solve((*C.gsl_matrix)(cholesky), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func CholeskySvx(cholesky [][]float64, x []float64) int {
	return (int)(C.gsl_linalg_cholesky_svx((*C.gsl_matrix)(cholesky), (*C.gsl_vector)(x)))
}

func CholeskyInvert(cholesky [][]float64) int {
	return (int)(C.gsl_linalg_cholesky_invert((*C.gsl_matrix)(cholesky)))
}

func CholeskyDecompUnit(A [][]float64, D []float64) int {
	return (int)(C.gsl_linalg_cholesky_decomp_unit((*C.gsl_matrix)(A), (*C.gsl_vector)(D)))
}

func ComplexCholeskyDecomp(A [][]complex128) int {
	return (int)(C.gsl_linalg_complex_cholesky_decomp((*C.gsl_matrix_complex)(A)))
}

func ComplexCholeskySolve(cholesky [][]complex128, b []complex128, x []complex128) int {
	return (int)(C.gsl_linalg_complex_cholesky_solve((*C.gsl_matrix_complex)(cholesky), (*C.gsl_vector_complex)(b), (*C.gsl_vector_complex)(x)))
}

func ComplexCholeskySvx(cholesky [][]complex128, x []complex128) int {
	return (int)(C.gsl_linalg_complex_cholesky_svx((*C.gsl_matrix_complex)(cholesky), (*C.gsl_vector_complex)(x)))
}

func ComplexCholeskyInvert(cholesky [][]complex128) int {
	return (int)(C.gsl_linalg_complex_cholesky_invert((*C.gsl_matrix_complex)(cholesky)))
}

func SymmtdDecomp(A [][]float64, tau []float64) int {
	return (int)(C.gsl_linalg_symmtd_decomp((*C.gsl_matrix)(A), (*C.gsl_vector)(tau)))
}

func SymmtdUnpack(A [][]float64, tau []float64, Q [][]float64, diag []float64, subdiag []float64) int {
	return (int)(C.gsl_linalg_symmtd_unpack((*C.gsl_matrix)(A), (*C.gsl_vector)(tau), (*C.gsl_matrix)(Q), (*C.gsl_vector)(diag), (*C.gsl_vector)(subdiag)))
}

func SymmtdUnpackT(A [][]float64, diag []float64, subdiag []float64) int {
	return (int)(C.gsl_linalg_symmtd_unpack_T((*C.gsl_matrix)(A), (*C.gsl_vector)(diag), (*C.gsl_vector)(subdiag)))
}

func HermtdDecomp(A [][]complex128, tau []complex128) int {
	return (int)(C.gsl_linalg_hermtd_decomp((*C.gsl_matrix_complex)(A), (*C.gsl_vector_complex)(tau)))
}

func HermtdUnpack(A [][]complex128, tau []complex128, U [][]complex128, diag []float64, sudiag []float64) int {
	return (int)(C.gsl_linalg_hermtd_unpack((*C.gsl_matrix_complex)(A), (*C.gsl_vector_complex)(tau), (*C.gsl_matrix_complex)(U), (*C.gsl_vector)(diag), (*C.gsl_vector)(sudiag)))
}

func HermtdUnpackT(A [][]complex128, diag []float64, subdiag []float64) int {
	return (int)(C.gsl_linalg_hermtd_unpack_T((*C.gsl_matrix_complex)(A), (*C.gsl_vector)(diag), (*C.gsl_vector)(subdiag)))
}

func HHSolve(A [][]float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_HH_solve((*C.gsl_matrix)(A), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func HHSvx(A [][]float64, x []float64) int {
	return (int)(C.gsl_linalg_HH_svx((*C.gsl_matrix)(A), (*C.gsl_vector)(x)))
}

func SolveSymmTridiag(diag []float64, offdiag []float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_solve_symm_tridiag((*C.gsl_vector)(diag), (*C.gsl_vector)(offdiag), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func SolveTridiag(diag []float64, abovediag []float64, belowdiag []float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_solve_tridiag((*C.gsl_vector)(diag), (*C.gsl_vector)(abovediag), (*C.gsl_vector)(belowdiag), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func SolveSymmCycTridiag(diag []float64, offdiag []float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_solve_symm_cyc_tridiag((*C.gsl_vector)(diag), (*C.gsl_vector)(offdiag), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func SolveCycTridiag(diag []float64, abovediag []float64, belowdiag []float64, b []float64, x []float64) int {
	return (int)(C.gsl_linalg_solve_cyc_tridiag((*C.gsl_vector)(diag), (*C.gsl_vector)(abovediag), (*C.gsl_vector)(belowdiag), (*C.gsl_vector)(b), (*C.gsl_vector)(x)))
}

func BidiagDecomp(A [][]float64, tau_U []float64, tau_V []float64) int {
	return (int)(C.gsl_linalg_bidiag_decomp((*C.gsl_matrix)(A), (*C.gsl_vector)(tau_U), (*C.gsl_vector)(tau_V)))
}

func BidiagUnpack(A [][]float64, tau_U []float64, U [][]float64, tau_V []float64, V [][]float64, diag []float64, superdiag []float64) int {
	return (int)(C.gsl_linalg_bidiag_unpack((*C.gsl_matrix)(A), (*C.gsl_vector)(tau_U), (*C.gsl_matrix)(U), (*C.gsl_vector)(tau_V), (*C.gsl_matrix)(V), (*C.gsl_vector)(diag), (*C.gsl_vector)(superdiag)))
}

func BidiagUnpack2(A [][]float64, tau_U []float64, tau_V []float64, V [][]float64) int {
	return (int)(C.gsl_linalg_bidiag_unpack2((*C.gsl_matrix)(A), (*C.gsl_vector)(tau_U), (*C.gsl_vector)(tau_V), (*C.gsl_matrix)(V)))
}

func BidiagUnpackB(A [][]float64, diag []float64, superdiag []float64) int {
	return (int)(C.gsl_linalg_bidiag_unpack_B((*C.gsl_matrix)(A), (*C.gsl_vector)(diag), (*C.gsl_vector)(superdiag)))
}

func BalanceMatrix(A [][]float64, D []float64) int {
	return (int)(C.gsl_linalg_balance_matrix((*C.gsl_matrix)(A), (*C.gsl_vector)(D)))
}

func BalanceAccum(A [][]float64, D []float64) int {
	return (int)(C.gsl_linalg_balance_accum((*C.gsl_matrix)(A), (*C.gsl_vector)(D)))
}

func BalanceColumns(A [][]float64, D []float64) int {
	return (int)(C.gsl_linalg_balance_columns((*C.gsl_matrix)(A), (*C.gsl_vector)(D)))
}
