package piscine

// func StrLen(str string) int {
// 	c := 0
// 	for range str {
// 		c++
// 	}
// 	return c
// }

func StrLen(s string) int {
	return len([]rune(s))
}
