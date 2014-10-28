#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_cdotu_sub (const int N, const void *X, const int incX, const void *Y,
             const int incY, void *result)
{
#define BASE float
#define CONJ_SIGN 1.0
#include "cblas_source_dot_c.h"
#undef CONJ_SIGN
#undef BASE
}