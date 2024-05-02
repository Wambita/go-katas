package main

import "testing"

// test for the order function to see if the output matches the expected
func TestOrder(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"is2 Thi1s T4est 3a", "Thi1s is2 3a T4est"},
		{"4of Fo1r pe6ople g3ood th5e the2", "Fo1r the2 g3ood 4of th5e pe6ople"},
		{"", ""},
	}
	for _, test := range tests{
		output := Order(test.input)
		if output != test.expected {
			t.Errorf("Order(%s) expected %s, got %s", test.input, test.expected, output)
		}
	}
}
