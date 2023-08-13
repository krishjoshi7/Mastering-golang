package main

import "fmt"

func main() {
	// Creating a slice
	numbers := []int{1, 2, 3, 4, 5}

	// Displaying the original slice
	fmt.Println("Original slice:", numbers)

	// Slicing the slice to create a new slice
	slicedNumbers := numbers[1:4]

	// Displaying the new slice
	fmt.Println("Sliced slice:", slicedNumbers)

	// Modifying the new slice
	slicedNumbers[1] = 99

	// Displaying both slices to see the changes
	fmt.Println("Original slice after modification:", numbers)
	fmt.Println("Sliced slice after modification:", slicedNumbers)

	// Appending to a slice
	numbers = append(numbers, 6, 7, 8)

	// Displaying the updated slice
	fmt.Println("Updated slice after appending:", numbers)
}
