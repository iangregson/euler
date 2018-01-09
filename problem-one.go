package main

import "fmt"

func main() {
	fmt.Println("Going to attempt problem 1 now...")

	var numbers []int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			numbers = append(numbers, i)
		}
	}

	var sum int

	for _, number := range numbers {
		sum += number
	}

	fmt.Println(sum)
}
