#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l3.h"

void
cblas_zsyrk (const enum CBLAS_ORDER Order, const enum CBLAS_UPLO Uplo,
             const enum CBLAS_TRANSPOSE Trans, const int N, const int K,
             const void *alpha, const void *A, const int lda,
             const void *beta, void *C, const int ldc)
{
#define BASE double
#include "cblas_source_syrk_c.h"
#undef BASE
}
