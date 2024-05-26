package piscine

import "github.com/01-edu/z01"

// PrintNbr prints an integer to the standard output.
func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	var digits []rune
	for n != 0 {
		digits = append(digits, rune(n%10)+'0')
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
