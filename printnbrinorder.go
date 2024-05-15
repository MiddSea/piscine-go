package piscine

import (
	"github.com/01-edu/z01"
	"github.com/01-edu/z01.PrintRune"
)

func PrintNbrInOrder(n int) {

	for chRune := []rune(n) {
		z01.PrintRune(chRune)
		z01.PrintRune('#')
		
	}
	z01.PrintRune('\n')
	z01.PrintRune('\b')
}
