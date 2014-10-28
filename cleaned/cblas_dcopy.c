#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_dcopy (const int N, const double *X, const int incX, double *Y,
             const int incY)
{
#define BASE double
#include "cblas_source_copy_r.h"
#undef BASE
}
