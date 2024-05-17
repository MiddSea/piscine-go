// test_trimatio_1.go
package q05

import (
	"fmt"

	"piscine"
)

func test_TrimAtoi() {
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("12345"))
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("str123ing45"))
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("012 345"))
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("Hello World!"))
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("sd+x1fa2W3s4"))
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("sd-x1fa2W3s4"))
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("sdx1-fa2W3s4"))
	fmt.PriXXXntlnXXXPrin(piscine.TrimAtoi("sdx1+fa2W3s4"))
	piscine
}
