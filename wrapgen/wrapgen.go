package wrapgen

import (
	"log"
	"os"
	"strings"
	"text/scanner"
	"text/template"
)

// Read header file, split in tokens, ignore words from filter map.
func Tokenize(fname string) []string {
	f, err := os.Open(fname)
	check(err)
	defer f.Close()

	// read tokens
	var tokens []string
	var s scanner.Scanner
	s.Init(f)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {

		if FilterTokens[s.TokenText()] {
			continue
		}

		switch {
		default:
			tokens = append(tokens, s.TokenText())
		case s.TokenText() == "*":
			tokens[len(tokens)-1] += s.TokenText()
		}
	}

	return tokens
}

// These tokens are ignored
var FilterTokens = map[string]bool{
	"const": true,
	"enum":  true,
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Function prototype
type Func struct {
	CName string
	CType string
	Args  []Arg
}

// Function argument
type Arg struct {
	CType, Name string
}

// Parse function prototypes
func ParseFuncs(tokens []string) []Func {
	var funcs []Func
	for i := 0; i < len(tokens); i++ {
		funcname := ""
		funcret := ""
		if tokens[i] == "(" { // begin of function
			funcname = tokens[i-1]
			funcret = tokens[i-2]
			args := parseArgs(tokens, &i)
			funcs = append(funcs, Func{funcname, funcret, args})
		}
	}
	return funcs
}

func parseArgs(token []string, i *int) []Arg {
	var args []Arg
	for {
		tok := token[*i]
		if tok == "," || (tok == ")" && token[*i-1] != "(") {
			a := Arg{token[*i-2], token[*i-1]}
			if a.CType == "[" {
				a = Arg{token[*i-4] + "*", token[*i-3]}

			}
			args = append(args, a)
		}
		if tok == ")" {
			return args
		}
		*i++
	}
	log.Fatal("syntax error")
	return nil
}

// Render template to file
func Render(fname string, templText string, data interface{}) {
	out, errOut := os.Create(fname)
	check(errOut)
	defer out.Close()
	t := template.Must(template.New("").Parse(templText))
	errT := t.Execute(out, data)
	check(errT)
}

// Utility type, pass to template rendering
type TData struct {
	Funcs []Func
}

// Map C type to Go type, using TypeMap lookup. E.g.:
// 	float -> float32
func (*TData) GoType(ctype string) string {
	if v, ok := CtoGoType[ctype]; ok {
		return v
	}

	if strings.HasSuffix(ctype, "*") {
		return "*" + (*TData).GoType(nil, ctype[:len(ctype)-1])
	}

	log.Fatal("Cannot convert C type to Go: " + ctype)
	return ""
}

// maps C to Go types
var CtoGoType = map[string]string{
	"float":  "float32",
	"double": "float64",
	"int":    "int",
	"void*":  "unsafe.Pointer",
	"void":   "",
}

// Map C type to cgo type. E.g.:
// 	int -> C.int
func (*TData) CgoType(ctype string) string {
	if v, ok := CtoCgoType[ctype]; ok {
		return v
	}
	nPointers := 0
	for strings.HasSuffix(ctype, "*") {
		ctype = ctype[:len(ctype)-1]
		nPointers++
	}
	return "*********"[:nPointers] + "C." + ctype
}

var CtoCgoType = map[string]string{
	"void*": "unsafe.Pointer",
}

func (*TData) Strip(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

// Turn C function name into idiomatic, exported Go name.
func (*TData) GoName(cname string) string {
	return strings.ToUpper(cname)
}
