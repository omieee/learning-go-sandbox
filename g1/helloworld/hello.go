package helloworld

import (
	"fmt"
)

func Hello() string {
	return "Hello!, World!"
}

func HelloName(name string) string {
	return "Hello! " + name
}

const englishGreet = "Hello! "

func HelloNameWithGreet(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "spanish":
		return "Hola! " + name
	case "french":
		return "Bonjour! " + name
	}
	return englishGreet + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloName("Om Shanker"))
	fmt.Println(HelloNameWithGreet("Bauua", "french"))
}
