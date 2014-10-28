#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_csscal (const int N, const float alpha, void *X, const int incX)
{
#define BASE float
#include "cblas_source_scal_c_s.h"
#undef BASE
}
