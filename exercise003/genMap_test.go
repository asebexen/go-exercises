package main

import (
	"reflect"
	"testing"
)

type Test struct {
	name     string
	input    int
	expected map[int]int
}

func TestGenMap(t *testing.T) {
	tests := [...]Test{
		{"4", 4, map[int]int{1: 1, 2: 4, 3: 9, 4: 16}},
		{"1", 1, map[int]int{1: 1}},
		{"-1", -1, map[int]int{}},
		{"0", 0, map[int]int{}},
		{"10", 10, map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36, 7: 49, 8: 64, 9: 81, 10: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := genMap(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, expected %v", result, tt.expected)
			}
		})
	}
}
