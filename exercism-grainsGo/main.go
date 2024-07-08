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
	number, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("Total number of grains on square %d is: %d\n", number, square(float64(number)))
}

func square(number float64) int {
	if number < 1 || number > 64 {
		return 0
	}
	sumOfGrains := 1.0
	oneSquare := 1.0
	for i := 0; i < int(number); i++ {
		sumOfGrains += oneSquare
	}
	return int(sumOfGrains) << int(number-1)
}
