#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"
#include "cblas_error_cblas_l2.h"

void
cblas_cher (const enum CBLAS_ORDER order, const enum CBLAS_UPLO Uplo,
            const int N, const float alpha, const void *X, const int incX,
            void *A, const int lda)
{
#define BASE float
#include "cblas_source_her.h"
#undef BASE
}