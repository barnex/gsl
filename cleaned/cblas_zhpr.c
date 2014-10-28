#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

void
cblas_zhpr (const enum CBLAS_ORDER order, const enum CBLAS_UPLO Uplo,
            const int N, const double alpha, const void *X, const int incX,
            void *Ap)
{
#define BASE double
#include "cblas_source_hpr.h"
#undef BASE
}
