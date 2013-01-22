package matrix

// Multiply two matrices by finding the dot products of each row-column
// combination.
func Multiply(A, B *Matrix) *Matrix {
	C := Zeros(A.rows, B.columns)
	for r := 1; r <= C.rows; r++ {
		for c := 1; c <= C.columns; c++ {
			C.Set(r, c, dotProduct(A.Row(r), B.Column(c)));
		}
	}
	return C
}

// Add two matrices together by adding corresponding elements of each.
func Add(A, B *Matrix) *Matrix {
	C := Zeros(A.rows, A.columns)
	for r := 1; r <= A.rows; r++ {
		for c := 1; c <= A.columns; c++ {
			C.Set(r, c, A.Get(r, c)+B.Get(r, c))
		}
	}
	return C
}

// Find the dot product by summing the products of corresponding elements of
// two slices.
func dotProduct(a, b []int) int {
	var total int
	for i := 0; i < len(a); i++ {
		total += a[i] * b[i]
	}
	return total
}