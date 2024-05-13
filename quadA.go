package piscine

import (
	zz "github.com/01-edu/z01"
)

func QuadA(x, y int) {
	// fmt.Prixxxntlnxxx("x, y QuadA")

	// top line
	if x >= 1 && y >= 1 { // x min 1 o
		zz.PrintRune('o') // top left corner
		if x > 2 {        // x min 3 o-o for dashes
			for width := 2; width < x; width++ {
				zz.PrintRune('-') // '-' dashes top  for the top line
			}
		}
		if x > 1 { // x min 2: oo
			zz.PrintRune('o')  // last 'o' top line
			zz.PrintRune('\n') // last 'o' top line
		}
	}

	// lines according to height
	if y > 2 {
		for height := 2; height < y; height++ {
			if x >= 1 && y >= 2 { // x min 1 o
				zz.PrintRune('|') // left column
				if x > 2 {        // x min 3 | | for spaces
					for width := 2; width < x; width++ {
						zz.PrintRune(' ') // ' ' space in lines
					}
				}
				if x > 1 { // x min 2: ||
					zz.PrintRune('|') // last '|' in line
				}
			}
		}
	}

	// bottom line
	if x >= 1 && y >= 2 { // x min 1 o
		zz.PrintRune('o') // bottom left corner
		if x > 2 {        // x min 3 o-o for dashes
			for width := 2; width < x; width++ {
				zz.PrintRune('-') // '-' dashes bottom line
			}
		}
		if x > 1 { // x min 2: oo
			zz.PrintRune('o')  // last 'o' bottom line
			zz.PrintRune('\n') // last 'o' bottom line

		}
	}
}
