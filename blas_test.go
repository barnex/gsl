package blas

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

import (
	"fmt"
	"math"
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
	X := []complex64{complex(1, 0), complex(1, 1)}
	incX := 1
	Y := []complex64{complex(0, 0), complex(2, 0)}
	incY := 1
	result := CDOTU(X, incX, Y, incY)
	fmt.Println(result)

	//Output:
	//(2+2i)
}

func ExampleCDOTC() {
	X := []complex64{complex(1, 0), complex(1, 1)}
	incX := 1
	Y := []complex64{complex(0, 0), complex(2, 0)}
	incY := 1
	result := CDOTC(X, incX, Y, incY)
	fmt.Println(result)

	//Output:
	//(2-2i)
}

func ExampleZDOTU() {
	X := []complex128{complex(1, 0), complex(1, 1)}
	incX := 1
	Y := []complex128{complex(0, 0), complex(2, 0)}
	incY := 1
	result := ZDOTU(X, incX, Y, incY)
	fmt.Println(result)

	//Output:
	//(2+2i)
}

func ExampleZDOTC() {
	X := []complex128{complex(1, 0), complex(1, 1)}
	incX := 1
	Y := []complex128{complex(0, 0), complex(2, 0)}
	incY := 1
	result := ZDOTC(X, incX, Y, incY)
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
	X := []complex64{complex(1, 0), complex(0, 1)}
	incX := 1
	result := SCNRM2(X, incX)
	fmt.Println(result)

	//Output:
	//1.4142135
}

func ExampleSCASUM() {
	X := []complex64{complex(1, 0), complex(1, 1)}
	incX := 1
	result := SCASUM(X, incX)
	fmt.Println(result)

	//Output:
	//3
}

func ExampleDZNRM2() {
	X := []complex128{complex(1, 0), complex(0, 1)}
	incX := 1
	result := DZNRM2(X, incX)
	fmt.Println(result)

	//Output:
	//1.4142135623730951
}

func ExampleDZASUM() {
	X := []complex128{complex(1, 0), complex(1, 1)}
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
	X := []complex64{complex(2, 0), complex(2, 2), complex(0, 2)}
	incX := 1
	result := ICAMAX(X, incX)
	fmt.Println(result)

	//Output:
	//1
}

