package piscine

import "github.com/01-edu/z01"

func printNum(num int) {
	z01.PrintRune(rune(num/10 + '0'))
	z01.PrintRune(rune(num%10 + '0'))
}

func PrintComb2() {
	num1 := 0
	for num1 < 100 {
		num2 := 0
		for num2 < 100 {
			if num1 < num2 {
				printNum(num1)
				z01.PrintRune(' ')
				printNum(num2)
				if num1 < 98 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
			num2++
		}
		num1++
	}
	z01.PrintRune('\n')
}
