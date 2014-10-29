#! /bin/bash

# No need to run this script unless you update the underlying gsl version or wrapgen.

go run wrapgen/cblaswrap.go cblas/gsl_gsl_cblas.h cblas/cblas.go
gofmt -w cblas/cblas.go

