package main

import "fmt"
import "os"
import "strconv"
import "math"

func isPrime(x int64) bool {
	var i int64 = 2
	for ; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	input_number, _ := strconv.Atoi(os.Args[1])
	var N int64 = int64(input_number)
	var i int64 = int64(math.Sqrt(float64(N)))
	for ; i > 1; i-- {
		if N%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
