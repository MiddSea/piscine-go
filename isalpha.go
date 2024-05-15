package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(('a' >= s[i] && s[i] <= 'z') || ('A' >= s[i] && s[i] <= 'Z')) {
			return false
		}
	}
	return true
}
