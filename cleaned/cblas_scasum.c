#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

float
cblas_scasum (const int N, const void *X, const int incX)
{
#define BASE float
#include "cblas_source_asum_c.h"
#undef BASE
}
