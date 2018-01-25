package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func main() {
	// get fibonacci sequence up to 4 million
	f := fibonacci()
	var last_number int = 0
	var fib_sum int = 0

	for last_number < 4e6 {
		last_number = f()
		if last_number%2 == 0 {
			fib_sum += last_number
		}
	}

	fmt.Println(fib_sum)
}
