package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// ****printdigits ***

	digits := "0123456789"

	for _, value := range digits {
		z01.PrintRune(value)
	}

	z01.PrintRune('\n')
}
