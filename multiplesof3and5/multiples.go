package main

import "fmt"

func Multiple3And5(number int) int {
	// declare var to store sum
	sum := 0
	// loop through the number in descnding order checking if it is a multiple of 3/5
	// add to sum if true
	for i := number - 1; i > 0; i-- {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	fmt.Println(sum)
	return sum
}

func main() {
	number := 30
	Multiple3And5(number)
}
