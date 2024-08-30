package main

import "testing"

type Test struct {
	name     string
	input    int
	expected int
}

func TestMyFunc(t *testing.T) {
	tests := [...]Test{
		{"8", 8, 40320},
		{"1", 1, 1},
		{"4", 4, 24},
		{"0", 0, -1},
		{"-5", -5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := myFunc(tt.input)
			if result != tt.expected {
				t.Errorf("got %d, expected %d", result, tt.expected)
			}
		})
	}
}
