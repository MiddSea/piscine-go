package piscine

// finds the next prime
func FindPrevPrime(nb int) int {
	a := nb
	for !IsPrime(a) {
		a--
	}
	return a
}
