package twoPointers

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "basic case with zeros in middle",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "no zeros",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "zeros at the beginning",
			input:    []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "zeros at the end",
			input:    []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "single element with zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "single element without zero",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "alternating zeros and non-zeros",
			input:    []int{0, 1, 0, 2, 0, 3},
			expected: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name:     "multiple consecutive zeros",
			input:    []int{1, 0, 0, 0, 2, 3},
			expected: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name:     "negative numbers with zeros",
			input:    []int{-1, 0, -2, 0, 3},
			expected: []int{-1, -2, 3, 0, 0},
		},
		{
			name:     "large numbers",
			input:    []int{0, 1000000, 0, -1000000},
			expected: []int{1000000, -1000000, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("MoveZeroes() = %v, want %v", tt.input, tt.expected)
			}
		})
	}
}
