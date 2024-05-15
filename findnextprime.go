package piscine

// finds the next prime
func FindNextPrime(nb int) int {
	a := nb
	for !IsPrime(a) {
		a++
	}
	return a
}
