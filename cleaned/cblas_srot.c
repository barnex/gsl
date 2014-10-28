#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_srot (const int N, float *X, const int incX, float *Y, const int incY,
            const float c, const float s)
{
#define BASE float
#include "cblas_source_rot.h"
#undef BASE
}
