package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		fmt.Println("Letter:", string(char))
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {
	go printNumbers() // Start a goroutine to print numbers concurrently
	printLetters()    // Print letters in the main goroutine

	// Wait for a while to allow goroutines to complete
	time.Sleep(time.Second * 3)
	fmt.Println("Done")
}
