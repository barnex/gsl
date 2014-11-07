# GSL [![GoDoc](https://godoc.org/github.com/barnex/gsl?status.svg)](https://godoc.org/github.com/barnex/gsl) [![Build Status](https://travis-ci.org/barnex/gsl.svg?branch=master)](https://travis-ci.org/barnex/gsl)

Go wrappers for part of the GNU Scientific library. This package does not depend on a GSL installation (GSL source files are embedded in this go package and statically linked).

## License

The Go wrapper files are BSD-licensed and may freely be used in conjunction with any BLAS library. When you use this package in its entirety (including the GSL C sources) then GSL's GPL license takes precedence.

## BLAS [![GoDoc](https://godoc.org/github.com/barnex/gsl/blas?status.svg)](https://godoc.org/github.com/barnex/gsl/blas)
package BLAS (Basic Linear Algebra Subprograms) is open for public testing. All levels (1, 2 and 3) are available for all precissions (float32, float64, complex64, complex128).

Please report issues, if any. This package is not yet recommended for production use.

This package is go-gettable without depending on a BLAS installation: `go get -x github.com/barnex/gsl/blas`
(installation may take long due to compilation of C sources).
