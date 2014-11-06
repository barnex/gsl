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
	A := MakeFloat32Matrix(2, 3)
	A[0][1] = 1
	A[0][2] = 2
	A[1][0] = 3

	X := []float32{-1, 4, 0}
	incX := 1
	Y := []float32{0, 0}
	incY := 1

	alpha := float32(1.0)
	beta := float32(0.0)

	SGEMV(NoTrans, alpha, A, X, incX, beta, Y, incY)
	fmt.Println(A, "*", X, "=", Y)

	SGEMV(Trans, alpha, A, Y, incX, beta, X, incY)
	fmt.Println(A, "^T*", Y, "=", X)

	//Output:
	//[[0 1 2] [3 0 0]] * [-1 4 0] = [4 -3]
	//[[0 1 2] [3 0 0]] ^T* [4 -3] = [-9 4 8]
}

func ExampleSTRMV() {
	A := MakeFloat32Matrix(2, 2)
	A[0][0] = 1
	A[1][1] = 1

	A[0][1] = 2

	X := []float32{-1, 4}
	incX := 1

	fmt.Println(A, "*", X)
	STRMV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println("=", X)

	//Output:
	//[[1 2] [0 1]] * [-1 4]
	//= [7 4]
}

func ExampleSTRSV() {
	// solve:
	//  x - y = 2
	//      y = 3
	// for x, y:
	A := MakeFloat32Matrix(2, 2)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 1
	X := []float32{2, 3}
	incX := 1
	STRSV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println(X)

	//Output:
	//[5 3]
}

func ExampleDGEMV() {
	A := MakeFloat64Matrix(2, 3)
	A[0][1] = 1
	A[0][2] = 2
	A[1][0] = 3

	X := []float64{-1, 4, 0}
	incX := 1
	Y := []float64{0, 0}
	incY := 1

	alpha := 1.0
	beta := 0.0

	DGEMV(NoTrans, alpha, A, X, incX, beta, Y, incY)
	fmt.Println(A, "*", X, "=", Y)

	DGEMV(Trans, alpha, A, Y, incX, beta, X, incY)
	fmt.Println(A, "^T*", Y, "=", X)

	//Output:
	//[[0 1 2] [3 0 0]] * [-1 4 0] = [4 -3]
	//[[0 1 2] [3 0 0]] ^T* [4 -3] = [-9 4 8]
}

func ExampleDTRMV() {
	A := MakeFloat64Matrix(2, 2)
	A[0][0] = 1
	A[1][1] = 1

	A[0][1] = 2

	X := []float64{-1, 4}
	incX := 1

	fmt.Println(A, "*", X)
	DTRMV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println("=", X)

	//Output:
	//[[1 2] [0 1]] * [-1 4]
	//= [7 4]
}

func ExampleDTRSV() {
	// solve:
	//  x - y = 2
	//      y = 3
	// for x, y:
	A := MakeFloat64Matrix(2, 2)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 1
	X := []float64{2, 3}
	incX := 1
	DTRSV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println(X)

	//Output:
	//[5 3]
}

func ExampleCGEMV() {
	A := MakeComplex64Matrix(2, 3)
	A[0][1] = 1
	A[0][2] = 2
	A[1][0] = 3

	X := []complex64{-1, 4, 0}
	incX := 1
	Y := []complex64{0, 0}
	incY := 1

	alpha := complex64(complex(1, 0))
	beta := complex64(complex(0, 0))

	CGEMV(NoTrans, alpha, A, X, incX, beta, Y, incY)
	fmt.Println(A, "*", X, "=", Y)

	CGEMV(Trans, alpha, A, Y, incX, beta, X, incY)
	fmt.Println(A, "^T*", Y, "=", X)

	//Output:
	//[[(0+0i) (1+0i) (2+0i)] [(3+0i) (0+0i) (0+0i)]] * [(-1+0i) (4+0i) (0+0i)] = [(4+0i) (-3+0i)]
	//[[(0+0i) (1+0i) (2+0i)] [(3+0i) (0+0i) (0+0i)]] ^T* [(4+0i) (-3+0i)] = [(-9+0i) (4+0i) (8+0i)]
}

