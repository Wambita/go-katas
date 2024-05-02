package main

import (
	"strings"
)
//main function
func main() {
	//test input
	sentence := "wha5t a1 b2d d3y m8net"
	Order(sentence)
}

// Order takes a string sentence and returns a new string with the words ordered
// according to embedded numerical values.
func Order(sentence string) string {
	words := strings.Split(sentence, " ")
	order := make(map[int]string)
	for _, word := range words {
		for _, char := range word {
			if char <= '9' && char > '0' {
				num := int(char - '0')
				order[num] = word
			}
		}
	}
 //slice to hold the results , length is the number of entries in the map
	results := make([]string, len(order))

	// iterate over map to populate the slice
	//adjust the map for zero based index. map was 1 based as the smallest digit was 1  not 0
	for k, v := range order {
		if k-1 < len(results) {
			results[k-1] = v
		}
	}
	sentence = strings.Join(results, " ")
	return sentence
}
