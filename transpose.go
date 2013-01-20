package matrix

// Transpose takes an m x n matrix and returns an n x m matrix where the rows 
// in the  first matrix are now the columns of the second.

func (A *Matrix) Transpose() *Matrix {
	var data []int
	for i := 1; i <= A.columns; i++ {
		col := A.Column(i)
		data = append(data, col...)
	}
	B := New(A.columns, A.rows, data)
	return B
}