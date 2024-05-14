package piscine

// https://gitea.anclarma.fr/anclarma/piscine-go/src/branch/main/piscine-go/nrune.go
// anclarma /piscine-go

func NRune(s string, n int) rune {
	if n <= StrLen(s) && n > 0 {
		return []rune(s)[n-1]
	} else {
		return 0
	}
}
