package main

import "fmt"

func main() {
	var gridSize int64 = 20
	var paths int64 = 1
	for i := int64(0); i < gridSize; i++ {
		paths *= (2 * gridSize) - i
		paths /= i + 1
	}
	fmt.Println(paths)
}
