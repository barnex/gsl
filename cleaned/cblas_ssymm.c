#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l3.h"

void
cblas_ssymm (const enum CBLAS_ORDER Order, const enum CBLAS_SIDE Side,
             const enum CBLAS_UPLO Uplo, const int M, const int N,
             const float alpha, const float *A, const int lda, const float *B,
             const int ldb, const float beta, float *C, const int ldc)
{
#define BASE float
#include "cblas_source_symm_r.h"
#undef BASE
}
