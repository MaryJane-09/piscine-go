package main

import "github.com/01-edu/z01"

// func main() {
// 	num := "0123456789"
// 	for _, i := range num {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }

func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}