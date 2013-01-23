// Package matrix implements a simple library for creating and
// manipulating matrices, and performing basic linear algebra.

package matrix

type matrix struct {
	rows, columns int   // the number of rows and columns.
	data          []int // the contents of the matrix as one long slice.
}

// Define the element of a matrix at a given row and column.
func (A *matrix) Set(r, c, val int) {
	A.data[A.findIndex(r, c)] = val
}

// Retrieve the contents of a matrix from a given row and column.
func (A *matrix) Get(r, c int) int {
	return A.data[A.findIndex(r, c)]
}

// Return a slice representing the column of a matrix.  We find this by examining
// each row, and adding the nth element of each to a slice.
func (A *matrix) Column(n int) []int {
	col := make([]int, A.rows)
	for i := 1; i <= A.rows; i++ {
		col[i-1] = A.Row(i)[n-1]
	}
	return col
}

// Return an entire row from a matrix.
func (A *matrix) Row(n int) []int {
	return A.data[A.findIndex(n, 1):A.findIndex(n, A.columns+1)]
}

// Transpose takes an m x n matrix and returns an n x m matrix where the rows 
// in the  first matrix are now the columns of the second.
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
