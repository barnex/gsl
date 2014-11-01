/*
package blas provides an interface to GSL's BLAS library.

Matrix representation

The library assumes that arrays, vectors and matrices passed as modifiable arguments are not aliased and do not overlap with each other. This removes the need for the library to handle overlapping memory regions as a special case, and allows additional optimizations to be used. If overlapping memory regions are passed as modifiable arguments then the results of such functions will be undefined. If the arguments will not be modified then overlapping or aliased memory regions can be safely used.


Naming

Each routine has a name which specifies the operation, the type of matrices involved and their precisions. Some of the most common operations and their names are given below,
 	DOT scalar product, x^T y
 	AXPY vector sum, \alpha x + y
 	MV matrix-vector product, A x
 	SV matrix-vector solve, inv(A) x
 	MM matrix-matrix product, A B
 	SM matrix-matrix solve, inv(A) B

The types of matrices are,

 	GE general
 	GB general band
 	SY symmetric
 	SB symmetric band
 	SP symmetric packed
 	HE hermitian
 	HB hermitian band
 	HP hermitian packed
 	TR triangular
 	TB triangular band
 	TP triangular packed

Each operation is defined for four precisions,

 	S single real
 	D double real
 	C single complex
 	Z double complex

Thus, for example, the name SGEMM stands for “single-precision general matrix-matrix multiply” and ZGEMM stands for “double-precision complex matrix-matrix multiply”.
*/
package blas
