package piscine

func ConcatParams(args []string) string {
	var str string
	for i := 0; i < len(args); i++ {
		str += args[i]
		if i < len(args)-1 {
			str += "\n"
		}
	}
	return str
}