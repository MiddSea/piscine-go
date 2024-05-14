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

	var result int
	result = 1

	if nb == 0 {
		return 1
	}
	// negative number
	if nb < 0 {
		return 0
	}
	// do factorial iterative loop
	for a := nb; a >= 1; a-- {
		result = a * result
	}
	// check for overflow i.e. negative number
	if result > 1 {
		return result // positive result
	} else {
			return 0 // neg number / overflow
	}

}
