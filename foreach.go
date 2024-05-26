// Instructions
// Write a function ForEach that, for an int slice, applies a function on each element of that slice.

// Expected function
package piscine

// TO BE DONE
// import

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// Usage
// Here is a possible program to test your function :

// package main

// import "piscine"

// func main() {
//	a := []int{1, 2, 3, 4, 5, 6}
//	piscine.ForEach(piscine.PrintNbr, a)
// }

// flaged at 2024-05-26 09:48
// And its output :

// $ go run .
// 123456
// $
