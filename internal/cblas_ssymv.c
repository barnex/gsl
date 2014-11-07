#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

void
cblas_ssymv (const enum CBLAS_ORDER order, const enum CBLAS_UPLO Uplo,
             const int N, const float alpha, const float *A, const int lda,
             const float *X, const int incX, const float beta, float *Y,
             const int incY)
{
#define BASE float
#include "cblas_source_symv.h"
#undef BASE
}
