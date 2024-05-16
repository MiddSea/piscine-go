package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	printStr(strSort(int2str(n)))
}
func strSort(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)-1; i++ {
		for j := 0; j < len(rs)-i-1; j++ {
			if rs[j] > rs[j+1] {
				rs[j], rs[j+1] = rs[j+1], rs[j]
			}
		}
	}
	return string(rs)
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
