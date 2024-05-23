package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	n := 0
	for n < len(s) {
		z01.PrintRune(rune(s[n]))
		n++
	}
}