func ExampleCTRMV() {
	A := MakeComplex64Matrix(2, 2)
	A[0][0] = 1
	A[1][1] = 1

	A[0][1] = 2

	X := []complex64{-1, 4}
	incX := 1

	fmt.Println(A, "*", X)
	CTRMV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println("=", X)

	//Output:
	//[[(1+0i) (2+0i)] [(0+0i) (1+0i)]] * [(-1+0i) (4+0i)]
	//= [(7+0i) (4+0i)]
}

func ExampleCTRSV() {
	// solve:
	//  x - y = 2
	//      y = 3
	// for x, y:
	A := MakeComplex64Matrix(2, 2)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 1
	X := []complex64{2, 3}
	incX := 1
	CTRSV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println(X)

	//Output:
	//[(5+0i) (3+0i)]
}

func ExampleZTRSV() {
	// solve:
	//  x - y = 2
	//      y = 3
	// for x, y:
	A := MakeComplex128Matrix(2, 2)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 1
	X := []complex128{2, 3}
	incX := 1
	ZTRSV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println(X)

	//Output:
	//[(5+0i) (3+0i)]
}

func ExampleZGEMV() {
	A := MakeComplex128Matrix(2, 3)
	A[0][1] = 1
	A[0][2] = 2
	A[1][0] = 3

	X := []complex128{-1, 4, 0}
	incX := 1
	Y := []complex128{0, 0}
	incY := 1

	alpha := complex(1, 0)
	beta := complex(0, 0)

	ZGEMV(NoTrans, alpha, A, X, incX, beta, Y, incY)
	fmt.Println(A, "*", X, "=", Y)

	ZGEMV(Trans, alpha, A, Y, incX, beta, X, incY)
	fmt.Println(A, "^T*", Y, "=", X)

	//Output:
	//[[(0+0i) (1+0i) (2+0i)] [(3+0i) (0+0i) (0+0i)]] * [(-1+0i) (4+0i) (0+0i)] = [(4+0i) (-3+0i)]
	//[[(0+0i) (1+0i) (2+0i)] [(3+0i) (0+0i) (0+0i)]] ^T* [(4+0i) (-3+0i)] = [(-9+0i) (4+0i) (8+0i)]
}

func ExampleZTRMV() {
	A := MakeComplex128Matrix(2, 2)
	A[0][0] = 1
	A[1][1] = 1

	A[0][1] = 2

	X := []complex128{-1, 4}
	incX := 1

	fmt.Println(A, "*", X)
	ZTRMV(Upper, NoTrans, NonUnit, A, X, incX)
	fmt.Println("=", X)

	//Output:
	//[[(1+0i) (2+0i)] [(0+0i) (1+0i)]] * [(-1+0i) (4+0i)]
	//= [(7+0i) (4+0i)]
}

func ExampleSSYMV() {
	A := MakeFloat32Matrix(2, 2)
	A[0][0] = 1
	A[1][1] = 2
	A[0][1] = 3
	A[1][0] = 3

	X := []float32{-1, 4}
	incX := 1
	Y := []float32{0, 0}
	incY := 1

	alpha := float32(1.0)
	beta := float32(0.0)

	SSYMV(Upper, alpha, A, X, incX, beta, Y, incY)
	fmt.Println(A, "*", X, "=", Y)

	//Output:
	//[[1 3] [3 2]] * [-1 4] = [11 5]
}

func ExampleSGER() {
	alpha := float32(1)
	X := []float32{1, 2}
	incX := 1
	Y := []float32{3, 4, 5}
	incY := 1
	A := MakeFloat32Matrix(2, 3)
	SGER(alpha, X, incX, Y, incY, A)
	fmt.Println(X, "*", Y, "^T = ", A)

	//Output:
	//[1 2] * [3 4 5] ^T =  [[3 4 5] [6 8 10]]
}

