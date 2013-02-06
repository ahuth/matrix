package matrix

// Multiply performs matrix multiplication by finding the dot product of all
// row/column combinations.
func Multiply(A, B *matrix) *matrix {
	C := Zeros(A.rows, B.columns)
	for r := 1; r <= C.rows; r++ {
		for c := 1; c <= C.columns; c++ {
			C.Set(r, c, dotProduct(A.Row(r), B.Column(c)))
		}
	}
	return C
}

// Add performs matrix addition by summing corresponding elements from each matrix.
func Add(A, B *matrix) *matrix {
	C := Zeros(A.rows, A.columns)
	for r := 1; r <= A.rows; r++ {
		for c := 1; c <= A.columns; c++ {
			C.Set(r, c, A.Get(r, c)+B.Get(r, c))
		}
	}
	return C
}

// dotProduct sums the product of corresponding elements from two slices.
func dotProduct(a, b []int) int {
	var total int
	for i := 0; i < len(a); i++ {
		total += a[i] * b[i]
	}
	return total
}
