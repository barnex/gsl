#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

double
cblas_dnrm2 (const int N, const double *X, const int incX)
{
#define BASE double
#include "cblas_source_nrm2_r.h"
#undef BASE
}
