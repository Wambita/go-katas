package main

import "testing"

func TestFindShort(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Basic", "Let's travel abroad shall we", 2},
		{"Mixed Case", "Hello world Golang", 5},
		{"Single Word", "hello", 5},
		{"Multiple Spaces", "  this  is  a test   ", 1},
		{"No Spaces", "spaces", 6},
		{"Numerics Included", "hi 1234 56", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FindShort(tc.input)
			if actual != tc.expected {
				t.Errorf("Failed %s: Expected %d, got %d for input '%s'", tc.name, tc.expected, actual, tc.input)
			}
		})
	}
}
