package piscine

func isCharset(char byte, charset string) bool {
	for i := 0; i < len(charset); i++ {
		if char == charset[i] {
			return true
		}
	}
	return false
}

func Split(str, charset string) []string {
	var result []string
	var word string
	for i := 0; i < len(str); i++ {
		if !isCharset(str[i], charset) {
			word += string(str[i])
		} else {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
