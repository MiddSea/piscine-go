package piscine

// converts a string to an int incl. signs
var isNeg = false


func isNeg(r rune) bool {
	if r == '-' {
		isNeg = true
	}
	return isNeg
}

func atoiAgain(str string) int {
	lStr := len(str)
	result := 0
	isNeg := false
	if lStr == 0 {
		return 999
	}
	// go through string
	// if is first amd R is "-" isNeg is true
	//
	// if
	if isNeg {
		return -result
	}
	return 1
}
