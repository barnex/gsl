#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

void
cblas_sspr2 (const enum CBLAS_ORDER order, const enum CBLAS_UPLO Uplo,
             const int N, const float alpha, const float *X, const int incX,
             const float *Y, const int incY, float *Ap)
{
#define BASE double
#include "cblas_source_spr2.h"
#undef BASE
}
