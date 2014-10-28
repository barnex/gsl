#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l3.h"

void
cblas_dsyrk (const enum CBLAS_ORDER Order, const enum CBLAS_UPLO Uplo,
             const enum CBLAS_TRANSPOSE Trans, const int N, const int K,
             const double alpha, const double *A, const int lda,
             const double beta, double *C, const int ldc)
{
#define BASE double
#include "cblas_source_syrk_r.h"
#undef BASE
}
