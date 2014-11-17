package blas

//THIS FILE IS AUTO-GENERATED, EDITING IS FUTILE.

import (
	. "github.com/barnex/matrix"
	"testing"
)

const N = 1024 // size of all vectors and matrices

func BenchmarkCAXPY(b *testing.B) {
	b.StopTimer()

	arg0 := complex64(1)
	arg1 := make([]complex64, N)
	arg2 := 1
	arg3 := make([]complex64, N)
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CAXPY(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkCCOPY(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex64, N)
	arg1 := 1
	arg2 := make([]complex64, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CCOPY(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkCDOTC(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex64, N)
	arg1 := 1
	arg2 := make([]complex64, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CDOTC(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkCDOTU(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex64, N)
	arg1 := 1
	arg2 := make([]complex64, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CDOTU(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkCGEMM(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := NoTrans
	arg2 := complex64(1)
	arg3 := MakeComplex64(N, N)
	arg4 := MakeComplex64(N, N)
	arg5 := complex64(1)
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CGEMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCGEMV(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := complex64(1)
	arg2 := MakeComplex64(N, N)
	arg3 := make([]complex64, N)
	arg4 := 1
	arg5 := complex64(1)
	arg6 := make([]complex64, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CGEMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkCGERC(b *testing.B) {
	b.StopTimer()

	arg0 := complex64(1)
	arg1 := make([]complex64, N)
	arg2 := 1
	arg3 := make([]complex64, N)
	arg4 := 1
	arg5 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CGERC(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkCGERU(b *testing.B) {
	b.StopTimer()

	arg0 := complex64(1)
	arg1 := make([]complex64, N)
	arg2 := 1
	arg3 := make([]complex64, N)
	arg4 := 1
	arg5 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CGERU(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkCHEMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := complex64(1)
	arg3 := MakeComplex64(N, N)
	arg4 := MakeComplex64(N, N)
	arg5 := complex64(1)
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CHEMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCHEMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := complex64(1)
	arg2 := MakeComplex64(N, N)
	arg3 := make([]complex64, N)
	arg4 := 1
	arg5 := complex64(1)
	arg6 := make([]complex64, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CHEMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkCHER(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float32(1)
	arg2 := make([]complex64, N)
	arg3 := 1
	arg4 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CHER(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkCHER2(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := complex64(1)
	arg2 := make([]complex64, N)
	arg3 := 1
	arg4 := make([]complex64, N)
	arg5 := 1
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CHER2(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCHER2K(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := complex64(1)
	arg3 := MakeComplex64(N, N)
	arg4 := MakeComplex64(N, N)
	arg5 := float32(1)
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CHER2K(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCHERK(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := float32(1)
	arg3 := MakeComplex64(N, N)
	arg4 := float32(1)
	arg5 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CHERK(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkCSCAL(b *testing.B) {
	b.StopTimer()

	arg0 := complex64(1)
	arg1 := make([]complex64, N)
	arg2 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSCAL(arg0, arg1, arg2)
	}
}

func BenchmarkCSSCAL(b *testing.B) {
	b.StopTimer()

	arg0 := float32(1)
	arg1 := make([]complex64, N)
	arg2 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSSCAL(arg0, arg1, arg2)
	}
}

func BenchmarkCSWAP(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex64, N)
	arg1 := 1
	arg2 := make([]complex64, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSWAP(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkCSYMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := complex64(1)
	arg3 := MakeComplex64(N, N)
	arg4 := MakeComplex64(N, N)
	arg5 := complex64(1)
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSYMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCSYR2K(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := complex64(1)
	arg3 := MakeComplex64(N, N)
	arg4 := MakeComplex64(N, N)
	arg5 := complex64(1)
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSYR2K(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCSYRK(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := complex64(1)
	arg3 := MakeComplex64(N, N)
	arg4 := complex64(1)
	arg5 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSYRK(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkCSize(b *testing.B) {
	b.StopTimer()

	arg0 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSize(arg0)
	}
}

func BenchmarkCSubmatrix(b *testing.B) {
	b.StopTimer()

	arg0 := MakeComplex64(N, N)
	arg1 := 1
	arg2 := 1
	arg3 := 1
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSubmatrix(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkCTRMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := complex64(1)
	arg5 := MakeComplex64(N, N)
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CTRMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCTRMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeComplex64(N, N)
	arg4 := make([]complex64, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CTRMV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkCTRSM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := complex64(1)
	arg5 := MakeComplex64(N, N)
	arg6 := MakeComplex64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CTRSM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkCTRSV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeComplex64(N, N)
	arg4 := make([]complex64, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CTRSV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkDASUM(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float64, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DASUM(arg0, arg1)
	}
}

func BenchmarkDAXPY(b *testing.B) {
	b.StopTimer()

	arg0 := float64(1)
	arg1 := make([]float64, N)
	arg2 := 1
	arg3 := make([]float64, N)
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DAXPY(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkDCOPY(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float64, N)
	arg1 := 1
	arg2 := make([]float64, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DCOPY(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkDDOT(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float64, N)
	arg1 := 1
	arg2 := make([]float64, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DDOT(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkDGEMM(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := NoTrans
	arg2 := float64(1)
	arg3 := MakeFloat64(N, N)
	arg4 := MakeFloat64(N, N)
	arg5 := float64(1)
	arg6 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DGEMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkDGEMV(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := float64(1)
	arg2 := MakeFloat64(N, N)
	arg3 := make([]float64, N)
	arg4 := 1
	arg5 := float64(1)
	arg6 := make([]float64, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DGEMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkDGER(b *testing.B) {
	b.StopTimer()

	arg0 := float64(1)
	arg1 := make([]float64, N)
	arg2 := 1
	arg3 := make([]float64, N)
	arg4 := 1
	arg5 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DGER(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkDNRM2(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float64, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DNRM2(arg0, arg1)
	}
}

func BenchmarkDSCAL(b *testing.B) {
	b.StopTimer()

	arg0 := float64(1)
	arg1 := make([]float64, N)
	arg2 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSCAL(arg0, arg1, arg2)
	}
}

func BenchmarkDSDOT(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float32, N)
	arg1 := 1
	arg2 := make([]float32, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSDOT(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkDSWAP(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float64, N)
	arg1 := 1
	arg2 := make([]float64, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSWAP(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkDSYMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := float64(1)
	arg3 := MakeFloat64(N, N)
	arg4 := MakeFloat64(N, N)
	arg5 := float64(1)
	arg6 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSYMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkDSYMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float64(1)
	arg2 := MakeFloat64(N, N)
	arg3 := make([]float64, N)
	arg4 := 1
	arg5 := float64(1)
	arg6 := make([]float64, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSYMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkDSYR(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float64(1)
	arg2 := make([]float64, N)
	arg3 := 1
	arg4 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSYR(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkDSYR2(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float64(1)
	arg2 := make([]float64, N)
	arg3 := 1
	arg4 := make([]float64, N)
	arg5 := 1
	arg6 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSYR2(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkDSYR2K(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := float64(1)
	arg3 := MakeFloat64(N, N)
	arg4 := MakeFloat64(N, N)
	arg5 := float64(1)
	arg6 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSYR2K(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkDSYRK(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := float64(1)
	arg3 := MakeFloat64(N, N)
	arg4 := float64(1)
	arg5 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSYRK(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkDSize(b *testing.B) {
	b.StopTimer()

	arg0 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSize(arg0)
	}
}

func BenchmarkDSubmatrix(b *testing.B) {
	b.StopTimer()

	arg0 := MakeFloat64(N, N)
	arg1 := 1
	arg2 := 1
	arg3 := 1
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DSubmatrix(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkDTRMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := float64(1)
	arg5 := MakeFloat64(N, N)
	arg6 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DTRMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkDTRMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeFloat64(N, N)
	arg4 := make([]float64, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DTRMV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkDTRSM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := float64(1)
	arg5 := MakeFloat64(N, N)
	arg6 := MakeFloat64(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DTRSM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkDTRSV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeFloat64(N, N)
	arg4 := make([]float64, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DTRSV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkDZASUM(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex128, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DZASUM(arg0, arg1)
	}
}

func BenchmarkDZNRM2(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex128, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		DZNRM2(arg0, arg1)
	}
}

func BenchmarkICAMAX(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex64, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ICAMAX(arg0, arg1)
	}
}

func BenchmarkIDAMAX(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float64, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IDAMAX(arg0, arg1)
	}
}

func BenchmarkISAMAX(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float32, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ISAMAX(arg0, arg1)
	}
}

func BenchmarkIZAMAX(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex128, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IZAMAX(arg0, arg1)
	}
}

func BenchmarkMakeComplex128(b *testing.B) {
	b.StopTimer()

	arg0 := 1
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		MakeComplex128(arg0, arg1)
	}
}

func BenchmarkMakeComplex64(b *testing.B) {
	b.StopTimer()

	arg0 := 1
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		MakeComplex64(arg0, arg1)
	}
}

func BenchmarkMakeFloat32(b *testing.B) {
	b.StopTimer()

	arg0 := 1
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		MakeFloat32(arg0, arg1)
	}
}

func BenchmarkMakeFloat64(b *testing.B) {
	b.StopTimer()

	arg0 := 1
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		MakeFloat64(arg0, arg1)
	}
}

func BenchmarkSASUM(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float32, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SASUM(arg0, arg1)
	}
}

func BenchmarkSAXPY(b *testing.B) {
	b.StopTimer()

	arg0 := float32(1)
	arg1 := make([]float32, N)
	arg2 := 1
	arg3 := make([]float32, N)
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SAXPY(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkSCASUM(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex64, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SCASUM(arg0, arg1)
	}
}

func BenchmarkSCNRM2(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex64, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SCNRM2(arg0, arg1)
	}
}

func BenchmarkSCOPY(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float32, N)
	arg1 := 1
	arg2 := make([]float32, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SCOPY(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkSDOT(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float32, N)
	arg1 := 1
	arg2 := make([]float32, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SDOT(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkSDSDOT(b *testing.B) {
	b.StopTimer()

	arg0 := float32(1)
	arg1 := make([]float32, N)
	arg2 := 1
	arg3 := make([]float32, N)
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SDSDOT(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkSGEMM(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := NoTrans
	arg2 := float32(1)
	arg3 := MakeFloat32(N, N)
	arg4 := MakeFloat32(N, N)
	arg5 := float32(1)
	arg6 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SGEMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkSGEMV(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := float32(1)
	arg2 := MakeFloat32(N, N)
	arg3 := make([]float32, N)
	arg4 := 1
	arg5 := float32(1)
	arg6 := make([]float32, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SGEMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkSGER(b *testing.B) {
	b.StopTimer()

	arg0 := float32(1)
	arg1 := make([]float32, N)
	arg2 := 1
	arg3 := make([]float32, N)
	arg4 := 1
	arg5 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SGER(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkSNRM2(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float32, N)
	arg1 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SNRM2(arg0, arg1)
	}
}

func BenchmarkSSCAL(b *testing.B) {
	b.StopTimer()

	arg0 := float32(1)
	arg1 := make([]float32, N)
	arg2 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSCAL(arg0, arg1, arg2)
	}
}

func BenchmarkSSWAP(b *testing.B) {
	b.StopTimer()

	arg0 := make([]float32, N)
	arg1 := 1
	arg2 := make([]float32, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSWAP(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkSSYMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := float32(1)
	arg3 := MakeFloat32(N, N)
	arg4 := MakeFloat32(N, N)
	arg5 := float32(1)
	arg6 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSYMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkSSYMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float32(1)
	arg2 := MakeFloat32(N, N)
	arg3 := make([]float32, N)
	arg4 := 1
	arg5 := float32(1)
	arg6 := make([]float32, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSYMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkSSYR(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float32(1)
	arg2 := make([]float32, N)
	arg3 := 1
	arg4 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSYR(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkSSYR2(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float32(1)
	arg2 := make([]float32, N)
	arg3 := 1
	arg4 := make([]float32, N)
	arg5 := 1
	arg6 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSYR2(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkSSYR2K(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := float32(1)
	arg3 := MakeFloat32(N, N)
	arg4 := MakeFloat32(N, N)
	arg5 := float32(1)
	arg6 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSYR2K(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkSSYRK(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := float32(1)
	arg3 := MakeFloat32(N, N)
	arg4 := float32(1)
	arg5 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSYRK(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkSSize(b *testing.B) {
	b.StopTimer()

	arg0 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSize(arg0)
	}
}

func BenchmarkSSubmatrix(b *testing.B) {
	b.StopTimer()

	arg0 := MakeFloat32(N, N)
	arg1 := 1
	arg2 := 1
	arg3 := 1
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SSubmatrix(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkSTRMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := float32(1)
	arg5 := MakeFloat32(N, N)
	arg6 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		STRMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkSTRMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeFloat32(N, N)
	arg4 := make([]float32, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		STRMV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkSTRSM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := float32(1)
	arg5 := MakeFloat32(N, N)
	arg6 := MakeFloat32(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		STRSM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkSTRSV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeFloat32(N, N)
	arg4 := make([]float32, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		STRSV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkZAXPY(b *testing.B) {
	b.StopTimer()

	arg0 := complex128(1)
	arg1 := make([]complex128, N)
	arg2 := 1
	arg3 := make([]complex128, N)
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZAXPY(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkZCOPY(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex128, N)
	arg1 := 1
	arg2 := make([]complex128, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZCOPY(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkZDOTC(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex128, N)
	arg1 := 1
	arg2 := make([]complex128, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZDOTC(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkZDOTU(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex128, N)
	arg1 := 1
	arg2 := make([]complex128, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZDOTU(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkZDSCAL(b *testing.B) {
	b.StopTimer()

	arg0 := float64(1)
	arg1 := make([]complex128, N)
	arg2 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZDSCAL(arg0, arg1, arg2)
	}
}

func BenchmarkZGEMM(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := NoTrans
	arg2 := complex128(1)
	arg3 := MakeComplex128(N, N)
	arg4 := MakeComplex128(N, N)
	arg5 := complex128(1)
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZGEMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZGEMV(b *testing.B) {
	b.StopTimer()

	arg0 := NoTrans
	arg1 := complex128(1)
	arg2 := MakeComplex128(N, N)
	arg3 := make([]complex128, N)
	arg4 := 1
	arg5 := complex128(1)
	arg6 := make([]complex128, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZGEMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkZGERC(b *testing.B) {
	b.StopTimer()

	arg0 := complex128(1)
	arg1 := make([]complex128, N)
	arg2 := 1
	arg3 := make([]complex128, N)
	arg4 := 1
	arg5 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZGERC(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkZGERU(b *testing.B) {
	b.StopTimer()

	arg0 := complex128(1)
	arg1 := make([]complex128, N)
	arg2 := 1
	arg3 := make([]complex128, N)
	arg4 := 1
	arg5 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZGERU(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkZHEMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := complex128(1)
	arg3 := MakeComplex128(N, N)
	arg4 := MakeComplex128(N, N)
	arg5 := complex128(1)
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZHEMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZHEMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := complex128(1)
	arg2 := MakeComplex128(N, N)
	arg3 := make([]complex128, N)
	arg4 := 1
	arg5 := complex128(1)
	arg6 := make([]complex128, N)
	arg7 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZHEMV(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func BenchmarkZHER(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := float64(1)
	arg2 := make([]complex128, N)
	arg3 := 1
	arg4 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZHER(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkZHER2(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := complex128(1)
	arg2 := make([]complex128, N)
	arg3 := 1
	arg4 := make([]complex128, N)
	arg5 := 1
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZHER2(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZHER2K(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := complex128(1)
	arg3 := MakeComplex128(N, N)
	arg4 := MakeComplex128(N, N)
	arg5 := float64(1)
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZHER2K(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZHERK(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := float64(1)
	arg3 := MakeComplex128(N, N)
	arg4 := float64(1)
	arg5 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZHERK(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkZSCAL(b *testing.B) {
	b.StopTimer()

	arg0 := complex128(1)
	arg1 := make([]complex128, N)
	arg2 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZSCAL(arg0, arg1, arg2)
	}
}

func BenchmarkZSWAP(b *testing.B) {
	b.StopTimer()

	arg0 := make([]complex128, N)
	arg1 := 1
	arg2 := make([]complex128, N)
	arg3 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZSWAP(arg0, arg1, arg2, arg3)
	}
}

func BenchmarkZSYMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := complex128(1)
	arg3 := MakeComplex128(N, N)
	arg4 := MakeComplex128(N, N)
	arg5 := complex128(1)
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZSYMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZSYR2K(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := complex128(1)
	arg3 := MakeComplex128(N, N)
	arg4 := MakeComplex128(N, N)
	arg5 := complex128(1)
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZSYR2K(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZSYRK(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := complex128(1)
	arg3 := MakeComplex128(N, N)
	arg4 := complex128(1)
	arg5 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZSYRK(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkZSize(b *testing.B) {
	b.StopTimer()

	arg0 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZSize(arg0)
	}
}

func BenchmarkZSubmatrix(b *testing.B) {
	b.StopTimer()

	arg0 := MakeComplex128(N, N)
	arg1 := 1
	arg2 := 1
	arg3 := 1
	arg4 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZSubmatrix(arg0, arg1, arg2, arg3, arg4)
	}
}

func BenchmarkZTRMM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := complex128(1)
	arg5 := MakeComplex128(N, N)
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZTRMM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZTRMV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeComplex128(N, N)
	arg4 := make([]complex128, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZTRMV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

func BenchmarkZTRSM(b *testing.B) {
	b.StopTimer()

	arg0 := Left
	arg1 := Upper
	arg2 := NoTrans
	arg3 := NonUnit
	arg4 := complex128(1)
	arg5 := MakeComplex128(N, N)
	arg6 := MakeComplex128(N, N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZTRSM(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

func BenchmarkZTRSV(b *testing.B) {
	b.StopTimer()

	arg0 := Upper
	arg1 := NoTrans
	arg2 := NonUnit
	arg3 := MakeComplex128(N, N)
	arg4 := make([]complex128, N)
	arg5 := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ZTRSV(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}
