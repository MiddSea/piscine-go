package piscine

// FirstRune does what it said on a package
func NRune(s string) rune {
	runeSlice := []rune(s)
	lastRuneNo := len(runeSlice) - 1
	r := runeSlice[lastRuneNo]
	return r
}

// done by Seán Middleton 14 May
nrune.go