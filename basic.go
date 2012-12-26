// Package matrix implements a simple library for creating and
// manipulating matrices, and performing basic linear algebra.

package matrix

type Matrix struct {
	rows, columns int   // the number of rows and columns.
	data          []int // the contents of the matrix as one long slice.
}

// Set lets you define the value of a matrix at the given row and
// column.

func (A *Matrix) Set(r, c, val int) {
	A.data[findIndex(r, c, A)] = val
}

// Get retrieves the contents of the matrix at the row and column.

func (A *Matrix) Get(r, c int) int {
	return A.data[findIndex(r, c, A)]
}

// Column returns a slice that represents a column from the matrix.
// This works by examining each row, and adding the nth element of
// each to the column slice.

func (A *Matrix) Column(n int) []int {
	col := make([]int, A.rows)
	for i := 1; i <= A.rows; i++ {
		col[i-1] = A.Row(i)[n-1]
	}
	return col
}

// Row returns a slice that represents a row from the matrix.

func (A *Matrix) Row(n int) []int {
	return A.data[findIndex(n, 1, A):findIndex(n, A.columns+1, A)]
}

// Identity creates an identity matrix with n rows and n columns.  When you
// multiply any matrix by its corresponding identity matrix, you get the
// original matrix.

func Identity(n int) Matrix {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i += (n + 1) {
		A.data[i] = 1
	}
	return A
}

// Zeros creates an r x c sized matrix that's filled with zeros.  The initial
// state of an int is 0, so we don't have to do any initialization.

func Zeros(r, c int) Matrix {
	return Matrix{r, c, make([]int, r*c)}
}

// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.

func New(r, c int, data []int) Matrix {
	if len(data) != r*c {
		panic("[]int data provided to matrix.New is greater than the provided capacity of the matrix!'")
	}
	A := Zeros(r, c)
	A.data = data
	return A
}

// findIndex takes a row and column and returns the corresponding index
// from the underlying data slice.

func findIndex(r, c int, A *Matrix) int {
	return (r-1)*A.columns + (c - 1)
}