package tests

import (
	"testing"

	twoPointers "github.com/ivivanov18/go-algo/two-pointers"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "basic case with solution",
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: true,
		},
		{
			name:     "no solution exists",
			nums:     []int{1, 2, 3, 4, 5},
			target:   15,
			expected: false,
		},
		{
			name:     "solution at boundaries",
			nums:     []int{1, 5, 10, 15, 20},
			target:   21,
			expected: true,
		},
		{
			name:     "negative numbers with solution",
			nums:     []int{-5, -2, 0, 3, 7},
			target:   2,
			expected: true,
		},
		{
			name:     "negative numbers no solution",
			nums:     []int{-5, -2, 0, 3, 7},
			target:   -10,
			expected: false,
		},
		{
			name:     "array with duplicates",
			nums:     []int{1, 2, 2, 3, 4},
			target:   4,
			expected: true,
		},
		{
			name:     "single element",
			nums:     []int{5},
			target:   5,
			expected: false,
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   10,
			expected: false,
		},
		{
			name:     "two elements with solution",
			nums:     []int{3, 7},
			target:   10,
			expected: true,
		},
		{
			name:     "two elements no solution",
			nums:     []int{3, 7},
			target:   5,
			expected: false,
		},
		{
			name:     "all same numbers with solution",
			nums:     []int{5, 5, 5, 5},
			target:   10,
			expected: true,
		},
		{
			name:     "all same numbers no solution",
			nums:     []int{5, 5, 5, 5},
			target:   7,
			expected: false,
		},
		{
			name:     "large numbers",
			nums:     []int{100, 200, 300, 400, 500},
			target:   600,
			expected: true,
		},
		{
			name:     "zero target",
			nums:     []int{-5, -3, 0, 3, 5},
			target:   0,
			expected: true,
		},
		{
			name:     "target with zero in array",
			nums:     []int{-5, 0, 5, 10},
			target:   5,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoPointers.TwoSum(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("TwoSum(%v, %d) = %v; expected %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func TestTwoSumEdgeCases(t *testing.T) {
	t.Run("minimum possible sum", func(t *testing.T) {
		nums := []int{-1000, -500, 0, 500, 1000}
		target := -1500
		result := twoPointers.TwoSum(nums, target)
		expected := true
		if result != expected {
			t.Errorf("TwoSum(%v, %d) = %v; expected %v", nums, target, result, expected)
		}
	})

	t.Run("maximum possible sum", func(t *testing.T) {
		nums := []int{-1000, -500, 0, 500, 1000}
		target := 1500
		result := twoPointers.TwoSum(nums, target)
		expected := true
		if result != expected {
			t.Errorf("TwoSum(%v, %d) = %v; expected %v", nums, target, result, expected)
		}
	})

	t.Run("sorted array ascending", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		target := 11
		result := twoPointers.TwoSum(nums, target)
		expected := true
		if result != expected {
			t.Errorf("TwoSum(%v, %d) = %v; expected %v", nums, target, result, expected)
		}
	})
}
