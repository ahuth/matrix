package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// String builds a human-readable string that represents the matrix.
func (A *matrix) String() string {

	// Find the width that each column needs to be.
	columnWidths := make([]int, A.columns)

	for i := range columnWidths {
		var maxLength int
		thisColumn := A.Column(i + 1)
		for j := range thisColumn {
			thisLength := len(strconv.Itoa(thisColumn[j]))
			if thisLength > maxLength {
				maxLength = thisLength
			}
		}
		columnWidths[i] = maxLength
	}

	// Now we know the width of each column, so build the string by converting
	// each element to a string and padding with the required number of spaces.
	var matrixString string

	for i := 0; i < A.rows; i++ {
		thisRow := A.Row(i + 1)
		matrixString += "["
		for j := range thisRow {
			numSpaces := columnWidths[j] - len(strconv.Itoa(thisRow[j]))
			matrixString += strings.Repeat(" ", numSpaces) + strconv.Itoa(thisRow[j])
			if j < len(thisRow)-1 {
				matrixString += " "
			}
		}
		matrixString += "]\n"
	}

	return matrixString
}

// Print converts the matrix into a string and then outputs it to fmt.Print.
func (A *matrix) Print() {
	fmt.Print(A.String())
}
