package piscine

// func FindNextPrime(nb int) int {
// 	ret := nb
// 	if ret <= 0 {
// 		ret = 2
// 	}
// 	for !IsPrime(ret) {
// 		ret++
// 	}
// 	return ret
// }



// https://github.com/4ubak/piscine-go/blob/master/findnextprime.go 
// frm oct 2019
// 4ubak 4ubak Update findnextprime.go

// 
func FindNextPrime(nb int) int {
	for a := nb; a >= nb; a++ {
		if IsPrime(a) {
			return a
		}
	}
	return 1
}
