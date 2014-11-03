package blas

import (
	"unsafe"
)

// Makes a matrix with contiguous underlying storage, usable by blas functions.
func MakeFloat32Matrix(rows, cols int) [][]float32 {
	checkSize(rows, cols)
	a := make([]float32, rows*cols)
	return reshapeR2(a, [2]int{rows, cols})
}

// Makes a matrix with contiguous underlying storage, usable by blas functions.
func MakeFloat64Matrix(rows, cols int) [][]float64 {
	checkSize(rows, cols)
	a := make([]float64, rows*cols)
	return reshapeD2(a, [2]int{rows, cols})
}

// Makes a matrix with contiguous underlying storage, usable by blas functions.
func MakeComplex64Matrix(rows, cols int) [][]complex64 {
	checkSize(rows, cols)
	a := make([]complex64, rows*cols)
	return reshapeC2(a, [2]int{rows, cols})
}

// Makes a matrix with contiguous underlying storage, usable by blas functions.
func MakeComplex128Matrix(rows, cols int) [][]complex128 {
	checkSize(rows, cols)
	a := make([]complex128, rows*cols)
	return reshapeZ2(a, [2]int{rows, cols})
}

// Returns the number of rows, columns and stride of matrix A.
// The matrix must be allocated with MakeFloat32Matrix to ensure contiguous underlying storage.
// Stride is the distance between vertically adjacent elements. For freshly allocated matrices,
// stride == cols, while for submatrices stride <= cols.
func SSize(A [][]float32) (rows, cols, stride int) {
	rows = len(A)
	if rows > 0 {
		cols = len(A[0])
	}
	stride = cols // value when there's only one row (stride unknown but irrelevant, cols is safe)
	if rows > 1 {
		stride = int(uintptr(unsafe.Pointer(&A[1][0]))-uintptr(unsafe.Pointer(&A[0][0]))) / sizeof_float32
	}
	checkStorage(cap(A[0]), rows, stride)
	return
}

// Returns the number of rows, columns and stride of matrix A.
// The matrix must be allocated with MakeFloat64Matrix to ensure contiguous underlying storage.
// Stride is the distance between vertically adjacent elements. For freshly allocated matrices,
// stride == cols, while for submatrices stride <= cols.
func DSize(A [][]float64) (rows, cols, stride int) {
	rows = len(A)
	if rows > 0 {
		cols = len(A[0])
	}
	stride = cols // value when there's only one row (stride unknown but irrelevant, cols is safe)
	if rows > 1 {
		stride = int(uintptr(unsafe.Pointer(&A[1][0]))-uintptr(unsafe.Pointer(&A[0][0]))) / sizeof_float64
	}
	checkStorage(cap(A[0]), rows, stride)
	return
}

// Returns the number of rows, columns and stride of matrix A.
// The matrix must be allocated with MakeComplex64Matrix to ensure contiguous underlying storage.
// Stride is the distance between vertically adjacent elements. For freshly allocated matrices,
// stride == cols, while for submatrices stride <= cols.
func CSize(A [][]complex64) (rows, cols, stride int) {
	rows = len(A)
	if rows > 0 {
		cols = len(A[0])
	}
	stride = cols // value when there's only one row (stride unknown but irrelevant, cols is safe)
	if rows > 1 {
		stride = int(uintptr(unsafe.Pointer(&A[1][0]))-uintptr(unsafe.Pointer(&A[0][0]))) / sizeof_complex64
	}
	checkStorage(cap(A[0]), rows, stride)
	return
}

// Returns the number of rows, columns and stride of matrix A.
// The matrix must be allocated with MakeComplex128Matrix to ensure contiguous underlying storage.
// Stride is the distance between vertically adjacent elements. For freshly allocated matrices,
// stride == cols, while for submatrices stride <= cols.
func ZSize(A [][]complex128) (rows, cols, stride int) {
	rows = len(A)
	if rows > 0 {
		cols = len(A[0])
	}
	stride = cols // value when there's only one row (stride unknown but irrelevant, cols is safe)
	if rows > 1 {
		stride = int(uintptr(unsafe.Pointer(&A[1][0]))-uintptr(unsafe.Pointer(&A[0][0]))) / sizeof_complex128
	}
	checkStorage(cap(A[0]), rows, stride)
	return
}

// sanity check that should catch many cases where matrix was not properly constructed
func checkStorage(capacity, rows, stride int) {
	if capacity < rows*stride || stride < 1 {
		panic("blas: non-contiguous matrix, should have been allocated with MakeXXXXMatrix.")
	}
}

const (
	sizeof_float32    = 4
	sizeof_float64    = 8
	sizeof_complex64  = 8
	sizeof_complex128 = 16
)

// Returns a submatrix of the matrix M.
// The upper-left element of the submatrix is the element (k1,k2) of the original matrix.
// The submatrix has n1 rows and n2 columns. The physical number of columns (stride) in memory is unchanged.
// The returned matrix shares underlying storage with the original: modifying either one will affect the other.
func SSubmatrix(M [][]float32, k1, k2, n1, n2 int) [][]float32 {
	m := make([][]float32, n1)
	for i := range m {
		m[i] = M[k1+i][k2 : k2+n2]
	}
	return m
}

// Returns a submatrix of the matrix M.
// The upper-left element of the submatrix is the element (k1,k2) of the original matrix.
// The submatrix has n1 rows and n2 columns. The physical number of columns (stride) in memory is unchanged.
// The returned matrix shares underlying storage with the original: modifying either one will affect the other.
func DSubmatrix(M [][]float64, k1, k2, n1, n2 int) [][]float64 {
	m := make([][]float64, n1)
	for i := range m {
		m[i] = M[k1+i][k2 : k2+n2]
	}
	return m
}

// Returns a submatrix of the matrix M.
// The upper-left element of the submatrix is the element (k1,k2) of the original matrix.
// The submatrix has n1 rows and n2 columns. The physical number of columns (stride) in memory is unchanged.
// The returned matrix shares underlying storage with the original: modifying either one will affect the other.
func CSubmatrix(M [][]complex64, k1, k2, n1, n2 int) [][]complex64 {
	m := make([][]complex64, n1)
	for i := range m {
		m[i] = M[k1+i][k2 : k2+n2]
	}
	return m
}

// Returns a submatrix of the matrix M.
// The upper-left element of the submatrix is the element (k1,k2) of the original matrix.
// The submatrix has n1 rows and n2 columns. The physical number of columns (stride) in memory is unchanged.
// The returned matrix shares underlying storage with the original: modifying either one will affect the other.
func ZSubmatrix(M [][]complex128, k1, k2, n1, n2 int) [][]complex128 {
	m := make([][]complex128, n1)
	for i := range m {
		m[i] = M[k1+i][k2 : k2+n2]
	}
	return m
}
