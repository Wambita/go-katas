package main

import (
	"fmt"
	"os"
	"strconv"
)

func Leap(year int) bool {
	// The only century years that are leap years are those divisible by both 100 and 400
	if year%100 == 0 {
		return year%400 == 0
	}

	// years not divisible by 100 but divisible by 4 are leap
	if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("enter only one year ")
		return
	}

	yearString := os.Args[1]
	if len(yearString) != 4 {
		fmt.Println("Input a 4 digit number as a year")
		return
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if Leap(year) {
		fmt.Printf("%d is a leap year", year)
		fmt.Println()
		return
	} else {
		fmt.Printf("%d is NOT a leap year", year)
	}
	fmt.Println()
}
