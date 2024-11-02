package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println("Testing RepeatAlpha")
	fmt.Println("RepeatAlpha(\"abc\":", piscine.RepeatAlpha("abc"), "\n")
	fmt.Println("RepeatAlpha(\"AbcAd\":", piscine.RepeatAlpha("AbcAd"), "\n")
	fmt.Println("RepeatAlpha(\"\":", piscine.RepeatAlpha(""), "\n")
	fmt.Println("RepeatAlpha(\"abZca\":", piscine.RepeatAlpha("abZca"), "\n")
	fmt.Println("RepeatAlpha(\"azBa\":", piscine.RepeatAlpha("azBa"), "\n")

}
