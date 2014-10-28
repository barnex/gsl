#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

CBLAS_INDEX
cblas_isamax (const int N, const float *X, const int incX)
{
#define BASE float
#include "cblas_source_iamax_r.h"
#undef BASE
}
