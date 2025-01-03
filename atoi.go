package piscine

func Atoi(s string) int {
	isNeg := false
	result := 0
	index := 0
	sR := []rune(s)
	len := len(sR) // added piscine
	if index < len && (sR[index] == '+' || sR[index] == '-') {
		if sR[index] == '-' {
			isNeg = true
		}
		index++
	}
	if index < len && sR[index] == '0' && result == 0 {
		return 0
	}
	for index < len && sR[index] >= '0' && sR[index] <= '9' {
		result *= 10
			result += int(sR[index] - '0')
			index++
	}
	if isNeg {
		result *= -1
	}
	if index < len {
		return 0
	}
	return result
}
