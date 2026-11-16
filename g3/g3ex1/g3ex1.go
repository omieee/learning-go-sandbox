package g3ex1

import "fmt"

///Write a program that defines a variable named greetings of type slice of strings with the following values:
// "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
// Create a subslice containing the first two values; a second subslice with the second, third, and fourth values;
// and a third subslice with the fourth and fifth values. Print out all four slices.

func g3ex1() ([]string, []string, []string) {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	gp1 := greetings[:2]
	gp2 := greetings[1:4]
	gp3 := greetings[3:]
	fmt.Println(gp1, gp2, gp3)
	return gp1, gp2, gp3
}
