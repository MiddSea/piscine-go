package piscine

// FirstRune does what it said on a package
func LastRune(s string) rune {
	runeSlice := []rune(s)
	lastRune := len(runeSlice) - 1
	r := runeSlice[lastRune]
	return r
}

// done by Seán Middleton 14 May
