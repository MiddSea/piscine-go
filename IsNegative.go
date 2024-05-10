package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	isNegResult := 'F'
	if nb < 0 {
		isNegResult = 'T'
	}
	z01.PrintRune(isNegResult)
	z01.PrintRune('\n')

}
