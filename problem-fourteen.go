package main

import "fmt"

func getSequence(n int) (c int) {
	c = 1
	for {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		c++
		if n == 1 {
			return
		}
	}
}

func main() {
	var longestChain int = 0
	var startingNumber int = 0
	for n := 999999; n > 1; n-- {
		seqLen := getSequence(n)
		if seqLen > longestChain {
			longestChain = seqLen
			startingNumber = n
		}
	}
	fmt.Println(startingNumber)
}
