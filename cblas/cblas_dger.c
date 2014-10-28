#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

void
cblas_dger (const enum CBLAS_ORDER order, const int M, const int N,
            const double alpha, const double *X, const int incX,
            const double *Y, const int incY, double *A, const int lda)
{
#define BASE double
#include "cblas_source_ger.h"
#undef BASE
}
