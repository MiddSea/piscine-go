package piscine

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("HfgtARDfCODE"))
	fmt.Println(piscine.IsNumeric("Helflo! How are you?"))
	fmt.Println(piscine.IsNumeric("HelloHowareyou"))
	fmt.Println(piscine.IsNumeric("What's this 4?"))
	fmt.Println(piscine.IsNumeric("adsfasdf"))
	fmt.Println(piscine.IsNumeric("12345609"))
}