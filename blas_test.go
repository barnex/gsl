package blas

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

import (
	"fmt"
)

func ExampleSDSDOT() {
	alpha := float32(4)
	X := []float32{2, 0}
	incX := 1
	Y := []float32{3, 5}
	incY := 1
	result := SDSDOT(alpha, X, incX, Y, incY)
	fmt.Println(result)

	//Output:
	//10
}

func ExampleDSDOT() {
	X := []float32{2, 0}
	incX := 1
	Y := []float32{3, 5}
	incY := 1
	result := DSDOT(X, incX, Y, incY)
	fmt.Println(result)

	//Output:
	//6
}

func ExampleSDOT() {
	X := []float32{2, 0}
	incX := 1
	Y := []float32{3, 5}
	incY := 1
	result := SDOT(X, incX, Y, incY)
	fmt.Println(result)

	//Output:
	//6
}

func ExampleDDOT() {
	X := []float64{2, 0}
	incX := 1
	Y := []float64{3, 5}
	incY := 1
	result := DDOT(X, incX, Y, incY)
	fmt.Println(result)

	//Output:
	//6
}

func ExampleCDOTU() {
	X := []complex64{complex(1,0), complex(1,1)}
	incX := 1
	Y := []complex64{complex(0,0), complex(2,0)}
	incY := 1
	result := CDOTU(X,incX, Y,incY)
	fmt.Println(result)

	//Output:
	//(2+2i)
}


func ExampleCDOTC() {
	X := []complex64{complex(1,0), complex(1,1)}
	incX := 1
	Y := []complex64{complex(0,0), complex(2,0)}
	incY := 1
	result := CDOTC(X,incX, Y,incY)
	fmt.Println(result)

	//Output:
	//(2-2i)
}


func ExampleZDOTU() {
	X := []complex128{complex(1,0), complex(1,1)}
	incX := 1
	Y := []complex128{complex(0,0), complex(2,0)}
	incY := 1
	result := ZDOTU(X,incX, Y,incY)
	fmt.Println(result)

	//Output:
	//(2+2i)
}


func ExampleZDOTC() {
	X := []complex128{complex(1,0), complex(1,1)}
	incX := 1
	Y := []complex128{complex(0,0), complex(2,0)}
	incY := 1
	result := ZDOTC(X,incX, Y,incY)
	fmt.Println(result)

	//Output:
	//(2-2i)
}

func ExampleSNRM2() {
	X := []float32{1, 666, 1, 666}
	incX := 2
	result := SNRM2(X, incX)
	fmt.Println(result)

	//Output:
	//1.4142135
}

func ExampleSASUM() {
	X := []float32{1, 666, 2, 666}
	incX := 2
	result := SASUM(X, incX)
	fmt.Println(result)

	//Output:
	//3
}

func ExampleDNRM2() {
	X := []float64{1, 666, 1, 666}
	incX := 2
	result := DNRM2(X, incX)
	fmt.Println(result)

	//Output:
	//1.4142135623730951
}


func ExampleDASUM() {
	X := []float64{1, 666, 2, 666}
	incX := 2
	result := DASUM(X, incX)
	fmt.Println(result)

	//Output:
	//3
}


func ExampleSCNRM2() {
	X := []complex64{complex(1,0), complex(0,1)}
	incX := 1
	result := SCNRM2(X, incX)
	fmt.Println(result)

	//Output:
	//1.4142135
}

func ExampleSCASUM() {
	X := []complex64{complex(1,0), complex(1,1)}
	incX := 1
	result := SCASUM(X, incX)
	fmt.Println(result)

	//Output:
	//3
}

func ExampleDZNRM2() {
	X := []complex128{complex(1,0), complex(0,1)}
	incX := 1
	result := DZNRM2(X, incX)
	fmt.Println(result)

	//Output:
	//1.4142135623730951
}

func ExampleDZASUM() {
	X := []complex128{complex(1,0), complex(1,1)}
	incX := 1
	result := DZASUM(X, incX)
	fmt.Println(result)

	//Output:
	//3
}

