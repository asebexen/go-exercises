package main

import (
	"reflect"
	"testing"
)

type Test struct {
	name        string
	input       string
	expected    []int
	expectedErr bool
}

func TestCreateIntArray(t *testing.T) {
	tests := [...]Test{
		{"Simple", "9,2,4", []int{9, 2, 4}, false},
		{"Whitespace", " 9  ,   2,4  ", []int{9, 2, 4}, false},
		{"Negatives", "-9,5,-201,-0", []int{-9, 5, -201, 0}, false},
		{"Empty", "", []int{}, true},
		{"Strings", "9,nine,f", []int{}, true},
		{"One element", "420", []int{420}, false},
		{"Empty element", "9,4,,2,5", []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := createIntArray(tt.input)
			if err != nil && !tt.expectedErr {
				t.Error("got error, expected none")
				return
			} else if err == nil && tt.expectedErr {
				t.Error("expected error, got none")
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
