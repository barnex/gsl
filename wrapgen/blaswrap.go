//+build ignore

package main

import (
	"flag"

	. "."
)

func main() {
	flag.Parse()
	tokens := Tokenize(flag.Arg(0))

	CtoGoType["CBLAS_ORDER_t"] = "Order"
	CtoGoType["CBLAS_TRANSPOSE_t"] = "Transpose"
	CtoGoType["CBLAS_UPLO_t"] = "Uplo"
	CtoGoType["CBLAS_DIAG_t"] = "Diag"
	CtoGoType["CBLAS_SIDE_t"] = "Side"
	CtoGoType["CBLAS_INDEX_t"] = "int"

	CtoGoType["CBLAS_ORDER"] = "uint32"
	CtoGoType["CBLAS_TRANSPOSE"] = "uint32"
	CtoGoType["CBLAS_UPLO"] = "uint32"
	CtoGoType["CBLAS_DIAG"] = "uint32"
	CtoGoType["CBLAS_SIDE"] = "uint32"
	CtoGoType["CBLAS_INDEX"] = "int"
	CtoGoType[""] = ""

	CtoGoType["gsl_complex"] = "complex128"
	CtoGoType["gsl_complex_float"] = "complex64"

	CtoGoType["gsl_vector*"] = "[]float64"
	CtoGoType["gsl_vector_float*"] = "[]float32"
	CtoGoType["gsl_vector_complex*"] = "[]complex128"
	CtoGoType["gsl_vector_complex_float*"] = "[]complex64"

	CtoGoType["gsl_matrix*"] = "[][]float64"
	CtoGoType["gsl_matrix_float*"] = "[][]float32"
	CtoGoType["gsl_matrix_complex*"] = "[][]complex128"
	CtoGoType["gsl_matrix_complex_float*"] = "[][]complex64"

	blasFuncs := ParseFuncs(tokens)

	// remove "result *something" arg
	for i,f:=range blasFuncs{
			args := f.Args
		if args[len(args)-1].Name == "result"{
				blasFuncs[i].Args = args[:len(args)-1]
		}
	}

	tokens = Tokenize(flag.Arg(1))
	cblasFuncs := ParseFuncs(tokens)

	Render(flag.Arg(2), templText, &TData2{TData{blasFuncs}, cblasFuncs})
}

type TData2 struct {
	TData
	CBlasFuncs []Func
}

func (x *TData2) CBlasFunc(blasFname string) Func {
	cblasFname := "cblas_" + x.Strip(blasFname, "gsl_blas_")
	for _, c := range x.CBlasFuncs {
		if c.CName == cblasFname {
			return c
		}
	}
	//log.Fatal("no CBlasFunc for " + blasFname + "->" + cblasFname, "in", x.CBlasFuncs)
	return Func{}
}

// wrapper code template text
const templText = `package blas

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

import(
	"github.com/barnex/blas/cblas"
)

{{range .Funcs}}
	{{$cf:=($.CBlasFunc .CName)}}
	func {{$.Strip ($.GoName .CName) "GSL_BLAS_"}}({{range $i,$a:=.Args}}{{if ne $i 0}},{{end}}{{$a.Name}} {{$.GoType $a.CType}}{{end}}) {{$.GoType $cf.CType}}{ {{range $cf.Args}}
				var {{.Name}}_ {{$.GoType .CType}} = 0 {{end}}
		{{with $cf.CName | $.GoName}}cblas.{{.}}( {{range $i,$a:=$cf.Args}}{{if ne $i 0}},{{end}}{{$a.Name}}_ {{end}} ){{end}}
	}
{{end}}

`
