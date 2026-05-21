package main

import "github.com/01-edu/z01"

func main() {
	alpha := "zyxwvutsrqponmlkjihgfedcba"
	for _, i := range alpha {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

