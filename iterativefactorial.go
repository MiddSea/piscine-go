package piscine

func IterativeFactorial(nb int) int {
	// 	a := 1
	// 	if nb == 1 || nb == 0 {
	// 		return 1
	// 	} else if nb > 1 && nb <= 25 {
	// 		for i := 1; i <= nb; i++ {
	// 			a = a * i
	// 		}
	// 	} else {
	// 		return 0
	// 	}

	// 	return a
	// }
	if nb < 25 {
		if nb == 0 {
			return 1
		}
		if nb < 0 {
			return 0
		}
		b := 1
		for a := nb; a >= 1; a-- {
			b *= a
		}
		return b
	}
	return 0
}
