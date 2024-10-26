package piscine

func FishAndChips(n int) string {
	ret := ""
	if n < 0 {
		return "error: number is negative"
	}
	if n%2 != 0 && n%3 != 0	{
		return "error: non divisible"
	}
	if n%2 == 0 {
		ret += "fish"
		if n%3 == 0 {
			ret += " and "
		}
	}
	if n%3 == 0 {
		ret += "chips"
	} 
	return ret
}
