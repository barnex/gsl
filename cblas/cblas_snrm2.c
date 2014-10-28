#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

float
cblas_snrm2 (const int N, const float *X, const int incX)
{
#define BASE float
#include "cblas_source_nrm2_r.h"
#undef BASE
}
