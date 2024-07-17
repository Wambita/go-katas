package main

import (
	"fmt"
)

func ordinal(v int) string {
	// exceptions
	if v == 11 || v == 12 || v == 13 {
		return "th"
	}

	v %= 10

	switch v {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func main() {
	for c := 1; c <= 20; c++ {
		fmt.Printf("%d%s\t%d%s\t%d%s\t%d%s\t%d%s\n",
			c, ordinal(c),
			c+20, ordinal(c+20),
			c+40, ordinal(c+40),
			c+60, ordinal(c+60),
			c+80, ordinal(c+80),
		)
	}
}
