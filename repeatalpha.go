package piscine

func RepeatAlpha(str string) string {
	res := ""
	for _, r := range str {
		// rep if A-Z
		if r >= 'A' && r <= 'Z' {
		// rep if A-Z
			for i := 0; i <= int(r-'A'); i++ {
				res += string(r)
			}
			
		} else if r >= 'a' && r <= 'z' {    
		// rep if a-z
			for i := 0; i <= int(r-'a'); i++ {
				res += string(r)
			}
		} else {
			res += string(r)
		}
	}
	return res
}

/*// rep if A-Z


import "unicode"
(s)
func RepeatAlpha(s string) string {
	res := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			rep := unicode.ToLower(r) - 'a' + 1
			for i := 0; i < int(rep); i++ {
				res += string(r)
			}
		} else {
			res += string(r)
		}
	}
	return res
}
*/
