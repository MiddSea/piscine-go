package piscine

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("abc\n"))
	fmt.Println(piscine.IsPrintable("a\bc\a"))
	fmt.Println(piscine.IsPrintable("ab\tc"))
	fmt.Println(piscine.IsPrintable("adfsa	sdf asdsfa asdfdsa "))
	fmt.Println(piscine.IsPrintable("a.,.,. eadsf a.-., /\\"))
	fmt.Println(piscine.IsPrintable("a true "))
}