func ExampleISAMAX() {
	X := []float32{1, -999, 2, 3, 999}
	incX := 1
	result := ISAMAX(X, incX)
	fmt.Println(result)

	//Output:
	//1
}

func ExampleIDAMAX() {
	X := []float32{1, -999, 2, 3, 999}
	incX := 1
	result := ISAMAX(X, incX)
	fmt.Println(result)

	//Output:
	//1
}

func ExampleICAMAX() {
	X := []complex64{complex(2,0), complex(2,2), complex(0,2)}
	incX := 1
	result := ICAMAX(X, incX)
	fmt.Println(result)

	//Output:
	//1
}

func ExampleIZAMAX() {
	X := []complex128{complex(2,0), complex(2,2), complex(0,2)}
	incX := 1
	result := IZAMAX(X, incX)
	fmt.Println(result)

	//Output:
	//1
}

func ExampleSSWAP() {
	X := []float32{1, 666, 2, 666}
	incX := 2
	Y := []float32{-1, -2}
	incY := 1
	SSWAP(X,incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[-1 666 -2 666] [1 2]
}

/*
func ExampleSCOPY() {
	X := []float32{0, 0, 0}
	Y := []float32{0, 0, 0}
	result := SCOPY(X, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSAXPY() {
	alpha := float32(2.0)
	X := []float32{0, 0, 0}
	Y := []float32{0, 0, 0}
	result := SAXPY(alpha, X, Y)
	fmt.Println(result)

	//Output:
	//
}
*/

func ExampleDSWAP() {
	X := []float64{1, 666, 2, 666}
	incX := 2
	Y := []float64{-1, -2}
	incY := 1
	DSWAP(X,incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[-1 666 -2 666] [1 2]
}

/*

func ExampleDCOPY() {
	X := []float64{0, 0, 0}
	Y := []float64{0, 0, 0}
	result := DCOPY(X, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDAXPY() {
	alpha := float64(2.0)
	X := []float64{0, 0, 0}
	Y := []float64{0, 0, 0}
	result := DAXPY(alpha, X, Y)
	fmt.Println(result)

	//Output:
	//
}

*/

func ExampleCSWAP() {
	X := []complex64{1, 666, 2, 666}
	incX := 2
	Y := []complex64{-1, -2}
	incY := 1
	CSWAP(X,incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[(-1+0i) (666+0i) (-2+0i) (666+0i)] [(1+0i) (2+0i)]
}
/*
func ExampleCCOPY() {
	X := []complex64{0, 0, 0}
	Y := []complex64{0, 0, 0}
	result := CCOPY(X, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCAXPY() {
	alpha := complex64(complex(2.0, 3.0))
	X := []complex64{0, 0, 0}
	Y := []complex64{0, 0, 0}
	result := CAXPY(alpha, X, Y)
	fmt.Println(result)

	//Output:
	//
}
*/
func ExampleZSWAP() {
	X := []complex128{1, 666, 2, 666}
	incX := 2
	Y := []complex128{-1, -2}
	incY := 1
	ZSWAP(X,incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[(-1+0i) (666+0i) (-2+0i) (666+0i)] [(1+0i) (2+0i)]
}

/*
func ExampleZCOPY() {
	X := []complex128{0, 0, 0}
	Y := []complex128{0, 0, 0}
	result := ZCOPY(X, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZAXPY() {
	alpha := complex128(complex(2.0, 3.0))
	X := []complex128{0, 0, 0}
	Y := []complex128{0, 0, 0}
	result := ZAXPY(alpha, X, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSROTG() {
	a := *float32(2.0)
	b := *float32(2.0)
	c := *float32(2.0)
	s := *float32(2.0)
	result := SROTG(a, b, c, s)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSROTMG() {
	d1 := *float32(2.0)
	d2 := *float32(2.0)
	b1 := *float32(2.0)
	b2 := float32(2.0)
	P := *float32(2.0)
	result := SROTMG(d1, d2, b1, b2, P)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSROT() {
	X := []float32{0, 0, 0}
	Y := []float32{0, 0, 0}
	c := float32(2.0)
	s := float32(2.0)
	result := SROT(X, Y, c, s)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSROTM() {
	X := []float32{0, 0, 0}
	Y := []float32{0, 0, 0}
	P := *float32(2.0)
	result := SROTM(X, Y, P)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDROTG() {
	a := *float64(2.0)
	b := *float64(2.0)
	c := *float64(2.0)
	s := *float64(2.0)
	result := DROTG(a, b, c, s)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDROTMG() {
	d1 := *float64(2.0)
	d2 := *float64(2.0)
	b1 := *float64(2.0)
	b2 := float64(2.0)
	P := *float64(2.0)
	result := DROTMG(d1, d2, b1, b2, P)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDROT() {
	X := []float64{0, 0, 0}
	Y := []float64{0, 0, 0}
	c := float64(2.0)
	s := float64(2.0)
	result := DROT(X, Y, c, s)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDROTM() {
	X := []float64{0, 0, 0}
	Y := []float64{0, 0, 0}
	P := *float64(2.0)
	result := DROTM(X, Y, P)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSSCAL() {
	alpha := float32(2.0)
	X := []float32{0, 0, 0}
	SSCAL(alpha, X)
	fmt.Println()

	//Output:
	//
}

func ExampleDSCAL() {
	alpha := float64(2.0)
	X := []float64{0, 0, 0}
	DSCAL(alpha, X)
	fmt.Println()

	//Output:
	//
}

func ExampleCSCAL() {
	alpha := complex64(complex(2.0, 3.0))
	X := []complex64{0, 0, 0}
	CSCAL(alpha, X)
	fmt.Println()

	//Output:
	//
}

func ExampleZSCAL() {
	alpha := complex128(complex(2.0, 3.0))
	X := []complex128{0, 0, 0}
	ZSCAL(alpha, X)
	fmt.Println()

	//Output:
	//
}

func ExampleCSSCAL() {
	alpha := float32(2.0)
	X := []complex64{0, 0, 0}
	CSSCAL(alpha, X)
	fmt.Println()

	//Output:
	//
}

func ExampleZDSCAL() {
	alpha := float64(2.0)
	X := []complex128{0, 0, 0}
	ZDSCAL(alpha, X)
	fmt.Println()

	//Output:
	//
}

func ExampleSGEMV() {
	TransA := Transpose(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	X := []float32{0, 0, 0}
	beta := float32(2.0)
	Y := []float32{0, 0, 0}
	result := SGEMV(TransA, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSTRMV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeFloat32Matrix(2, 2)
	X := []float32{0, 0, 0}
	result := STRMV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSTRSV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeFloat32Matrix(2, 2)
	X := []float32{0, 0, 0}
	result := STRSV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDGEMV() {
	TransA := Transpose(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	X := []float64{0, 0, 0}
	beta := float64(2.0)
	Y := []float64{0, 0, 0}
	result := DGEMV(TransA, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDTRMV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeFloat64Matrix(2, 2)
	X := []float64{0, 0, 0}
	result := DTRMV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDTRSV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeFloat64Matrix(2, 2)
	X := []float64{0, 0, 0}
	result := DTRSV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCGEMV() {
	TransA := Transpose(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	X := []complex64{0, 0, 0}
	beta := complex64(complex(2.0, 3.0))
	Y := []complex64{0, 0, 0}
	result := CGEMV(TransA, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCTRMV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeComplex64Matrix(2, 2)
	X := []complex64{0, 0, 0}
	result := CTRMV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCTRSV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeComplex64Matrix(2, 2)
	X := []complex64{0, 0, 0}
	result := CTRSV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZGEMV() {
	TransA := Transpose(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	X := []complex128{0, 0, 0}
	beta := complex128(complex(2.0, 3.0))
	Y := []complex128{0, 0, 0}
	result := ZGEMV(TransA, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZTRMV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeComplex128Matrix(2, 2)
	X := []complex128{0, 0, 0}
	result := ZTRMV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZTRSV() {
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	A := MakeComplex128Matrix(2, 2)
	X := []complex128{0, 0, 0}
	result := ZTRSV(Uplo, TransA, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSSYMV() {
	Uplo := Uplo(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	X := []float32{0, 0, 0}
	beta := float32(2.0)
	Y := []float32{0, 0, 0}
	result := SSYMV(Uplo, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSGER() {
	alpha := float32(2.0)
	X := []float32{0, 0, 0}
	Y := []float32{0, 0, 0}
	A := MakeFloat32Matrix(2, 2)
	result := SGER(alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSSYR() {
	Uplo := Uplo(2.0)
	alpha := float32(2.0)
	X := []float32{0, 0, 0}
	A := MakeFloat32Matrix(2, 2)
	result := SSYR(Uplo, alpha, X, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSSYR2() {
	Uplo := Uplo(2.0)
	alpha := float32(2.0)
	X := []float32{0, 0, 0}
	Y := []float32{0, 0, 0}
	A := MakeFloat32Matrix(2, 2)
	result := SSYR2(Uplo, alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDSYMV() {
	Uplo := Uplo(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	X := []float64{0, 0, 0}
	beta := float64(2.0)
	Y := []float64{0, 0, 0}
	result := DSYMV(Uplo, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDGER() {
	alpha := float64(2.0)
	X := []float64{0, 0, 0}
	Y := []float64{0, 0, 0}
	A := MakeFloat64Matrix(2, 2)
	result := DGER(alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDSYR() {
	Uplo := Uplo(2.0)
	alpha := float64(2.0)
	X := []float64{0, 0, 0}
	A := MakeFloat64Matrix(2, 2)
	result := DSYR(Uplo, alpha, X, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDSYR2() {
	Uplo := Uplo(2.0)
	alpha := float64(2.0)
	X := []float64{0, 0, 0}
	Y := []float64{0, 0, 0}
	A := MakeFloat64Matrix(2, 2)
	result := DSYR2(Uplo, alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHEMV() {
	Uplo := Uplo(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	X := []complex64{0, 0, 0}
	beta := complex64(complex(2.0, 3.0))
	Y := []complex64{0, 0, 0}
	result := CHEMV(Uplo, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCGERU() {
	alpha := complex64(complex(2.0, 3.0))
	X := []complex64{0, 0, 0}
	Y := []complex64{0, 0, 0}
	A := MakeComplex64Matrix(2, 2)
	result := CGERU(alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCGERC() {
	alpha := complex64(complex(2.0, 3.0))
	X := []complex64{0, 0, 0}
	Y := []complex64{0, 0, 0}
	A := MakeComplex64Matrix(2, 2)
	result := CGERC(alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHER() {
	Uplo := Uplo(2.0)
	alpha := float32(2.0)
	X := []complex64{0, 0, 0}
	A := MakeComplex64Matrix(2, 2)
	result := CHER(Uplo, alpha, X, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHER2() {
	Uplo := Uplo(2.0)
	alpha := complex64(complex(2.0, 3.0))
	X := []complex64{0, 0, 0}
	Y := []complex64{0, 0, 0}
	A := MakeComplex64Matrix(2, 2)
	result := CHER2(Uplo, alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHEMV() {
	Uplo := Uplo(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	X := []complex128{0, 0, 0}
	beta := complex128(complex(2.0, 3.0))
	Y := []complex128{0, 0, 0}
	result := ZHEMV(Uplo, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZGERU() {
	alpha := complex128(complex(2.0, 3.0))
	X := []complex128{0, 0, 0}
	Y := []complex128{0, 0, 0}
	A := MakeComplex128Matrix(2, 2)
	result := ZGERU(alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZGERC() {
	alpha := complex128(complex(2.0, 3.0))
	X := []complex128{0, 0, 0}
	Y := []complex128{0, 0, 0}
	A := MakeComplex128Matrix(2, 2)
	result := ZGERC(alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHER() {
	Uplo := Uplo(2.0)
	alpha := float64(2.0)
	X := []complex128{0, 0, 0}
	A := MakeComplex128Matrix(2, 2)
	result := ZHER(Uplo, alpha, X, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHER2() {
	Uplo := Uplo(2.0)
	alpha := complex128(complex(2.0, 3.0))
	X := []complex128{0, 0, 0}
	Y := []complex128{0, 0, 0}
	A := MakeComplex128Matrix(2, 2)
	result := ZHER2(Uplo, alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSGEMM() {
	TransA := Transpose(2.0)
	TransB := Transpose(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	beta := float32(2.0)
	C := MakeFloat32Matrix(2, 2)
	result := SGEMM(TransA, TransB, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSSYMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	beta := float32(2.0)
	C := MakeFloat32Matrix(2, 2)
	result := SSYMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSSYRK() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	beta := float32(2.0)
	C := MakeFloat32Matrix(2, 2)
	result := SSYRK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSSYR2K() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	beta := float32(2.0)
	C := MakeFloat32Matrix(2, 2)
	result := SSYR2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSTRMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	result := STRMM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	result := STRSM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDGEMM() {
	TransA := Transpose(2.0)
	TransB := Transpose(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	beta := float64(2.0)
	C := MakeFloat64Matrix(2, 2)
	result := DGEMM(TransA, TransB, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDSYMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	beta := float64(2.0)
	C := MakeFloat64Matrix(2, 2)
	result := DSYMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDSYRK() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	beta := float64(2.0)
	C := MakeFloat64Matrix(2, 2)
	result := DSYRK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDSYR2K() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	beta := float64(2.0)
	C := MakeFloat64Matrix(2, 2)
	result := DSYR2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDTRMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	result := DTRMM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	result := DTRSM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCGEMM() {
	TransA := Transpose(2.0)
	TransB := Transpose(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := complex64(complex(2.0, 3.0))
	C := MakeComplex64Matrix(2, 2)
	result := CGEMM(TransA, TransB, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCSYMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := complex64(complex(2.0, 3.0))
	C := MakeComplex64Matrix(2, 2)
	result := CSYMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCSYRK() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	beta := complex64(complex(2.0, 3.0))
	C := MakeComplex64Matrix(2, 2)
	result := CSYRK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCSYR2K() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := complex64(complex(2.0, 3.0))
	C := MakeComplex64Matrix(2, 2)
	result := CSYR2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCTRMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	result := CTRMM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	result := CTRSM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZGEMM() {
	TransA := Transpose(2.0)
	TransB := Transpose(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := complex128(complex(2.0, 3.0))
	C := MakeComplex128Matrix(2, 2)
	result := ZGEMM(TransA, TransB, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZSYMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := complex128(complex(2.0, 3.0))
	C := MakeComplex128Matrix(2, 2)
	result := ZSYMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZSYRK() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	beta := complex128(complex(2.0, 3.0))
	C := MakeComplex128Matrix(2, 2)
	result := ZSYRK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZSYR2K() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := complex128(complex(2.0, 3.0))
	C := MakeComplex128Matrix(2, 2)
	result := ZSYR2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZTRMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	result := ZTRMM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	TransA := Transpose(2.0)
	Diag := Diag(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	result := ZTRSM(Side, Uplo, TransA, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHEMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := complex64(complex(2.0, 3.0))
	C := MakeComplex64Matrix(2, 2)
	result := CHEMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHERK() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := float32(2.0)
	A := MakeComplex64Matrix(2, 2)
	beta := float32(2.0)
	C := MakeComplex64Matrix(2, 2)
	result := CHERK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHER2K() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := float32(2.0)
	C := MakeComplex64Matrix(2, 2)
	result := CHER2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHEMM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := complex128(complex(2.0, 3.0))
	C := MakeComplex128Matrix(2, 2)
	result := ZHEMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHERK() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := float64(2.0)
	A := MakeComplex128Matrix(2, 2)
	beta := float64(2.0)
	C := MakeComplex128Matrix(2, 2)
	result := ZHERK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHER2K() {
	Uplo := Uplo(2.0)
	Trans := Transpose(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := float64(2.0)
	C := MakeComplex128Matrix(2, 2)
	result := ZHER2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}
*/
