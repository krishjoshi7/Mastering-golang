package main

import "fmt"

func main() {
   //arrays

   var arr [3]int
   arr[0] = 1
   arr[1] = 2
   arr[2] = 3

   //slices
   var slice = []int{1, 2, 3, 4, 5}

   fmt.Println("Array:", arr)
   fmt.Println("Slice:", slice)
}




