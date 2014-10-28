#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

float
cblas_sasum (const int N, const float *X, const int incX)
{
#define BASE float
#include "cblas_source_asum_r.h"
#undef BASE
}
