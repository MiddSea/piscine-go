package piscine

// finds the next prime
func FindNextPrime(nb int) int {
	a := nb
	for IsPrime(a) == false {
		a++
	}
	return a
}
