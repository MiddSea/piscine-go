package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			return false
		}
	}
	return true
}
