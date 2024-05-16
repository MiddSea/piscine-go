package piscine

func ToLower(s string) string {
	var newString string
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			newString += string(s[i] - 'A' + 'a')
		} else {
			newString += string(s[i])
		}
	}
	return newString
}

// anclarma-rep
// func ToLower(s string) string {
// 	x := map[rune]rune{'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i', 'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r', 'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z'}
// 	var result []rune
// 	for _, r := range s {
// 		lower, ok := x[r]
// 		if ok {
// 			result = append(result, lower)
// 		} else {
// 			result = append(result, r)
// 		}
// 	}
// 	return string(result)
// }

// this is the reversed map for toLower
