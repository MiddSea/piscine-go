package piscine

import "fmt"

func QuadA(x, y int) {
	fmt.Println("x, y QuadA")

	//top line
	if x >= 1 && y >= 1 {
		print("o") // top left corner
		if x > 2 {
			for width := 2; width < x; width++ {
				print("-") // top line
			}
		}
		if x > 3 {
			print("o") // last
		}
	}

	//lines according to height

	//last line
}
