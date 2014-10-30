package blas

import "fmt"

func MakeFloat32Matrix(rows, cols int)[][]float32{
	checkSize(rows, cols)
	a := make([]float32, rows*cols)
	return reshapeR2(a, [2]int{rows, cols})
}

func checkSize(rows, cols int){
	//TODO
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeR2(array []float32, N [2]int) [][]float32 {
	return reshapeR3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeR3(array []float32, N [3]int) [][][]float32 {
	return reshapeR4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeR4(array []float32, N [4]int) [][][][]float32 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]float32, N[0])
	for i := range sliced {
		sliced[i] = make([][][]float32, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]float32, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeC2(array []complex64, N [2]int) [][]complex64 {
	return reshapeC3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeC3(array []complex64, N [3]int) [][][]complex64 {
	return reshapeC4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeC4(array []complex64, N [4]int) [][][][]complex64 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]complex64, N[0])
	for i := range sliced {
		sliced[i] = make([][][]complex64, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]complex64, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeD2(array []float64, N [2]int) [][]float64 {
	return reshapeD3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeD3(array []float64, N [3]int) [][][]float64 {
	return reshapeD4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeD4(array []float64, N [4]int) [][][][]float64 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]float64, N[0])
	for i := range sliced {
		sliced[i] = make([][][]float64, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]float64, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeZ2(array []complex128, N [2]int) [][]complex128 {
	return reshapeZ3(array, [3]int{1, N[0], N[1]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeZ3(array []complex128, N [3]int) [][][]complex128 {
	return reshapeZ4(array, [4]int{1, N[0], N[1], N[2]})[0]
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func reshapeZ4(array []complex128, N [4]int) [][][][]complex128 {
	if prod(N[:]) != len(array) {
		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
	}
	sliced := make([][][][]complex128, N[0])
	for i := range sliced {
		sliced[i] = make([][][]complex128, N[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]complex128, N[2])
		}
	}
	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
			}
		}
	}
	return sliced
}

func prod(s []int) int {
	p := 1
	for _, s := range s {
		p *= s
	}
	return p
}
