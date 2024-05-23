package piscine

func StrRev(s string) string {
	str := []rune(s)
	lenStr := StrLen(s)
	for i := 0; i < lenStr/2; i++ {
		tmp := str[i]
		str[i] = str[lenStr-i-1]
		str[lenStr-i-1] = tmp
	}
	return string(str)
}
