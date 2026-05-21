package main

import "github.com/01-edu/z01"

// func main() {
// 	alpha := "abcdefghijklmnopqrstuvwxyz"
// 	for _, i := range alpha {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}