#! /bin/bash

# No need to run this script unless you update the underlying gsl version or wrapgen.

#go run wrapgen/cblaswrap.go cblas/gsl_gsl_cblas.h cblas/cblas.go
#gofmt -w cblas/cblas.go

go run wrapgen/blaswrap.go cblas/gsl_gsl_blas.h cblas/gsl_gsl_cblas.h blas.go.tmp blas_test.go
gofmt -w blas.go.tmp blas_test.go
