package piscine

const up int = 1
const down int = -1

var downUp = up

func FromTo(from int, to int) string {
	ret := ""
	if (from < 0 && to > 99) || (from > 99 && to > 0) {
		ret = "Invalid"
	} else {
		if from > to &&  downUp == up  {
			downUp = down
			to = -to
		}
		//trap := 0
		for n := from; (n <= to && downUp == up) || (n >= to && downUp == down); n += downUp {
			 print(">n", n, "du", downUp, "<")
			if n < 10 {
				ret += "0"
				ret += string('0' + rune(n))
			}  
			if n >= 10 {
				ret += string('0' + rune(n/10)) //
				ret += string('0' + rune(n%10)) //
			}
			if n < to && downUp == up || n > to && downUp == down {
				ret += ", "
			}
			/*if trap >= 20 {
				print("PPANIC")
				return "PPANIC return"
			} else { 
			trap++*/
			
		}
	}
	ret += "\n"
	return ret
}
