package matrix

import (
	"fmt"
	"strconv"
)

// Print converts the matrix into a string and then outputs it to fmt.Printf.
func (A *matrix) Print() {

	// Find the width (in characters) that each column needs to be.  We hold these
	// widths as strings, not ints, because we're going to use these in a printf
	// function.

	columnWidths := make([]string, A.columns)

	for i := range columnWidths {
		var maxLength int
		thisColumn := A.Column(i + 1)
		for j := range thisColumn {
			thisLength := len(strconv.Itoa(thisColumn[j]))
			if thisLength > maxLength {
				maxLength = thisLength
			}
		}
		columnWidths[i] = strconv.Itoa(maxLength)
	}

	// We have the widths, so now output each element with the correct column
	// width so that they line up properly.

	for i := 0; i < A.rows; i++ {
		thisRow := A.Row(i + 1)
		fmt.Printf("[")
		for j := range thisRow {
			var printFormat string
			if j == 0 {
				printFormat = "%" + columnWidths[j] + "s"
			} else {
				printFormat = " %" + columnWidths[j] + "s"
			}
			fmt.Printf(printFormat, strconv.Itoa(thisRow[j]))
		}
		fmt.Printf("]\n")
	}
}
