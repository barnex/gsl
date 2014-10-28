#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l3.h"

void
cblas_sgemm (const enum CBLAS_ORDER Order, const enum CBLAS_TRANSPOSE TransA,
             const enum CBLAS_TRANSPOSE TransB, const int M, const int N,
             const int K, const float alpha, const float *A, const int lda,
             const float *B, const int ldb, const float beta, float *C,
             const int ldc)
{
#define BASE float
#include "cblas_source_gemm_r.h"
#undef BASE
}
