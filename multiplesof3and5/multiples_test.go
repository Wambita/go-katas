package main

import "testing"

func TestMultiple3And5(t *testing.T) {
	tests := []struct {
		number   int
		expected int
	}{
		{10, 23},
		{20, 78},
		{30, 195},
	}

	for _, test := range tests {
		output := Multiple3And5(test.number)
		if output != test.expected {
			t.Errorf("Error: Multple3And5(%d), \nExpected:%d \nGot:%d", test.number, test.expected, output)
		}
	}
}
