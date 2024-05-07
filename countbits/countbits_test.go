package main

import "testing"

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		input    uint
		expected int
	}{
		{"Zero", 0, 0},
		{"One", 1, 1},
		{"Two", 2, 1},
		{"Three", 3, 2},
		{"Fifteen", 15, 4},
		{"Sixteen", 16, 1},
		{"SeventyEight", 78, 4},
		{"MaxUint32", 4294967295, 32}, // Max for uint32
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := CountBits(test.input)
			if output != test.expected {
				t.Errorf("Failed %s: /nExpected %d /nGot %d", test.name, test.expected, output)
			}
		})
	}
}
