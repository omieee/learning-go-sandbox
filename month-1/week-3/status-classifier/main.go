package main

import "fmt"

func classifyStatusCode(statusCode int) string {
	if statusCode >= 200 && statusCode <= 299 {
		return "PASS"
	}
	if statusCode >= 300 && statusCode <= 399 {
		return "WARN"
	}
	if statusCode >= 400 && statusCode <= 499 {
		return "BLOCK"
	}
	if statusCode >= 500 && statusCode <= 599 {
		return "BLOCK`"
	}
	return "UNKNOWN"
}

func main() {
	fmt.Println("Status code is: " + classifyStatusCode(327))
}
