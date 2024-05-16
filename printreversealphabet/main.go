package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// **** alphabet forwards ***
	// alphabet := "abcdefghijklmnopqrstuvwxyz"
	// *** alphabet backwards ***
	// a so called "raw character literal string"
	// alphabet := `zyxwvutsrqponmlkjihgfedcba`

	alphabet := "zyxwvutsrqponmlkjihgfedcba"

	for _, value := range alphabet {
		z01.PrintRune(value)
	}

	z01.PrintRune('\n')
}

// could reverse it by doing