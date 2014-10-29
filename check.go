package blas

import (
	"fmt"
)

func checkSizeS(s ...[]float32) {
	N := len(s[0])
	for _, other := range s {
		if len(other) != N {
			panic(fmt.Sprintf("slice length mismatch: %v != %v", N, len(other)))
		}
	}
}
