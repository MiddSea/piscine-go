package piscine

func UltimateDivMod(a, b *int) {
	newA := *a / *b
	*b = *a % *b
	*a = newA
}
