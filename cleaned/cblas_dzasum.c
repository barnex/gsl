#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

double
cblas_dzasum (const int N, const void *X, const int incX)
{
#define BASE double
#include "cblas_source_asum_c.h"
#undef BASE
}
