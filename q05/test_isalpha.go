package piscine

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.IsAlpha("HfgtARDfCODE"))
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	fmt.Println(piscine.IsAlpha("HelloHowareyfou"))
	fmt.Println(piscine.IsAlpha("What's this 4?"))
	fmt.Println(piscine.IsAlpha("Wrthatsthis4"))
	fmt.Println(piscine.IsAlpha(""))
}
