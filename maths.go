package matrix

// Multiply multiplies two matrices together and return the resulting matrix.
// For each element of the result matrix, we get the dot product of the
// corresponding row from matrix A and column from matrix B.

func Multiply(A, B *Matrix) *Matrix {
	C := Zeros(A.rows, B.columns)
	for r := 1; r <= C.rows; r++ {
		for c := 1; c <= C.columns; c++ {
			C.Set(r, c, dotProduct(A.Row(r), B.Column(c)));
		}
	}
	return C
}

// Add adds two matrices together and returns the resulting matrix.  To do
// this, we just add together the corresponding elements from each matrix.

func Add(A, B *Matrix) *Matrix {
	C := Zeros(A.rows, A.columns)
	for r := 1; r <= A.rows; r++ {
		for c := 1; c <= A.columns; c++ {
			C.Set(r, c, A.Get(r, c)+B.Get(r, c))
		}
	}
	return C
}

// dotProduct calculates the algebraic dot product of two slices.  This is just
// the sum  of the products of corresponding elements in the slices.  We use
// this when we multiply matrices together.

func dotProduct(a, b []int) int {
	var total int
	for i := 0; i < len(a); i++ {
		total += a[i] * b[i]
	}
	return total
}