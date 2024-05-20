package main

import "github.com/01-edu/z01"

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	EvenMsg := "I have an odd number of arguments"
	OddMsg := "I have an even number of arguments"

	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
