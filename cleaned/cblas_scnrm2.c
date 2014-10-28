#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

float
cblas_scnrm2 (const int N, const void *X, const int incX)
{
#define BASE float
#include "cblas_source_nrm2_c.h"
#undef BASE
}
