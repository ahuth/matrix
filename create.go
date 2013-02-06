package matrix

// Identity creates an n x n identity matrix.
func Identity(n int) *matrix {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i += (n + 1) {
		A.data[i] = 1
	}
	return A
}

// Zeros creates an r x c matrix filled with zeros.
func Zeros(r, c int) *matrix {
	return &matrix{r, c, make([]int, r*c)}
}

// New creates an r x c matrix filled with the provided data slice.
func New(r, c int, data []int) *matrix {
	if len(data) != r*c {
		panic("[]int data provided to matrix.New() is greater than the provided capacity of the matrix!")
	}
	A := Zeros(r, c)
	A.data = data
	return A
}
