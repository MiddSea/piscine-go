package piscine

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println("testing piscine.ToLower")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println(piscine.ToLower("abcABA d d d ... .[[87/||\\]]"))
	fmt.Println(piscine.ToLower("abcABA d d d ... .[[87/||\\]]"))
	fmt.Println(piscine.ToLower("abcABA d243fn fe  d d ... .[[87/||\\]]"))
	fmt.Println(piscine.ToLower("adfsa	s243fn fe df asdsfa asdfdsa "))
	fmt.Println(piscine.ToLower("a.,.,. e243fn fe adsf a.-., /\\"))
	fmt.Println(piscine.ToLower("a true Story! "))
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("DONE testing piscine.ToLower")
}

// testing to upper done works 2024-05-16