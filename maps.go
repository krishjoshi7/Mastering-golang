package main

import "fmt"

func main() {
	// Creating a map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Eve":   28,
	}

	// Accessing values from the map
	fmt.Println("Alice's age:", ages["Alice"])

	// Adding a new key-value pair
	ages["Charlie"] = 22

	// Deleting a key-value pair
	delete(ages, "Eve")

	// Iterating over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
