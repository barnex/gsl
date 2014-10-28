package main

//#cgo LDFLAGS: -lm
//#include "cblas_gsl_cblas.h"
import "C"

func main(){
	println(C.CblasLeft)
}
