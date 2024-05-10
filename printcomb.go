package piscine

import (
	"github.com/01-edu/z01"
	
)

func PrintComb() {

	for i := 0; i < 9; i++ {
		// z01.PrintRune(i)
		print(i,", ")
	}	

	z01.PrintRune('\n')
}