func ExampleSSYR() {
	alpha := float32(1)
	X := []float32{1, 2}
	incX := 1
	A := MakeFloat32Matrix(2, 2)
	SSYR(Upper, alpha, X, incX, A)
	fmt.Println(X, "*", X, "^T = ", A)

	//Output:
	//[1 2] * [1 2] ^T =  [[1 2] [0 4]]
}

/*
func ExampleSSYR2() {

	alpha := float32(2.0)
	X := []float32{0, 0, 0}
	Y := []float32{0, 0, 0}
	A := MakeFloat32Matrix(2, 2)
	result := SSYR2(Uplo, alpha, X, Y, A)
	fmt.Println(result)

	//Output:
	//
}
*/

func ExampleDSYMV() {
	A := MakeFloat64Matrix(2, 2)
	A[0][0] = 1
	A[1][1] = 2
	A[0][1] = 3
	A[1][0] = 3

	X := []float64{-1, 4}
	incX := 1
	Y := []float64{0, 0}
	incY := 1

	alpha := float64(1.0)
	beta := float64(0.0)

	DSYMV(Upper, alpha, A, X, incX, beta, Y, incY)
	fmt.Println(A, "*", X, "=", Y)

	//Output:
	//[[1 3] [3 2]] * [-1 4] = [11 5]
}

func ExampleDGER() {
	alpha := 1.0
	X := []float64{1, 2}
	incX := 1
	Y := []float64{3, 4, 5}
	incY := 1
	A := MakeFloat64Matrix(2, 3)
	DGER(alpha, X, incX, Y, incY, A)
	fmt.Println(X, "*", Y, "^T = ", A)

	//Output:
	//[1 2] * [3 4 5] ^T =  [[3 4 5] [6 8 10]]
}

func ExampleDSYR() {
	alpha := 1.0
	X := []float64{1, 2}
	incX := 1
	A := MakeFloat64Matrix(2, 2)
	DSYR(Upper, alpha, X, incX, A)
	fmt.Println(X, "*", X, "^T = ", A)

	//Output:
	//[1 2] * [1 2] ^T =  [[1 2] [0 4]]
}

/*

func ExampleDSYR2() {

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


*/
func ExampleCGERU() {
	alpha := complex64(1)
	X := []complex64{1, 2}
	incX := 1
	Y := []complex64{3, 4, 5}
	incY := 1
	A := MakeComplex64Matrix(2, 3)
	CGERU(alpha, X, incX, Y, incY, A)
	fmt.Println(X, "*", Y, "^T = ", A)

	//Output:
	//[(1+0i) (2+0i)] * [(3+0i) (4+0i) (5+0i)] ^T =  [[(3+0i) (4+0i) (5+0i)] [(6+0i) (8+0i) (10+0i)]]
}

func ExampleCGERC() {
	alpha := complex64(1)
	X := []complex64{1, 2}
	incX := 1
	Y := []complex64{complex(0, 3), complex(0, 4), complex(0, 5)}
	incY := 1
	A := MakeComplex64Matrix(2, 3)
	CGERC(alpha, X, incX, Y, incY, A)
	fmt.Println(X, "* conj", Y, "^T = ", A)

	//Output:
	//[(1+0i) (2+0i)] * conj [(0+3i) (0+4i) (0+5i)] ^T =  [[(0-3i) (0-4i) (0-5i)] [(0-6i) (0-8i) (0-10i)]]
}

func ExampleCHER() {
	alpha := float32(1.0)
	X := []complex64{complex(0, 1), 2}
	incX := 1
	A := MakeComplex64Matrix(2, 2)
	CHER(Upper, alpha, X, incX, A)
	fmt.Println(A)

	//Output:
	//[[(1+0i) (0+2i)] [(0+0i) (4+0i)]]
}

