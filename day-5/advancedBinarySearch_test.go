package day5

import "testing"

func TestFindMinRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "rotated at middle",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "rotated at end",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "not rotated",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "two elements rotated",
			nums:     []int{2, 1},
			expected: 1,
		},
		{
			name:     "two elements not rotated",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "rotated by one",
			nums:     []int{2, 3, 4, 5, 6, 7, 1},
			expected: 1,
		},
		{
			name:     "larger array rotated",
			nums:     []int{11, 13, 15, 17, 1, 3, 5, 7, 9},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMinRotatedSortedArray(tt.nums)
			if result != tt.expected {
				t.Errorf("FindMinRotatedSortedArray(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "target in left sorted portion",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			name:     "target in right sorted portion",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   1,
			expected: 5,
		},
		{
			name:     "target at rotation point",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "target at beginning",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   4,
			expected: 0,
		},
		{
			name:     "target at end",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   2,
			expected: 6,
		},
		{
			name:     "target not found",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "single element found",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "single element not found",
			nums:     []int{1},
			target:   2,
			expected: -1,
		},
		{
			name:     "not rotated array",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "two elements first",
			nums:     []int{3, 1},
			target:   3,
			expected: 0,
		},
		{
			name:     "two elements second",
			nums:     []int{3, 1},
			target:   1,
			expected: 1,
		},
		{
			name:     "larger array target in left",
			nums:     []int{11, 13, 15, 17, 1, 3, 5, 7, 9},
			target:   13,
			expected: 1,
		},
		{
			name:     "larger array target in right",
			nums:     []int{11, 13, 15, 17, 1, 3, 5, 7, 9},
			target:   7,
			expected: 7,
		},
		{
			name:     "target less than all elements",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   -1,
			expected: -1,
		},
		{
			name:     "target greater than all elements",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   10,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArray(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("SearchInRotatedSortedArray(%v, %d) = %d; want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
