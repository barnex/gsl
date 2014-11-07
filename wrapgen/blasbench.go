//+build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"path"
	"reflect"
	"runtime"
	"strings"

	. "."
	. "../blas"
)

var funcs = []interface{}{CAXPY, CCOPY, CDOTC, CDOTU, CGEMM, CGEMV, CGERC, CGERU, CHEMM, CHEMV, CHER, CHER2, CHER2K, CHERK, CSCAL, CSSCAL, CSWAP, CSYMM, CSYR2K, CSYRK, CSize, CSubmatrix, CTRMM, CTRMV, CTRSM, CTRSV, DASUM, DAXPY, DCOPY, DDOT, DGEMM, DGEMV, DGER, DNRM2, DSCAL, DSDOT, DSWAP, DSYMM, DSYMV, DSYR, DSYR2, DSYR2K, DSYRK, DSize, DSubmatrix, DTRMM, DTRMV, DTRSM, DTRSV, DZASUM, DZNRM2, ICAMAX, IDAMAX, ISAMAX, IZAMAX, MakeComplex128Matrix, MakeComplex64Matrix, MakeFloat32Matrix, MakeFloat64Matrix, SASUM, SAXPY, SCASUM, SCNRM2, SCOPY, SDOT, SDSDOT, SGEMM, SGEMV, SGER, SNRM2, SSCAL, SSWAP, SSYMM, SSYMV, SSYR, SSYR2, SSYR2K, SSYRK, SSize, SSubmatrix, STRMM, STRMV, STRSM, STRSV, ZAXPY, ZCOPY, ZDOTC, ZDOTU, ZDSCAL, ZGEMM, ZGEMV, ZGERC, ZGERU, ZHEMM, ZHEMV, ZHER, ZHER2, ZHER2K, ZHERK, ZSCAL, ZSWAP, ZSYMM, ZSYR2K, ZSYRK, ZSize, ZSubmatrix, ZTRMM, ZTRMV, ZTRSM, ZTRSV}

func main() {
	flag.Parse()
	var fs []Function
	for _, f := range funcs {
		fv := reflect.ValueOf(f)
		ft := fv.Type()
		var args []string
		for i := 0; i < ft.NumIn(); i++ {
			args = append(args, noPkg(ft.In(i).String()))
		}
		name := runtime.FuncForPC(fv.Pointer()).Name()
		name = path.Ext(name)[1:]
		fs = append(fs, Function{name, args})
	}
	for _, f := range fs {
		fmt.Println(f)
	}

	Render(flag.Arg(0), testTempl, &Dta{fs})
}

type Function struct {
	Name string
	Args []string
}

type Dta struct{ Funcs []Function }

func (*Dta) Make(typ string) string {
	switch typ {
	default:
		log.Fatal("canot make ", typ)
	case "int":
		return "1"
	case "complex128":
		return "complex128(1)"
	case "[]complex128":
		return "make([]complex128, N)"
	case "[][]complex128":
		return "MakeComplex128Matrix(N, N)"
	case "complex64":
		return "complex64(1)"
	case "[]complex64":
		return "make([]complex64, N)"
	case "[][]complex64":
		return "MakeComplex64Matrix(N, N)"
	case "float32":
		return "float32(1)"
	case "[]float32":
		return "make([]float32, N)"
	case "[][]float32":
		return "MakeFloat32Matrix(N, N)"
	case "float64":
		return "float64(1)"
	case "[]float64":
		return "make([]float64, N)"
	case "[][]float64":
		return "MakeFloat64Matrix(N, N)"
	case "Transpose":
		return "NoTrans"
	case "Uplo":
		return "Upper"
	case "Diag":
		return "NonUnit"
	case "Side":
		return "Left"
	}
	return "$err$"
}

func noPkg(s string) string {
	if strings.HasPrefix(s, "blas.") {
		return s[len("blas."):]
	} else {
		return s
	}
}

const testTempl = `
package blas

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

import(
	"testing"
)

const N = 1024 // size of all vectors and matrices

{{range .Funcs}}
	func Benchmark{{.Name}}(b*testing.B) { 
			b.StopTimer()
			{{range $i,$a:=.Args}}
				arg{{$i}} := {{$.Make $a}} {{end}}
			b.StartTimer()
			for i:=0; i<b.N; i++{
				{{.Name}}( {{range $i,$a:=.Args}} {{if ne $i 0}},{{end}} arg{{$i}} {{end}})
			}
	}
{{end}}

`
