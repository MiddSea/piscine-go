package q05

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello 78 World!f AB  11 4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}

// And its output :

// $ go run .
// 10
// $
