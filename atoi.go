package piscine

import "piscine"

func Atoi(s string) int {
	isNeg := false
	result := 0
	index := 0
	len := piscine.StrLen(s) // added piscine

	if index < len && (s[index] == '+' || s[index] == '-') {
		if s[index] == '-' {
			isNeg = true
		}
		index++
	}
	for index < len && s[index] == '0' {
		index++
	}
	for index < len && s[index] >= '0' && s[index] <= '9' {
		result *= 10
		result += int(s[index] - '0')
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
