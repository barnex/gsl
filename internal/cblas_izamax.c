#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

CBLAS_INDEX
cblas_izamax (const int N, const void *X, const int incX)
{
#define BASE double
#include "cblas_source_iamax_c.h"
#undef BASE
}
