#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_drotm (const int N, double *X, const int incX, double *Y,
             const int incY, const double *P)
{
#define BASE double
#include "cblas_source_rotm.h"
#undef BASE
}
