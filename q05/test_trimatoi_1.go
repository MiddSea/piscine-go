// test_trimatio_1.go
package q05

import (
	"fmt"
	"piscine"
)

func test_TrimAtoi() {
	fmt.Println(piscine.TrimAtoi("12345"))
	fmt.Println(piscine.TrimAtoi("str123ing45"))
	fmt.Println(piscine.TrimAtoi("012 345"))
	fmt.Println(piscine.TrimAtoi("Hello World!"))
	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
	piscine
}
