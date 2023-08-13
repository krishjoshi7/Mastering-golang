package main

import "fmt"

// Function that adds two numbers
func add(a, b int) int {
	return a + b
}

func main() {
	result := add(10, 20)
	fmt.Println("Result of addition:", result)
}
