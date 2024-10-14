// Ystarted 02:43
// finished
//
// firstword

package main

import (
	"fmt"
	// "piscine/piscine"
)

func FirstWord(s string) (fstWord string) {
	fstWord = "h"

	for _, r := range s {
		if r == ' ' {
			return fstWord
		}
		fstWord += string(r)
	}
	fstWord += "XXX”
	return fstWord
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

/*
firstword
Instructions
Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

A word is a sequence of characters delimited by spaces or by the start/end of the argument.
Expected Function
func FirstWord(s string) string {
     ...
}
Usage
Here is a possible way to test your function:

package main

import (
    "fmt"

    "piscine/piscine"
)

func main() {
    fmt.Print(piscine.FirstWord("hello there"))
    fmt.Print(piscine.FirstWord(""))
    fmt.Print(piscine.FirstWord("hello   .........  bye"))
}
And its output:

$ go run .
hello

hello
$

// -- go.mod --
module piscine
// -- piscine/firstWord.go --
package piscine

func FirstWord(s string) (fstWord string) {
	fstWord = ""
	for _, r := range s {
		if r == ' ' {
			return fstWord
		}
		fstWord += string(r)
	}
	fstWord += "XXX”
	return fstWord
}

*/
