package ex2

import "fmt"

///
// Write a program that declares a constant called value that can be assigned to both an integer
// and a floating-point variable.Assign it to an integer called i and a floating-point variable
// called f. Print out i and f.
///

func ex2() (int, float64) {
	const value int = 10
	var i int = value
	var f float64 = float64(value)
	fmt.Println("I:", i)
	fmt.Println("F:", f)
	return i, f
}
