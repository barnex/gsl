#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

#include "cblas_hypot.c.inc"

void
cblas_ctpsv (const enum CBLAS_ORDER order, const enum CBLAS_UPLO Uplo,
             const enum CBLAS_TRANSPOSE TransA, const enum CBLAS_DIAG Diag,
             const int N, const void *Ap, void *X, const int incX)
{
#define BASE float
#include "cblas_source_tpsv_c.h"
#undef BASE
}
