package matrix

// Create an n x n identity matrix.
func Identity(n int) *matrix {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i += (n + 1) {
		A.data[i] = 1
	}
	return A
}

// Create an r x c matrix filled with zeros.  The empty state of an int is 0,
// so we don't have to set any data.
func Zeros(r, c int) *matrix {
	return &matrix{r, c, make([]int, r*c)}
}

// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.
func New(r, c int, data []int) *matrix {
	if len(data) != r*c {
		panic("[]int data provided to matrix.New is greater than the provided capacity of the matrix!'")
	}
	A := Zeros(r, c)
	A.data = data
	return A
}
