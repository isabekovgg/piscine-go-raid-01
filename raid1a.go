package main

import "github.com/01-edu/z01"

func Raid1a(x, y int) {

	if x > 0 && y > 0 {
		//draw top width
		z01.PrintRune('o')
		for a := 0; a < x-2; a++ {
			z01.PrintRune('-')
		}
		if x > 1 {
			z01.PrintRune('o')
		}
		z01.PrintRune('\n')
		//draw length
		for b := 0; b < y-2; b++ {
			z01.PrintRune('|')
			for c := 0; c < x-2; c++ {
				z01.PrintRune(' ')
			}
			if x > 1 {
				z01.PrintRune('|')
			}
			z01.PrintRune('\n')
		}
		//draw bottom width
		if y > 1 {
			z01.PrintRune('o')
			for a := 0; a < x-2; a++ {
				z01.PrintRune('-')
			}
			if x > 1 {
				z01.PrintRune('o')
			}
			z01.PrintRune('\n')
		}
	} else {
		z01.PrintRune('0')
	}
}

// func main() {
// 	Raid1a(5,1)
// }
