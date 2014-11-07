#include "gsl_gsl_math.h"
#include "gsl_gsl_cblas.h"
#include "cblas_cblas.h"

void
cblas_drot (const int N, double *X, const int incX, double *Y, const int incY,
            const double c, const double s)
{
#define BASE double
#include "cblas_source_rot.h"
#undef BASE
}
