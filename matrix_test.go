package blas

import "testing"

func TestSSize(t *testing.T) {
	R := 100
	C := 5
	A := MakeFloat32Matrix(R, C)
	r, c, s := SSize(A)
	if r != R || c != C || s != C {
		t.Fail()
	}
}

func TestSSubmatrix(t *testing.T) {
	R := 100
	C := 200
	A := MakeFloat32Matrix(R, C)

	k1, k2 := 17, 13
	n1, n2 := 6, 19

	A[k1][k2] = 1

	a := SSubmatrix(A, k1, k2, n1, n2)

	r, c, s := SSize(a)
	if r != n1 || c != n2 || s != C {
		t.Error("r,c,s=", r, c, s)
	}

	if a[0][0] != 1 {
		t.Fail()
	}
}

func TestMatrixCheck(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fail()
		}
	}()
	badMatrix := make([][]float32, 10)
	for i := range badMatrix {
		badMatrix[i] = make([]float32, 20)
	}
	_, _, _ = SSize(badMatrix) // should panic
}