func ExampleCHER2() {
	alpha := complex64(1.0)
	X := []complex64{complex(0, 1), 2}
	incX := 1
	Y := []complex64{2, complex(3, 1)}
	incY := 1
	A := MakeComplex64Matrix(2, 2)
	CHER2(Upper, alpha, X, incX, Y, incY, A)
	fmt.Println(A)

	//Output:
	//[[(2+0i) (0+4i)] [(0+0i) (8+0i)]]
}

/*
func ExampleZHEMV() {

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


*/
func ExampleZGERU() {
	alpha := complex128(1)
	X := []complex128{1, 2}
	incX := 1
	Y := []complex128{3, 4, 5}
	incY := 1
	A := MakeComplex128Matrix(2, 3)
	ZGERU(alpha, X, incX, Y, incY, A)
	fmt.Println(X, "*", Y, "^T = ", A)

	//Output:
	//[(1+0i) (2+0i)] * [(3+0i) (4+0i) (5+0i)] ^T =  [[(3+0i) (4+0i) (5+0i)] [(6+0i) (8+0i) (10+0i)]]
}

func ExampleZGERC() {
	alpha := complex128(1)
	X := []complex128{1, 2}
	incX := 1
	Y := []complex128{complex(0, 3), complex(0, 4), complex(0, 5)}
	incY := 1
	A := MakeComplex128Matrix(2, 3)
	ZGERC(alpha, X, incX, Y, incY, A)
	fmt.Println(X, "* conj", Y, "^T = ", A)

	//Output:
	//[(1+0i) (2+0i)] * conj [(0+3i) (0+4i) (0+5i)] ^T =  [[(0-3i) (0-4i) (0-5i)] [(0-6i) (0-8i) (0-10i)]]
}

func ExampleZHER() {
	alpha := 1.0
	X := []complex128{complex(0, 1), 2}
	incX := 1
	A := MakeComplex128Matrix(2, 2)
	ZHER(Upper, alpha, X, incX, A)
	fmt.Println(A)

	//Output:
	//[[(1+0i) (0+2i)] [(0+0i) (4+0i)]]
}

func ExampleZHER2() {
	alpha := complex128(1.0)
	X := []complex128{complex(0, 1), 2}
	incX := 1
	Y := []complex128{2, complex(3, 1)}
	incY := 1
	A := MakeComplex128Matrix(2, 2)
	ZHER2(Upper, alpha, X, incX, Y, incY, A)
	fmt.Println(A)

	//Output:
	//[[(2+0i) (0+4i)] [(0+0i) (8+0i)]]
}

func ExampleSGEMM() {
	alpha := float32(1.0)
	A := MakeFloat32Matrix(2, 4)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 2

	B := MakeFloat32Matrix(4, 3)
	B[1][0] = 2
	B[2][1] = 3

	beta := float32(0.0)
	C := MakeFloat32Matrix(2, 3)
	SGEMM(NoTrans, NoTrans, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[1 -1 0 0] [0 2 0 0]] * [[0 0 0] [2 0 0] [0 3 0] [0 0 0]] = [[-2 0 0] [4 0 0]]
}

func ExampleSSYMM() {
	alpha := float32(1.0)
	A := MakeFloat32Matrix(3, 3)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 2

	B := MakeFloat32Matrix(3, 2)
	B[1][0] = 2
	B[2][1] = 3

	beta := float32(0.0)
	C := MakeFloat32Matrix(3, 2)
	SSYMM(Left, Upper, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[1 -1 0] [0 2 0] [0 0 0]] * [[0 0] [2 0] [0 3]] = [[-2 0] [4 0] [0 0]]
}

func ExampleSSYRK() {
	alpha := float32(1.0)
	A := MakeFloat32Matrix(2, 3)
	A[0][0] = 1

	beta := float32(0.0)
	C := MakeFloat32Matrix(3, 3)

	SSYRK(Upper, Trans, alpha, A, beta, C)
	fmt.Println(C)

	//Output:
	//[[1 0 0] [0 0 0] [0 0 0]]
}

