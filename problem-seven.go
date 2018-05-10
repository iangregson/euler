package main

import "fmt"

// send a sequence of integers from 2 onwards to a channel 'ch'
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // send i into the channel
	}
}

// take the primes from the incoming channel (that were generated above)
// and push them out to a new channel
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // receive int from input channel
		if i%prime != 0 {
			out <- i // send the integer to the output channel
		}
	}
}

// The prime seive: Diasy-chain Filter process
func main() {
	intChannel := make(chan int) // create a new channel
	go Generate(intChannel)      // Launch the Generate goroutine to start generating our sequence of integers
	for i := 0; i < 10000; i++ {
		prime := <-intChannel
		primeChannel := make(chan int)
		go Filter(intChannel, primeChannel, prime)
		intChannel = primeChannel
	}
	prime := <-intChannel
	fmt.Println(prime)
}

// credit: https://github.com/JanLaussmann/Project-Euler-Golang/blob/master/Solution-7.go
