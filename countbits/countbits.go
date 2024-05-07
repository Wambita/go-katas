package main

import "strconv"

func CountBits(u uint) int {
	// var to store the binary equivalent of u
	base_str := strconv.FormatUint(uint64(u), 2)

	// counter to store the count of 1
	count := 0
	for _, i := range base_str {
		if i == '1' {
			count++
		}
	}
	return count
}
