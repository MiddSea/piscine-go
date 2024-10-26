package main

import ( 
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.braketss("Hello! Ho{w ar[e ]}you?", " "))
	fmt.Println(piscine.brackets("  Hello! \t\nHow  are you?", "\t\n "))
}