package piscine

func isAlpha(r rune) bool {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return true
	}
	return false
}

func CountAlpha(s string) int {
	count := 0
	for _, r := range s {
		if isAlpha(r) {
			count++
		}
	}
	return count
}
