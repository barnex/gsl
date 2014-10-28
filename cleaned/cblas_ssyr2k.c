#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l3.h"

void
cblas_ssyr2k (const enum CBLAS_ORDER Order, const enum CBLAS_UPLO Uplo,
              const enum CBLAS_TRANSPOSE Trans, const int N, const int K,
              const float alpha, const float *A, const int lda,
              const float *B, const int ldb, const float beta, float *C,
              const int ldc)
{
#define BASE float
#include "cblas_source_syr2k_r.h"
#undef BASE
}
