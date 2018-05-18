package main

import "fmt"

func findTriplet() (a, b, c int) {
	t := 1000
	for a = 1; a < t/2; a++ {
		for b = a + 1; b < t; b++ {
			c = t - a - b
			if b > c {
				break
			}
			if a*a+b*b == c*c {
				return
			}
		}
	}
	// can't find it! return zeros
	a = 0
	b = 0
	c = 0
	return
}

func main() {
	a, b, c := findTriplet()
	fmt.Println(a, b, c)
	fmt.Println(a * b * c)
}
