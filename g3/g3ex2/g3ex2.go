package g3ex2

import "fmt"

///Write a program that defines a string variable called message with the value
// "Hi 👩 and 👨" and prints the fourth rune in it as a character, not a number.

func g3ex2() {
	message := "Hi 👩 and 👨"
	ru := []rune(message)
	fmt.Printf("%c", ru[3])
}
