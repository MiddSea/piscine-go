package piscine

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}

// from q05 index
// And its output :

// $ go run .
// 2
// 1
// -1
// $

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	fmt.Println(strings.Index("chandler", "and"))
// 	fmt.Println(strings.Index("navy", "as"))
// }
