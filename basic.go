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

// findIndex takes a row and column and returns the corresponding index
// from the underlying data slice.

func findIndex(r, c int, A *Matrix) int {
	return (r-1)*A.columns + (c - 1)
}