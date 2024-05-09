package main

import (

	"github.com/01-edu/z01"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, value := range alphabet {
		z01.PrintRune(value)
	}
	
	z01.PrintRune('\n')
}
