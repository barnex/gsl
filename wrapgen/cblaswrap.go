//+build ignore

package main

import(
		. "."
		"flag"
)

func main() {
	flag.Parse()
	tokens:= Tokenize(flag.Arg(0))

	CtoGoType["CBLAS_ORDER"] = "uint32"
	CtoGoType["CBLAS_TRANSPOSE"] = "uint32"
	CtoGoType["CBLAS_UPLO"] = "uint32"
	CtoGoType["CBLAS_DIAG"] = "uint32"
	CtoGoType["CBLAS_SIDE"] = "uint32"
	CtoGoType["CBLAS_INDEX"] = "uint32"

	CtoCgoType["CBLAS_ORDER"] = "uint32"
	CtoCgoType["CBLAS_TRANSPOSE"] = "uint32"
	CtoCgoType["CBLAS_UPLO"] = "uint32"
	CtoCgoType["CBLAS_DIAG"] = "uint32"
	CtoCgoType["CBLAS_SIDE"] = "uint32"
	CtoCgoType["CBLAS_INDEX"] = "uint32"

	funcs := ParseFuncs(tokens)

	Render(flag.Arg(1), templText, &TData{funcs})
}


// wrapper code template text
const templText = `//Package cblas provides an unsafe, low-level C interface used by package blas.
//Users should use package blas directly.
package cblas

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

//#cgo LDFLAGS: -lm
//#cgo CFLAGS: -O3
//#include "gsl_gsl_cblas.h"
import "C"

import(
	"unsafe"
)

{{range .Funcs}}
	func {{$.GoName .CName}}({{range $i,$a:=.Args}}{{if ne $i 0}},{{end}}{{$a.Name}} {{$.GoType $a.CType}}{{end}}) {{$.GoType .CType}}{
		{{with $.GoType .CType}}return ({{.}})({{end}}C.{{.CName}}( {{range $i,$a:=.Args}} {{if ne $i 0}},{{end}} ({{$.CgoType $a.CType}})({{$a.Name}}) {{end}}){{with $.GoType .CType}}){{end}}	
	}
{{end}}

`
