package main

import (
	"fmt"
	"matrix"
)

func main() {
	a := matrix.New(3, 2, []int{-1, 3, 0, 5, 2, 5})
	b := matrix.New(2, 3, []int{3, 1, 2, -2, 0, 5})
	c := matrix.Multiply(a, b)
	c.Print()
	fmt.Println("")
	d := c.Transpose()
	d.Print()
}
