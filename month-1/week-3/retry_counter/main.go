package main

import "fmt"

func retryAttempts(maxRetries int) {
	if maxRetries <= 0 {
		fmt.Printf("No retries, as max retries is %v", maxRetries)
		return
	}
	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Attempt %v of %v\n", i, maxRetries)
	}
	fmt.Println("No more retries left")
}

func main() {
	retryAttempts(-10)
}
