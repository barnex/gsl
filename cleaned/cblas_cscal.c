#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_cscal (const int N, const void *alpha, void *X, const int incX)
{
#define BASE float
#include "cblas_source_scal_c.h"
#undef BASE
}
