package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
	sum := big.NewInt(0)
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		input := stdin.Text()
		val, _ := new(big.Int).SetString(input, 10)
		sum = new(big.Int).Add(sum, val)
	}
	sum_string := sum.String()
	sum_rune := []rune(sum_string)
	ret_string := string(sum_rune[:10])
	ret_int, err := strconv.Atoi(ret_string)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret_int)
}
