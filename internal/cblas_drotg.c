#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_drotg (double *a, double *b, double *c, double *s)
{
#define BASE double
#include "cblas_source_rotg.h"
#undef BASE
}
