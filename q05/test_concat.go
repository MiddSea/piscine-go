package piscine

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.IsAlpha("Hello! ff How are you?"))
	fmt.Println(piscine.IsAlpha("HelloAgfaikn"))
	fmt.Println(piscine.IsAlpha("HelloAZ[\\]^_`again"))
}

// And its output :

// $ go run .
// Hello! How are you?
// $
