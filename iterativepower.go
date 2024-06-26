package piscine

// q04 excercise
func IterativePower(nb, power int) int {
	if power < 0 {
		return 0
	}
	ret := 1
	for ; power > 0; power-- {
		ret *= nb
	}
	return ret
}
