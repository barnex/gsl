#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_dscal (const int N, const double alpha, double *X, const int incX)
{
#define BASE double
#include "cblas_source_scal_r.h"
#undef BASE
}
