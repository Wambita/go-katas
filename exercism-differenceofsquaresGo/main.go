package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Difference of squares of %d is %d\n", num, DifferenceOfSquares(num))
}

func DifferenceOfSquares(n int) int {
	// square of sum
	sum := 0
	square := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	square = sum * sum
	// sum of squares
	sumSquare := 0
	for j := 1; j <= n; j++ {
		sumSquare += j * j
	}
	return square - sumSquare
}
