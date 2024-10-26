package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Split("Hello! How are you?", " "))
	fmt.Println(piscine.Split("  Hello! \t\nHow  are you?", "\t\n "))

}
