// fromto
package main

import (
	"fmt"
	"piscine"
)

func main() {
	//	fmt.Print(piscine.FromTo(11, 29))
	//	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.FromTo(55, 11))
	//	fmt.Print(piscine.FromTo(32, 7))
	fmt.Print(piscine.FromTo(100, 10))
}

/*
Instructions
Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

The numbers must be separated by a comma and a space.
If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
Prepend a 0 to any number that is less than 10.
Add a new line \n at the end of the string.
Expected function
func FromTo(from int, to int) string {
// fromto
}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
	"psicine"
)

func main() {
	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.FromTo(10, 1))
	fmt.Print(piscine.FromTo(10, 10))
	fmt.Print(piscine.FromTo(100, 10))
}
And its output:

$ go run . | cat -e
01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
10$
Invalid$
*/
