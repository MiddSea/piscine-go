// test_trimatio_1.go
package q05

import (
	"fmt"
	"piscine"
)

func test_TrimAtoi() {
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("12345"))
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("str123ing45"))
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("012 345"))
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("Hello World!"))
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("sd+x1fa2W3s4"))
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("sd-x1fa2W3s4"))
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("sdx1-fa2W3s4"))
	fmt.PrintlnXXXPrin(piscine.TrimAtoi("sdx1+fa2W3s4"))
	piscine
}
