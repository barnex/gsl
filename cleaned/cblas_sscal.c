#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_sscal (const int N, const float alpha, float *X, const int incX)
{
#define BASE float
#include "cblas_source_scal_r.h"
#undef BASE
}
