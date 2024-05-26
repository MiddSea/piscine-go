// Usage
// Here is a possible program to test your function :
package main

import "piscine"

func main() {
	a := []int{6, 1, 2, 35, 4, 5, 6, 7}
	piscine.ForEach(piscine.PrintNbr, a)
	piscine.WriteHelloWorld("Now then, Now heeek")
}

// And its output :

//$ go run .
//123456
//$
