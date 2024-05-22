package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	xStr := "x = "
	yStr := "y = "
	points := &point{}

	setPoint(points)

	xmassiv := []rune{}
	ymassiv := []rune{}

	xVal := points.x
	yVal := points.y

	for xVal != 0 {
		xmassiv = append(xmassiv, rune(xVal%10))
		xVal /= 10
	}
	for _, val := range xStr {
		z01.PrintRune(rune(val))
	}
	for i := len(xmassiv) - 1; i >= 0; i-- {
		z01.PrintRune(xmassiv[i] + 48)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')

	for yVal != 0 {
		ymassiv = append(ymassiv, rune(yVal%10))
		yVal /= 10
	}
	for _, val := range yStr {
		z01.PrintRune(rune(val))
	}
	for i := len(ymassiv) - 1; i >= 0; i-- {
		z01.PrintRune(ymassiv[i] + 48)
	}
	z01.PrintRune('\n')
}

// package main

// import (
// 	"github.com/01-edu/"

// )

// // Create a new directory called point.

// // The code below must be copied into a file called main.go inside the point directory.

// // The necessary changes must be applied so that the program works.

// // The function setPoint() must work with int.

// // Code to be copied
// // func setPoint(ptr *point) {
// // 	ptr.x = 42
// // 	ptr.y = 21
// // }

// // func main() {
// // 	points := &point{}

// // 	setPoint(points)

// // 	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
// // }
// // Usage
// // $ go run .
// // x = 42, y = 21
// // $

// // Assuming the definition of the point struct

// type point struct {
// 	x int
// 	y int
// }

// func setPoint(ptr *point) {
// 	ptr.x = 42
// 	ptr.y = 21
// }

// func Int2String(num int) string {
// 	i := 0
// 	if num == 0 {
// 		return "0"
// 	}
// 	for j := 1; j <= num%mod; j++ {
// 		i++
// 	}
// 	for j := -1; j >= num%mod; j-- {
// 		i++
// 	}
// 	return NumFromBase(num/mod, base) + string(base[i])
// }

// printRuneStr(s string) {
// 	for _, r := range s {
// 		z0
// 	}

// 	}

// }

// func main() {
// 	points := &point{}

// 	setPoint(points)

// 	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
// }
