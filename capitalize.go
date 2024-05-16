package piscine

// induced and Capital
func Capitalize(s string) string {
	var newString string
	isFirst := true
	for i := 0; i < len(s); i++ {
		if IsAlpha(string(s[i])) {
			if isFirst {
				newString += ToUpper(string(s[i]))
			} else {
				newString += ToLower(string(s[i]))
			}
			isFirst = false
		} else {
			newString += string(s[i])
			isFirst = true
		}
	}
	return newString
}

func IsAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

// this is the reversed map for toLower
func ToLower(s string) string {
	x := map[rune]rune{'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i', 'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r', 'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z'}
	var result []rune
	for _, r := range s {
		lower, ok := x[r]
		if ok {
			result = append(result, lower)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// this is the old one for ToUpper
func ToUpper(s string) string {
	x := map[rune]rune{'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D', 'e': 'E', 'f': 'F', 'g': 'G', 'h': 'H', 'i': 'I', 'j': 'J', 'k': 'K', 'l': 'L', 'm': 'M', 'n': 'N', 'o': 'O', 'p': 'P', 'q': 'Q', 'r': 'R', 's': 'S', 't': 'T', 'u': 'U', 'v': 'V', 'w': 'W', 'x': 'X', 'y': 'Y', 'z': 'Z'}
	var result []rune
	for _, r := range s {
		big, ok := x[r]
		if ok {
			result = append(result, big)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}



