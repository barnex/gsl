#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_saxpy (const int N, const float alpha, const float *X, const int incX,
             float *Y, const int incY)
{
#define BASE float
#include "cblas_source_axpy_r.h"
#undef BASE
}
