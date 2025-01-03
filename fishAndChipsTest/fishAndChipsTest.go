// fishandchips

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FishAndChips(4))
	fmt.Println(piscine.FishAndChips(9))
	fmt.Println(piscine.FishAndChips(6))
		fmt.Println(piscine.FishAndChips(13))
	fmt.Println(piscine.FishAndChips(0))
	fmt.Println(piscine.FishAndChips(-12))

}

/*Instructions
Write a function called FishAndChips() that takes an int and returns a string.

If the number is divisible by 2, print fish.
If the number is divisible by 3, print chips.
If the number is divisible by 2 and 3, print fish and chips.
If the number is negative return error: number is negative.
If the number is non divisible by 2 or 3 return error: non divisible.
Expected function
func FishAndChips(n int) string {

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FishAndChips(4))
	fmt.Println(piscine.FishAndChips(9))
	fmt.Println(piscine.FishAndChips(6))
	fmt.Println(piscine.FishAndChips(13))
	fmt.Println(piscine.FishAndChips(0))
	fmt.Println(piscine.FishAndChips(-12))

}
And its output:

$ go run . | cat -e
fish$
chips$
fish and chips$
*/
