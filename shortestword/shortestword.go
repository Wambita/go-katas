package main

import (
	"strings"
)

// Algorithm : 1. Store the string in slice of strings separated by spaces
// 2. Sort the slice in ascending order or length of individual slices
// 3. Return the lenght of the first element in slice
func FindShort(s string) int {
	slicestr := strings.Fields(s)

	for i := 0; i < len(slicestr); i++ {
		for j := 0; j < len(slicestr); j++ {
			if len(slicestr[i]) < len(slicestr[j]) {
				slicestr[i], slicestr[j] = slicestr[j], slicestr[i]
			}
		}
	}c
	return len(slicestr[0])
}
