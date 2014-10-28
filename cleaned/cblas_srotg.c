#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_srotg (float *a, float *b, float *c, float *s)
{
#define BASE float
#include "cblas_source_rotg.h"
#undef BASE
}
