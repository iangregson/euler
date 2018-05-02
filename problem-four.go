package main

import "strconv"
import "fmt"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	r := Reverse(s)
	return s == r
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	largest_palindrome := 0

	for a := 100; a < 1000; a++ {
		for b := a + 1; b < 1001; b++ {
			product := a * b
			product_is_palindrome := isPalindrome(product)
			if product > largest_palindrome && product_is_palindrome {
				largest_palindrome = product
			}
		}
	}
	fmt.Println(largest_palindrome)
}