func ExampleSSYR2K() {
	alpha := float32(1.0)
	A := MakeFloat32Matrix(2, 3)
	A[0][0] = 1

	B := MakeFloat32Matrix(2, 3)
	B[0][0] = 1

	beta := float32(0.0)
	C := MakeFloat32Matrix(3, 3)

	SSYR2K(Upper, Trans, alpha, A, B, beta, C)
	fmt.Println(C)

	//Output:
	//[[2 0 0] [0 0 0] [0 0 0]]
}

func ExampleSTRMM() {
	alpha := float32(1.0)
	A := MakeFloat32Matrix(3, 3)
	A[1][1] = 1
	A[2][2] = 1
	B := MakeFloat32Matrix(3, 2)
	B[1][1] = 1

	STRMM(Left, Upper, NoTrans, NonUnit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[0 0] [0 1] [0 0]]
}

func ExampleSTRSM() {
	alpha := float32(1.0)
	A := MakeFloat32Matrix(3, 3)
	B := MakeFloat32Matrix(3, 2)
	B[1][1] = 1

	STRSM(Left, Upper, NoTrans, Unit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[0 0] [0 1] [0 0]]
}

func ExampleDGEMM() {
	alpha := 1.0
	A := MakeFloat64Matrix(2, 4)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 2

	B := MakeFloat64Matrix(4, 3)
	B[1][0] = 2
	B[2][1] = 3

	beta := 0.0
	C := MakeFloat64Matrix(2, 3)
	DGEMM(NoTrans, NoTrans, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[1 -1 0 0] [0 2 0 0]] * [[0 0 0] [2 0 0] [0 3 0] [0 0 0]] = [[-2 0 0] [4 0 0]]
}

func ExampleDSYMM() {
	alpha := 1.0
	A := MakeFloat64Matrix(3, 3)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 2

	B := MakeFloat64Matrix(3, 2)
	B[1][0] = 2
	B[2][1] = 3

	beta := 0.0
	C := MakeFloat64Matrix(3, 2)
	DSYMM(Left, Upper, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[1 -1 0] [0 2 0] [0 0 0]] * [[0 0] [2 0] [0 3]] = [[-2 0] [4 0] [0 0]]
}

func ExampleDSYRK() {
	alpha := 1.0
	A := MakeFloat64Matrix(2, 3)
	A[0][0] = 1

	beta := 0.0
	C := MakeFloat64Matrix(3, 3)

	DSYRK(Upper, Trans, alpha, A, beta, C)
	fmt.Println(C)

	//Output:
	//[[1 0 0] [0 0 0] [0 0 0]]
}

func ExampleDSYR2K() {
	alpha := float64(1.0)
	A := MakeFloat64Matrix(2, 3)
	A[0][0] = 1

	B := MakeFloat64Matrix(2, 3)
	B[0][0] = 1

	beta := float64(0.0)
	C := MakeFloat64Matrix(3, 3)

	DSYR2K(Upper, Trans, alpha, A, B, beta, C)
	fmt.Println(C)

	//Output:
	//[[2 0 0] [0 0 0] [0 0 0]]
}

func ExampleDTRMM() {
	alpha := float64(1.0)
	A := MakeFloat64Matrix(3, 3)
	A[1][1] = 1
	A[1][2] = 1
	B := MakeFloat64Matrix(2, 3)
	B[1][1] = 1

	DTRMM(Right, Lower, Trans, NonUnit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[0 0 0] [0 1 0]]
}

func ExampleDTRSM() {
	alpha := float64(1.0)
	A := MakeFloat64Matrix(3, 3)
	B := MakeFloat64Matrix(3, 2)
	B[1][1] = 1

	DTRSM(Left, Upper, NoTrans, Unit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[0 0] [0 1] [0 0]]
}

