package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if !('a' <= s[i] && s[i] <= 'z') {
			return false
		}
	}
	return true
}
