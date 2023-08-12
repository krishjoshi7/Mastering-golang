package main

import "fmt" 

func main() {
	// For loop
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration", i)
	}
	//while loop (Go only has 'for' loop)
	sum := 0
	for sum < 10 {
		sum += 2

	}
	fmt.Println("Sum:", sum)
	
}
