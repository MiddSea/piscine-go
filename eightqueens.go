package piscine

import "github.com/01-edu/z01"

func printResultQueens(result [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(result[i] + '0'))
	}
	z01.PrintRune('\n')
}

func checkQueens(index int, tab [8]int) bool {
	for i := 0; i < index; i++ {
		diff := i - index
		if tab[i] == tab[index] || tab[i]+diff == tab[index] || tab[i]-diff == tab[index] {
			return false
		}
	}
	return true
}

func solveQueens(index int, tab [8]int) {
	if index == 8 {
		printResultQueens(tab)
		return
	}
	for i := 1; i <= 8; i++ {
		tab[index] = i
		if checkQueens(index, tab) {
			solveQueens(index+1, tab)
		}
	}
}

func EightQueens() {
	var tab [8]int
	for i := 0; i < 8; i++ {
		tab[i] = 1
	}
	solveQueens(0, tab)
}
