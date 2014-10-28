#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_scopy (const int N, const float *X, const int incX, float *Y,
             const int incY)
{
#define BASE float
#include "cblas_source_copy_r.h"
#undef BASE
}
