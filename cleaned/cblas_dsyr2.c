#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

void
cblas_dsyr2 (const enum CBLAS_ORDER order, const enum CBLAS_UPLO Uplo,
             const int N, const double alpha, const double *X, const int incX,
             const double *Y, const int incY, double *A, const int lda)
{
#define BASE double
#include "cblas_source_syr2.h"
#undef BASE
}
