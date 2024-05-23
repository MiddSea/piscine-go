package main

import (
	"fmt"
	"piscine"
)

func main() {
	// Define a slice of integers
	a := []int{23, 123, 1, 11, 55, 93}

	// Call the Max function to find the maximum value in the slice
	max := piscine.Max(a)

	// Print the maximum value
	fmt.Println(max)
}
