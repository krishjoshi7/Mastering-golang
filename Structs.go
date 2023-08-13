package main

import "fmt"

// Defining a struct
type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Creating an instance of the struct
	person := Person{
		Name:    "Alice",
		Age:     25,
		Country: "USA",
	}

	// Accessing struct fields
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Country:", person.Country)
}
