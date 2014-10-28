#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_srotmg (float *d1, float *d2, float *b1, const float b2, float *P)
{
#define BASE float
#include "cblas_source_rotmg.h"
#undef BASE
}
