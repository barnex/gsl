// +build ignore

// This program generates Go wrappers for cblas sources

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/scanner"
	"text/template"
)

func main() {
	flag.Parse()
	for _, fname := range flag.Args() {
		cuda2go(fname)
	}
}

type Func struct {
	CName         string
	GoType, CType string
	Args          []arg
}

func cuda2go(fname string) {
	// open cuda file
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read tokens
	var token []string
	var s scanner.Scanner
	s.Init(f)
	tok := s.Scan()
	for tok != scanner.EOF {
		if !filter[s.TokenText()] {
			if s.TokenText() == "*" {
				token[len(token)-1] += s.TokenText()
			} else {
				token = append(token, s.TokenText())
			}
		}
		tok = s.Scan()
	}

	funcs := make(map[string]Func)
	// parse tokens
	for i := 0; i < len(token); i++ {
		funcname := ""
		funcret := ""
		if token[i] == "(" { // begin of function
			funcname = token[i-1]
			funcret = token[i-2]
			args := parseArgs(token, &i)
			goname := strings.ToUpper(funcname[len("cblas_"):])
			funcs[goname] = Func{funcname, typemap(funcret), funcret, args}
		}
	}

	{
		out, errOut := os.Create("cblas_wrapper.go")
		if errOut != nil {
			log.Fatal(errOut)
		}
		errT := templ.Execute(out, &TData{funcs})
		if errT != nil {
			log.Fatal(errT)
		}
	}

	{
		out, errOut := os.Create("blas_wrapper.go")
		if errOut != nil {
			log.Fatal(errOut)
		}
		errT := templateSafe.Execute(out, &TData{funcs})
		if errT != nil {
			log.Fatal(errT)
		}
	}

}

type TData struct {
	Funcs map[string]Func
}

func (*TData) TrimComma(s string) string {
	return strings.Trim(s, ",")
}

func (*TData) Cast(ctype, arg string) string {
	switch {
	default:
		return cast(ctype) + "(" + arg + ")"
	case strings.HasPrefix(ctype, "CBLAS_"):
		return "uint32(" + arg + ")"
	case ctype == "void*":
		return arg
	}
}

func (*TData) CBlasCast(gotype, arg string) string {
	switch {
	default:
		return arg
	case gotype == "Order" || gotype == "Transpose":
		return "uint32(" + arg + ")"
	}
}

var castm = map[string]string{
	"int":     "C.int",
	"float":   "C.float",
	"float*":  "(*C.float)",
	"double":  "C.double",
	"double*": "(*C.double)",
	"void*":   "",
}

func cast(t string) string {
	if c, ok := castm[t]; ok {
		return c
	} else {
		panic("cannot cast " + t)
	}
}

type arg struct {
	GoType, CType, Name string
}

func newArg(typ, name string) arg {
	return arg{typemap(typ), typ, name}
}

func parseArgs(token []string, i *int) []arg {
	var args []arg
	for {
		tok := token[*i]
		if tok == ")" {
			if token[*i-1] != ")" {
				args = append(args, newArg(token[*i-2], token[*i-1]))
			}
			return args
		}
		if tok == "," {
			args = append(args, newArg(token[*i-2], token[*i-1]))
		}
		*i++
	}
	log.Fatal("syntax error")
	return nil
}

// translate C type to Go type.
func typemap(ctype string) string {
	if gotype, ok := tm[ctype]; ok {
		return gotype
	}
	log.Fatal(fmt.Errorf("unsupported cuda type: %v", ctype))
	return "666"
}

var tm = map[string]string{
	"float":           "float32",
	"float*":          "*float32",
	"double":          "float64",
	"double*":         "*float64",
	"int":             "int",
	"void*":           "unsafe.Pointer",
	"void":            "",
	"CBLAS_ORDER":     "uint32",
	"CBLAS_TRANSPOSE": "uint32",
	"CBLAS_UPLO":      "uint32",
	"CBLAS_DIAG":      "uint32",
	"CBLAS_SIDE":      "uint32",
	"CBLAS_INDEX":     "int",
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

{{range $k,$v:=.Funcs}}
func {{$k}} ( {{range $i,$a := $v.Args}} {{if ne $i 0}},{{end}} {{$a.Name}} {{$a.GoType}} {{end}}  ) {{$v.GoType}}{
{{with $v.GoType}} return {{.}} ( {{end}} C.{{$v.CName}}( {{range $i,$v := .Args}} {{if ne $i 0}}, {{end}} {{$.Cast .CType .Name}} {{end}}) {{with $v.GoType}} ) {{end}}
}
{{end}}

`

const templSafe = ` package blas

import(
	"cblas"
	"unsafe"
)

type Order int
type Transpose int
type Uplo int
type Diag int
type Side int

{{range $k,$v:=.Funcs}}
func {{$k}} ( {{range $i,$a := $v.Args}} {{if ne $i 0}},{{end}} {{$a.Name}} {{$a.GoType}} {{end}}  ) {{$v.GoType}}{
{{with $v.GoType}} return {{.}} ( {{end}} C.{{$v.CName}}(
	{{range $i,$v := .Args}} {{if ne $i 0}},
{{end}} {{$.CBlasCast .GoType .Name}} {{end}}) {{with $v.GoType}} ) {{end}}
}
{{end}}

`

// wrapper code template
var templ = template.Must(template.New("wrap").Parse(templText))
var templateSafe = template.Must(template.New("wrap").Parse(templSafe))

// should token be filtered out of stream?
var filter = map[string]bool{
	"const": true,
	"enum":  true,
}
