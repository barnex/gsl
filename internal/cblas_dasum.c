#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

double
cblas_dasum (const int N, const double *X, const int incX)
{
#define BASE double
#include "cblas_source_asum_r.h"
#undef BASE
}
