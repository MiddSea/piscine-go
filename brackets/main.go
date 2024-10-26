package main

import (
	"fmt"
	"os"
)

func isBracketed(s string) bool {
	var stack []rune
	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if brackets[stack[len(stack)-1]] != r {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	for _, arg := range os.Args[1:] {
		if isBracketed(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}

}

/*brackets
Instructions
Write a program that takes an undefined number of string in arguments. For each argument, if the expression is correctly bracketed, the program prints on the standard output OK followed by a newline ('\n'), otherwise it prints Error followed by a newline.

Symbols considered as brackets are parentheses ( and ), square brackets [ and ] and curly braces { and }. Every other symbols are simply ignored.

An opening bracket must always be closed by the good closing bracket in the correct order. A string which does not contain any bracket is considered as a correctly bracketed string.

If there is no argument, the program must print nothing.

Usage
$ go run . '(johndoe)' | cat -e
OK$
$ go run . '([)]' | cat -e
Error$
$ go run . '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
OK$
OK$
$ go run .
$
*/
