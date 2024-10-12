package main

import (
	"github.com/01-edu/z01"
	"piscine"
)

// func toUpper(r rune) 

func camelCase(s string) string {
  return "CamelCase: " + s + " : CamelCase" 
} 


func main() {
  z01.PrintRune('a')
  piscine.PrintStrUsingRunes("Hello World")
  piscine.PrintStrUsingRunes(camelCase("Hello world"))
}