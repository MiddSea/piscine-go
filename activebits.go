// piscine/active_bits.go
package piscine

// ActiveBits returns the number of active bits (bits with the value 1) in the binary representation of an integer.
func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1 // Check if the least significant bit is 1
		n >>= 1        // Right shift to check the next bit
	}
	return count
}

// Write a function, ActiveBits, that returns the number of active bits (bits with the value 1) in the binary representation of an integer number.

// Expected function
// func ActiveBits(n int) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.ActiveBits(7))
// }
// And its output :

// $ go run .
// 3
// $
// Something is wrong ? Submit an issue or even better propose a change !
