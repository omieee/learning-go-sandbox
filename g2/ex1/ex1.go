package ex1

///
// Write a program that declares an integer variable called i with the value 20.
// Assign i to a floating-point variable named f. Print out i and f.
///
import "fmt"

func ex1() (int, float64) {
	i := 20
	f := float64(i)
	fmt.Println("I:", i)
	fmt.Println("F:", f)
	return i, f
}
