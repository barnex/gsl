#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_zswap (const int N, void *X, const int incX, void *Y, const int incY)
{
#define BASE double
#include "cblas_source_swap_c.h"
#undef BASE
}
