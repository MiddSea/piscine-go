package piscine

import "fmt"

func QuadA(x, y int) {
	fmt.Println("x, y QuadA")

	// top line
	if x >= 1 && y >= 1 { // x min 1 o
		print("o") // top left corner
		if x > 2 { // x min 3 o-o for dashes
			for width := 2; width < x; width++ {
				print("-") // '-' dashes top line
			}
		}
		if x > 1 { // x min 2: oo
			print("o\n") // last 'o' top line
		}
	}

	// lines according to height
	if y > 2 {
		for height := 2; height < y; height++ {
			if x >= 1 && y >= 2 { // x min 1 o
				print("|") // left column
				if x > 2 { // x min 3 | | for spaces
					for width := 2; width < x; width++ {
						print(" ") // ' ' space in lines
					}
				}
				if x > 1 { // x min 2: ||
					print("|\n") // last '|' in line
				}
			}
		}
	}

	// bottom line
	if x >= 1 && y >= 2 { // x min 1 o
		print("o") // bottom left corner
		if x > 2 { // x min 3 o-o for dashes
			for width := 2; width < x; width++ {
				print("-") // '-' dashes bottom line
			}
		}
		if x > 1 { // x min 2: oo
			print("o\n") // last 'o' bottom line
		}
	}
}
