package main

import (
	"strings"
)

func main() {
	sentence := "wha5t a1 b2d d3y m8net"
	Order(sentence)
}

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

	results := make([]string, len(order))
	for k, v := range order {
		if k-1 < len(results) {
			results[k-1] = v
		}
	}
	sentence = strings.Join(results, " ")
	return sentence
}
