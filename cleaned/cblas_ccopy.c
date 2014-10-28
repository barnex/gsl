#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_ccopy (const int N, const void *X, const int incX, void *Y,
             const int incY)
{
#define BASE float
#include "cblas_source_copy_c.h"
#undef BASE
}
