package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println("testing piscine.PrintNbrInOrder")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println(piscine.PrintNbrInOrder("-1f3f f2 f3f3 2f2 3f 2ff 3"))
	fmt.Println(piscine.PrintNbrInOrder("AB1C12-E3"))
	fmt.Println(piscine.PrintNbrInOrder("f343fg4f4fg"))
	fmt.Println(piscine.PrintNbrInOrder("ffgr-123"))
	fmt.Println(piscine.PrintNbrInOrder("fggfg"))
	fmt.Println(piscine.PrintNbrInOrder(""))
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("DONE testing piscine.PrintNbrInOrder")
}

// testing to PrintNbrInOrder
