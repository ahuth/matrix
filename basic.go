// Package matrix implements a simple library for creating and
// manipulating matrices, and performing basic linear algebra.

package matrix

type matrix struct {
	rows, columns int   // the number of rows and columns.
	data          []int // the contents of the matrix as one long slice.
}

// Set allows you to define one element of a matrix.
func (A *matrix) Set(r, c, val int) {
	A.data[A.findIndex(r, c)] = val
}

// Get retrieves the contents of a matrix at a row and column.
func (A *matrix) Get(r, c int) int {
	return A.data[A.findIndex(r, c)]
}

// Column returns a slice representing one column of a matrix.
func (A *matrix) Column(n int) []int {
	col := make([]int, A.rows)
	for i := 1; i <= A.rows; i++ {
		col[i-1] = A.Get(i, n)
	}
	return col
}

// Row returns a slice representing one row of a matrx.
func (A *matrix) Row(n int) []int {
	return A.data[A.findIndex(n, 1):A.findIndex(n, A.columns+1)]
}

// Transpose takes an m x n matrix and returns an n x m matrix where the rows of
// the  first matrix are now the columns of the second.
func (A *matrix) Transpose() *matrix {
	var data []int
	for i := 1; i <= A.columns; i++ {
		col := A.Column(i)
		data = append(data, col...)
	}
	B := New(A.columns, A.rows, data)
	return B
}

// Translate a row and column into an index of a matrix's underlying slice.
func (A *matrix) findIndex(r, c int) int {
	return (r-1)*A.columns + (c - 1)
}
