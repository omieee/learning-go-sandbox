package helloworld

import "fmt"

func Hello() string {
	return "Hello!, World!"
}

func HelloName(name string) string {
	return "Hello! " + name
}

const greet = "Hello"

func HelloNameWithGreet(name string) string {
	return greet + " " + name + "!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloName("Om Shanker"))
	fmt.Println(HelloNameWithGreet("Bauua"))
}
