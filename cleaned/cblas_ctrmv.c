#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

void
cblas_ctrmv (const enum CBLAS_ORDER order, const enum CBLAS_UPLO Uplo,
             const enum CBLAS_TRANSPOSE TransA, const enum CBLAS_DIAG Diag,
             const int N, const void *A, const int lda, void *X,
             const int incX)
{
#define BASE float
#include "cblas_source_trmv_c.h"
#undef BASE
}
