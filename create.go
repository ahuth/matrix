package matrix

// Identity creates an identity matrix with n rows and n columns.  When you
// multiply any matrix by its corresponding identity matrix, you get the
// original matrix.

func Identity(n int) *Matrix {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i += (n + 1) {
		A.data[i] = 1
	}
	return A
}

// Zeros creates an r x c sized matrix that's filled with zeros.  The initial
// state of an int is 0, so we don't have to do any initialization.

func Zeros(r, c int) *Matrix {
	return &Matrix{r, c, make([]int, r*c)}
}

// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.

func New(r, c int, data []int) *Matrix {
	if len(data) != r*c {
		panic("[]int data provided to matrix.New is greater than the provided capacity of the matrix!'")
	}
	A := Zeros(r, c)
	A.data = data
	return A
}