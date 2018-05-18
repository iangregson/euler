package main

import "fmt"

func SieveOfEratosthenes(n int64) []int64 {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, n+1)
	for i := int64(2); i < n+1; i++ {
		integers[i] = true
	}
	for p := int64(2); p*p <= n; p++ {
		// It integers[p] is not changed, it is prime
		if integers[p] == true {
			// Update all the multiples
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int64
	for p := int64(2); p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}

func main() {
	primes := SieveOfEratosthenes(2e6)
	fmt.Println(len(primes))
	sum := int64(0)
	for _, prime := range primes {
		sum1 := sum + prime
		if sum1 < sum {
			fmt.Println(sum)
		} else {
			sum = sum1
		}
	}
	fmt.Println(sum)
}