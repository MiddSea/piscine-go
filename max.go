package piscine

// Max returns the maximum value in a slice of integers. If the slice is empty, it returns 0.
func Max(a []int) int {
	// Check if the slice is empty
	if len(a) == 0 {
		return 0
	}

	// Assume the first element is the maximum initially
	maxVal := a[0]

	// Iterate over the slice starting from the second element
	for _, value := range a[1:] {
		// Update maxVal if the current element is greater
		if value > maxVal {
			maxVal = value
		}
	}

	// Return the maximum value found
	return maxVal
}

// max
// FAILED
// LEVEL 29
// REQUIRED
// XP
// Files to submit
// Allowed functions
// 4.70 kB
// max.go
// --allow-builtin
// Instructions
// Write a function Max that will return the maximum value in a slice of integers. If the slice is empty it will return 0.

// Expected function
// func Max(a []int) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := []int{23, 123, 1, 11, 55, 93}
// 	max := piscine.Max(a)
// 	fmt.Println(max)
// }
// And its output :

// $ go run .
// 123
// $
