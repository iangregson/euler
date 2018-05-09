package main

import (
	"fmt"
	"math"
)

func main() {
	sumsquares := sumSquares(100)
	squaresums := squareSum(100)
	result := int(sumsquares - squaresums)
	fmt.Println(result)
}

func sumSquares(l float64) float64 {
	sum := float64(0)
	for i := float64(1); i <= l; i++ {
		sum += math.Pow(i, 2)
	}
	return sum
}
func squareSum(l float64) float64 {
	sum := float64(0)
	for i := float64(1); i <= l; i++ {
		sum += i
	}
	return math.Pow(sum, 2)
}
