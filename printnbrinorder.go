package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	printStr(int2str(n))
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func int2str(n int) (s string) {
	for n%10 > 0 {
		s = string(rune(n%10)+'0') + s
		n /= 10
	}
	return
}