func ExampleIZAMAX() {
	X := []complex128{complex(2, 0), complex(2, 2), complex(0, 2)}
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
	SSWAP(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[-1 666 -2 666] [1 2]
}

func ExampleSCOPY() {
	X := []float32{1, 666, 2, 666}
	incX := 2
	Y := []float32{-1, -2}
	incY := 1
	SCOPY(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[1 666 2 666] [1 2]
}

func ExampleSAXPY() {
	alpha := float32(2)
	X := []float32{1, 2, 3}
	incX := 1
	Y := []float32{4, 5, 6}
	incY := 1
	SAXPY(alpha, X, incX, Y, incY)
	fmt.Println(Y)

	//Output:
	//[6 9 12]
}

func ExampleDSWAP() {
	X := []float64{1, 666, 2, 666}
	incX := 2
	Y := []float64{-1, -2}
	incY := 1
	DSWAP(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[-1 666 -2 666] [1 2]
}

func ExampleDCOPY() {
	X := []float64{1, 666, 2, 666}
	incX := 2
	Y := []float64{-1, -2}
	incY := 1
	DCOPY(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[1 666 2 666] [1 2]
}

func ExampleDAXPY() {
	alpha := float64(2)
	X := []float64{1, 2, 3}
	incX := 1
	Y := []float64{4, 5, 6}
	incY := 1
	DAXPY(alpha, X, incX, Y, incY)
	fmt.Println(Y)

	//Output:
	//[6 9 12]
}

func ExampleCSWAP() {
	X := []complex64{1, 666, 2, 666}
	incX := 2
	Y := []complex64{-1, -2}
	incY := 1
	CSWAP(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[(-1+0i) (666+0i) (-2+0i) (666+0i)] [(1+0i) (2+0i)]
}

func ExampleCCOPY() {
	X := []complex64{1, 666, 2, 666}
	incX := 2
	Y := []complex64{-1, -2}
	incY := 1
	CCOPY(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[(1+0i) (666+0i) (2+0i) (666+0i)] [(1+0i) (2+0i)]
}

func ExampleCAXPY() {
	alpha := complex64(complex(0, 1))
	X := []complex64{1, 2, 3}
	incX := 1
	Y := []complex64{4, 5, 6}
	incY := 1
	CAXPY(alpha, X, incX, Y, incY)
	fmt.Println(Y)

	//Output:
	//[(4+1i) (5+2i) (6+3i)]
}

func ExampleZSWAP() {
	X := []complex128{1, 666, 2, 666}
	incX := 2
	Y := []complex128{-1, -2}
	incY := 1
	ZSWAP(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[(-1+0i) (666+0i) (-2+0i) (666+0i)] [(1+0i) (2+0i)]
}

func ExampleZCOPY() {
	X := []complex128{1, 666, 2, 666}
	incX := 2
	Y := []complex128{-1, -2}
	incY := 1
	ZCOPY(X, incX, Y, incY)
	fmt.Println(X, Y)

	//Output:
	//[(1+0i) (666+0i) (2+0i) (666+0i)] [(1+0i) (2+0i)]
}

func ExampleZAXPY() {
	alpha := complex128(complex(0, 1))
	X := []complex128{1, 2, 3}
	incX := 1
	Y := []complex128{4, 5, 6}
	incY := 1
	ZAXPY(alpha, X, incX, Y, incY)
	fmt.Println(Y)

	//Output:
	//[(4+1i) (5+2i) (6+3i)]
}

func ExampleSROTG() {
	a := float32(1)
	b := float32(1)
	r, c, s := SROTG(a, b)
	fmt.Println(r, c, s)

	//Output:
	//1.4142135 0.70710677 0.70710677
}

func ExampleSROTMG() {
	d1 := float32(1)
	d2 := float32(1)
	b1 := float32(1)
	b2 := float32(1)
	result := SROTMG(d1, d2, b1, b2)
	fmt.Println(result)

	//Output:
	//[1 1 0 0 1]
}

func ExampleSROT() {
	X := []float32{1, 0, 1}
	incX := 1
	Y := []float32{0, 1, 1}
	incY := 1
	theta := math.Pi / 4
	c := float32(math.Cos(theta))
	s := float32(math.Sin(theta))
	SROT(X, incX, Y, incY, c, s)
	fmt.Println(X, Y)

	//Output:
	//[0.70710677 0.70710677 1.4142135] [-0.70710677 0.70710677 0]
}

func ExampleSROTM() {
	d1 := float32(1.0)
	d2 := float32(1.0)
	b1 := float32(1.0)
	b2 := float32(1.0)
	P := SROTMG(d1, d2, b1, b2)

	X := []float32{1}
	incX := 1
	Y := []float32{1}
	incY := 1
	SROTM(X, incX, Y, incY, P)
	fmt.Println(X, Y)

	//Output:
	//[2] [0]
}

func ExampleDROTMG() {
	d1 := 1.0
	d2 := 1.0
	b1 := 1.0
	b2 := 1.0
	result := DROTMG(d1, d2, b1, b2)
	fmt.Println(result)

	//Output:
	//[1 1 0 0 1]
}

func ExampleDROT() {
	X := []float64{1, 0, 1}
	incX := 1
	Y := []float64{0, 1, 1}
	incY := 1
	theta := math.Pi / 4
	c := math.Cos(theta)
	s := math.Sin(theta)
	DROT(X, incX, Y, incY, c, s)
	fmt.Println(X, Y)

	//Output:
	//[0.7071067811865476 0.7071067811865475 1.414213562373095] [-0.7071067811865475 0.7071067811865476 1.1102230246251565e-16]
}

func ExampleDROTM() {
	d1 := 1.0
	d2 := 1.0
	b1 := 1.0
	b2 := 1.0
	P := DROTMG(d1, d2, b1, b2)

	X := []float64{1}
	incX := 1
	Y := []float64{1}
	incY := 1
	DROTM(X, incX, Y, incY, P)
	fmt.Println(X, Y)

	//Output:
	//[2] [0]
}

func ExampleSSCAL() {
	alpha := float32(2)
	X := []float32{1, 666, 2, 666}
	incX := 2
	SSCAL(alpha, X, incX)
	fmt.Println(X)

	//Output:
	//[2 666 4 666]
}

func ExampleDSCAL() {
	alpha := 2.0
	X := []float64{1, 666, 2, 666}
	incX := 2
	DSCAL(alpha, X, incX)
	fmt.Println(X)

	//Output:
	//[2 666 4 666]
}

func ExampleCSCAL() {
	alpha := complex64(complex(0, 1))
	X := []complex64{1, 666, 2, 666}
	incX := 2
	CSCAL(alpha, X, incX)
	fmt.Println(X)

	//Output:
	//[(0+1i) (666+0i) (0+2i) (666+0i)]
}

func ExampleZSCAL() {
	alpha := complex(0, 1)
	X := []complex128{1, 666, 2, 666}
	incX := 2
	ZSCAL(alpha, X, incX)
	fmt.Println(X)

	//Output:
	//[(0+1i) (666+0i) (0+2i) (666+0i)]
}

func ExampleCSSCAL() {
	alpha := float32(2)
	X := []complex64{1, 666, 2, 666}
	incX := 2
	CSSCAL(alpha, X, incX)
	fmt.Println(X)

	//Output:
	//[(2+0i) (666+0i) (4+0i) (666+0i)]
}

func ExampleZDSCAL() {
	alpha := 2.0
	X := []complex128{1, 666, 2, 666}
	incX := 2
	ZDSCAL(alpha, X, incX)
	fmt.Println(X)

	//Output:
	//[(2+0i) (666+0i) (4+0i) (666+0i)]
}

func ExampleSGEMV() {
	alpha := float32(1.0)
	A := MakeFloat32Matrix(2, 3)
	A[0][1] = 1
	A[0][2] = 2
	A[1][0] = 3

	X := []float32{-1, 4, 0}
	incX := 1
	beta := float32(0.0)
	Y := []float32{0, 0}
	incY := 1
	SGEMV(NoTrans, alpha, A, X, incX, beta, Y, incY)
	fmt.Println(A, "*", X, "=", Y)
	SGEMV(Trans, alpha, A, Y, incX, beta, X, incY)
	fmt.Println(A, "^T*", Y, "=", X)

	//Output:
	//[[0 1 2] [3 0 0]] * [-1 4 0] = [4 -3]
	//[[0 1 2] [3 0 0]] ^T* [4 -3] = [-9 4 8]
}

/*
func ExampleSTRMV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeFloat32Matrix(2, 2)
	X := []float32{0, 0, 0}
	result := STRMV(Uplo, NoTrans, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSTRSV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeFloat32Matrix(2, 2)
	X := []float32{0, 0, 0}
	result := STRSV(Uplo, NoTrans, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDGEMV() {

	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	X := []float64{0, 0, 0}
	beta := float64(2.0)
	Y := []float64{0, 0, 0}
	result := DGEMV(NoTrans, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDTRMV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeFloat64Matrix(2, 2)
	X := []float64{0, 0, 0}
	result := DTRMV(Uplo, NoTrans, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDTRSV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeFloat64Matrix(2, 2)
	X := []float64{0, 0, 0}
	result := DTRSV(Uplo, NoTrans, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCGEMV() {

	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	X := []complex64{0, 0, 0}
	beta := complex64(complex(2.0, 3.0))
	Y := []complex64{0, 0, 0}
	result := CGEMV(NoTrans, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCTRMV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeComplex64Matrix(2, 2)
	X := []complex64{0, 0, 0}
	result := CTRMV(Uplo, NoTrans, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCTRSV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeComplex64Matrix(2, 2)
	X := []complex64{0, 0, 0}
	result := CTRSV(Uplo, NoTrans, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZGEMV() {

	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	X := []complex128{0, 0, 0}
	beta := complex128(complex(2.0, 3.0))
	Y := []complex128{0, 0, 0}
	result := ZGEMV(NoTrans, alpha, A, X, beta, Y)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZTRMV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeComplex128Matrix(2, 2)
	X := []complex128{0, 0, 0}
	result := ZTRMV(Uplo, NoTrans, Diag, A, X)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZTRSV() {
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	A := MakeComplex128Matrix(2, 2)
	X := []complex128{0, 0, 0}
	result := ZTRSV(Uplo, NoTrans, Diag, A, X)
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

	TransB := Transpose(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	beta := float32(2.0)
	C := MakeFloat32Matrix(2, 2)
	result := SGEMM(NoTrans, TransB, alpha, A, B, beta, C)
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

	Diag := Diag(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	result := STRMM(Side, Uplo, NoTrans, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleSTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	alpha := float32(2.0)
	A := MakeFloat32Matrix(2, 2)
	B := MakeFloat32Matrix(2, 2)
	result := STRSM(Side, Uplo, NoTrans, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDGEMM() {

	TransB := Transpose(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	beta := float64(2.0)
	C := MakeFloat64Matrix(2, 2)
	result := DGEMM(NoTrans, TransB, alpha, A, B, beta, C)
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

	Diag := Diag(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	result := DTRMM(Side, Uplo, NoTrans, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleDTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	alpha := float64(2.0)
	A := MakeFloat64Matrix(2, 2)
	B := MakeFloat64Matrix(2, 2)
	result := DTRSM(Side, Uplo, NoTrans, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCGEMM() {

	TransB := Transpose(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := complex64(complex(2.0, 3.0))
	C := MakeComplex64Matrix(2, 2)
	result := CGEMM(NoTrans, TransB, alpha, A, B, beta, C)
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

	Diag := Diag(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	result := CTRMM(Side, Uplo, NoTrans, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	result := CTRSM(Side, Uplo, NoTrans, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZGEMM() {

	TransB := Transpose(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := complex128(complex(2.0, 3.0))
	C := MakeComplex128Matrix(2, 2)
	result := ZGEMM(NoTrans, TransB, alpha, A, B, beta, C)
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

	Diag := Diag(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	result := ZTRMM(Side, Uplo, NoTrans, Diag, alpha, A, B)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZTRSM() {
	Side := Side(2.0)
	Uplo := Uplo(2.0)

	Diag := Diag(2.0)
	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	result := ZTRSM(Side, Uplo, NoTrans, Diag, alpha, A, B)
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
