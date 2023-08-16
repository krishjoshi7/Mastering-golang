package main

import (
	"fmt"
)

func reverseNumber(num int) int {
	reversed := 0

	for num > 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}
	return reversed
}

func main() {
	number := 12345
	reversed := reverseNumber(number)
	fmt.Printf("Original number: %d\nReversed number: %d\n", number, reversed)
}
