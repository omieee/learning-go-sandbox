package ex3

import "fmt"

///Write a program with three variables, one named b of type byte, one named smallI of type int32,
//and one named bigI of type uint64. Assign each variable the maximum legal value for its type;
//then add 1 to each variable. Print out their values.
///

func ex3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
