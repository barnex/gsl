#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_drotmg (double *d1, double *d2, double *b1, const double b2, double *P)
{
#define BASE double
#include "cblas_source_rotmg.h"
#undef BASE
}