func ExampleCGEMM() {
	alpha := complex64(1.0)
	A := MakeComplex64Matrix(2, 4)
	A[0][0] = 1
	A[0][1] = complex(0, -1)
	A[1][1] = 2

	B := MakeComplex64Matrix(4, 3)
	B[1][0] = 2
	B[2][1] = 3

	beta := complex64(0.0)
	C := MakeComplex64Matrix(2, 3)
	CGEMM(NoTrans, NoTrans, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[(1+0i) (0-1i) (0+0i) (0+0i)] [(0+0i) (2+0i) (0+0i) (0+0i)]] * [[(0+0i) (0+0i) (0+0i)] [(2+0i) (0+0i) (0+0i)] [(0+0i) (3+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]] = [[(0-2i) (0+0i) (0+0i)] [(4+0i) (0+0i) (0+0i)]]
}

func ExampleCSYMM() {
	alpha := complex64(1)
	A := MakeComplex64Matrix(3, 3)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 2

	B := MakeComplex64Matrix(3, 2)
	B[1][0] = 2
	B[2][1] = 3

	beta := complex64(0)
	C := MakeComplex64Matrix(3, 2)
	CSYMM(Left, Upper, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[(1+0i) (-1+0i) (0+0i)] [(0+0i) (2+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]] * [[(0+0i) (0+0i)] [(2+0i) (0+0i)] [(0+0i) (3+0i)]] = [[(-2+0i) (0+0i)] [(4+0i) (0+0i)] [(0+0i) (0+0i)]]
}

func ExampleCSYRK() {
	alpha := complex64(1)
	A := MakeComplex64Matrix(2, 3)
	A[0][0] = 1

	beta := complex64(0)
	C := MakeComplex64Matrix(3, 3)

	CSYRK(Upper, Trans, alpha, A, beta, C)
	fmt.Println(C)

	//Output:
	//[[(1+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}

func ExampleCSYR2K() {
	alpha := complex64(1.0)
	A := MakeComplex64Matrix(2, 3)
	A[0][0] = 1

	B := MakeComplex64Matrix(2, 3)
	B[0][0] = 1

	beta := complex64(0.0)
	C := MakeComplex64Matrix(3, 3)

	CSYR2K(Upper, Trans, alpha, A, B, beta, C)
	fmt.Println(C)

	//Output:
	//[[(2+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}

func ExampleCTRMM() {
	alpha := complex64(1)
	A := MakeComplex64Matrix(3, 3)
	A[1][1] = 1
	A[1][2] = complex(0, 1)
	B := MakeComplex64Matrix(2, 3)
	B[1][2] = 1

	CTRMM(Right, Upper, ConjTrans, NonUnit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0-1i) (0+0i)]]
}

func ExampleCTRSM() {
	alpha := complex64(1)
	A := MakeComplex64Matrix(3, 3)
	A[1][1] = 1
	A[1][2] = complex(0, 1)
	B := MakeComplex64Matrix(2, 3)
	B[1][2] = 1

	CTRSM(Right, Upper, ConjTrans, Unit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0+1i) (1+0i)]]
}

