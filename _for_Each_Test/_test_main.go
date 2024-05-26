// Usage
// Here is a possible program to test your function :
package main

import "piscine"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, a)
}

// And its output :

//$ go run .
//123456
//$