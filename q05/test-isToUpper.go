package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ToUpper("abcABA d d d ... .[[87/||\\]]"))
	fmt.Println(piscine.ToUpper("abcABA d d d ... .[[87/||\\]]"))
	fmt.Println(piscine.ToUpper("abcABA d243fn fe  d d ... .[[87/||\\]]"))
	fmt.Println(piscine.ToUpper("adfsa	s243fn fe df asdsfa asdfdsa "))
	fmt.Println(piscine.ToUpper("a.,.,. e243fn fe adsf a.-., /\\"))
	fmt.Println(piscine.ToUpper("a true "))
}

// testing to upper
