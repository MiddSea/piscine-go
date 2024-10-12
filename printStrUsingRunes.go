package piscine

// hand made by Seán
import "github.com/01-edu/z01"


func PrintStrUsingRunes(str string) {
for _, r := range str {
    z01.PrintRune(r)
}
}