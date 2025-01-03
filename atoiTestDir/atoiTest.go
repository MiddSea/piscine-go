package main

import (
	"fmt"
	"piscine"
	"strconv"
)

func main() {
	testCases := []string{"++1234", "", "+1234", "-1234", "00123", "+01234"}
	fmt.Println("Testing atoi")
	fmt.Println("atoi(\"0234\"):",
		piscine.Atoi("0234"), "\n")

	for _, str := range testCases {
		fmt.Printf("atoi(\"% 15v\"):%15v", str, piscine.Atoi(str))
		strConvAtoi, _ := strconv.Atoi(str)
		fmt.Printf("-- strconv.atoi(\"% 15v\"):%15v", str, strConvAtoi)
		fmt.Printf("-- picine.AtoiAgainAgain(\"% 15v\"):%15v\n", str, piscine.AtoiAgainAgain(str))
	}
}
