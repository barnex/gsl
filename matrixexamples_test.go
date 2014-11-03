package blas

import (
	"fmt"
)

func ExampleMakeFloat32Matrix() {
	m := MakeFloat32Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[0 0 0] [0 0 0]]
}

func ExampleMakeFloat64Matrix() {
	m := MakeFloat64Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[0 0 0] [0 0 0]]
}

func ExampleMakeComplex64Matrix() {
	m := MakeComplex64Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}

func ExampleMakeComplex128Matrix() {
	m := MakeComplex128Matrix(2, 3)
	fmt.Println(m)

	//Output:
	//[[(0+0i) (0+0i) (0+0i)] [(0+0i) (0+0i) (0+0i)]]
}