func ExampleZGEMM() {
	alpha := complex128(1.0)
	A := MakeComplex128Matrix(2, 4)
	A[0][0] = 1
	A[0][1] = complex(0, -1)
	A[1][1] = 2

	B := MakeComplex128Matrix(4, 3)
	B[1][0] = 2
	B[2][1] = 3

	beta := complex128(0.0)
	C := MakeComplex128Matrix(2, 3)
	ZGEMM(NoTrans, NoTrans, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[(1+0i) (0-1i) (0+0i) (0+0i)] [(0+0i) (2+0i) (0+0i) (0+0i)]] * [[(0+0i) (0+0i) (0+0i)] [(2+0i) (0+0i) (0+0i)] [(0+0i) (3+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]] = [[(0-2i) (0+0i) (0+0i)] [(4+0i) (0+0i) (0+0i)]]
}

func ExampleZSYMM() {
	alpha := complex128(1)
	A := MakeComplex128Matrix(3, 3)
	A[0][0] = 1
	A[0][1] = -1
	A[1][1] = 2

	B := MakeComplex128Matrix(3, 2)
	B[1][0] = 2
	B[2][1] = 3

	beta := complex128(0)
	C := MakeComplex128Matrix(3, 2)
	ZSYMM(Left, Upper, alpha, A, B, beta, C)
	fmt.Println(A, "*", B, "=", C)

	//Output:
	//[[(1+0i) (-1+0i) (0+0i)] [(0+0i) (2+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]] * [[(0+0i) (0+0i)] [(2+0i) (0+0i)] [(0+0i) (3+0i)]] = [[(-2+0i) (0+0i)] [(4+0i) (0+0i)] [(0+0i) (0+0i)]]
}

func ExampleZSYRK() {
	alpha := complex128(1)
	A := MakeComplex128Matrix(2, 3)
	A[0][0] = 1

	beta := complex128(0)
	C := MakeComplex128Matrix(3, 3)

	ZSYRK(Upper, Trans, alpha, A, beta, C)
	fmt.Println(C)

	//Output:
	//[[(1+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}

func ExampleZSYR2K() {
	alpha := complex128(1.0)
	A := MakeComplex128Matrix(2, 3)
	A[0][0] = 1

	B := MakeComplex128Matrix(2, 3)
	B[0][0] = 1

	beta := complex128(0.0)
	C := MakeComplex128Matrix(3, 3)

	ZSYR2K(Upper, Trans, alpha, A, B, beta, C)
	fmt.Println(C)

	//Output:
	//[[(2+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}

func ExampleZTRMM() {
	alpha := complex128(1)
	A := MakeComplex128Matrix(3, 3)
	A[1][1] = 1
	A[1][2] = complex(0, 1)
	B := MakeComplex128Matrix(2, 3)
	B[1][2] = 1

	ZTRMM(Right, Upper, ConjTrans, NonUnit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0-1i) (0+0i)]]
}

func ExampleZTRSM() {
	alpha := complex128(1)
	A := MakeComplex128Matrix(3, 3)
	A[1][1] = 1
	A[1][2] = complex(0, 1)
	B := MakeComplex128Matrix(2, 3)
	B[1][2] = 1

	ZTRSM(Right, Upper, ConjTrans, Unit, alpha, A, B)
	fmt.Println(B)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0+1i) (1+0i)]]
}

/*
func ExampleCHEMM() {
	Side := Side(2.0)

	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := complex64(complex(2.0, 3.0))
	C := MakeComplex64Matrix(2, 2)
	CHEMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHERK() {


	alpha := float32(2.0)
	A := MakeComplex64Matrix(2, 2)
	beta := float32(2.0)
	C := MakeComplex64Matrix(2, 2)
	CHERK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleCHER2K() {


	alpha := complex64(complex(2.0, 3.0))
	A := MakeComplex64Matrix(2, 2)
	B := MakeComplex64Matrix(2, 2)
	beta := float32(2.0)
	C := MakeComplex64Matrix(2, 2)
	CHER2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHEMM() {
	Side := Side(2.0)

	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := complex128(complex(2.0, 3.0))
	C := MakeComplex128Matrix(2, 2)
	ZHEMM(Side, Uplo, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHERK() {


	alpha := float64(2.0)
	A := MakeComplex128Matrix(2, 2)
	beta := float64(2.0)
	C := MakeComplex128Matrix(2, 2)
	ZHERK(Uplo, Trans, alpha, A, beta, C)
	fmt.Println(result)

	//Output:
	//
}

func ExampleZHER2K() {


	alpha := complex128(complex(2.0, 3.0))
	A := MakeComplex128Matrix(2, 2)
	B := MakeComplex128Matrix(2, 2)
	beta := float64(2.0)
	C := MakeComplex128Matrix(2, 2)
	ZHER2K(Uplo, Trans, alpha, A, B, beta, C)
	fmt.Println(result)

	//Output:
	//
}
*/
