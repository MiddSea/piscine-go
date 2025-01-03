package piscine

/*
iscapitalized
Instructions
Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter or a non-alphabetic character.

If any of the words begin with a lowercase letter return false.
If the string is empty return false.
Expected function
func IsCapitalized(s string) bool {

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
And its output:

$ go run .
false
true
true
true
true
false
*/

// isCaptialized takes a string, checks if first letter of word is ASCII capitals 
func IsCapitalized(s string) bool {
	words := []rune(s)
	if len(words) == 0 {
		return false
	}
	for i := 0; i < len(words); i++ {
		// if current character is the first letter of string and is lowercase, return false
		if i == 0 && words[i] >= 'a' && words[i] <= 'z' {
			return false
		}

		// if previous character is a space or the first character and current character is lowercase, return false
		if i > 0 && words[i-1] == ' ' && words[i] >= 'a' && words[i] <= 'z' {
			return false
		}
		// if current character is uppercase, skip to next space
		if words[i] >= 'A' && words[i] <= 'Z' {
			for i < len(words) && words[i] != ' ' {
				i++
			}
		}
	}
	return true
}