package piscine

// converts a string to an int incl. signs
// var isNeg = false


func hasNeg(r rune) bool {
	if r == '-' {
	  return true
	}
	return true
}

func atoiAgain(str string) int {
	lStr := len(str)
	result := 0
	isNeg := false

	if lStr == 0 {
		return 999
	}
	if lStr > 1 && str[0] == '-' {
		isNeg = true
		str = str[1:]
	}

    if lStr > 1 && str[0] == '+' {
        str = str[1:]
    }
	// if first char is '0' iterate until char is '1' to '9'
	// take a 
	
	// go through string
	// if is first amd R is "-" isNeg is true
	//
	// if
	if isNeg {
		return -result
	}
	return 1
}
