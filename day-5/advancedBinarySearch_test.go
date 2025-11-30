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

func TestFindPeakArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "single element",
			nums: []int{5},
		},
		{
			name: "two elements ascending",
			nums: []int{1, 2},
		},
		{
			name: "two elements descending",
			nums: []int{2, 1},
		},
		{
			name: "peak at beginning",
			nums: []int{5, 4, 3, 2, 1},
		},
		{
			name: "peak at end",
			nums: []int{1, 2, 3, 4, 5},
		},
		{
			name: "peak in middle",
			nums: []int{1, 2, 3, 2, 1},
		},
		{
			name: "peak at index 1",
			nums: []int{1, 3, 2},
		},
		{
			name: "peak near beginning",
			nums: []int{1, 5, 3, 2},
		},
		{
			name: "peak near end",
			nums: []int{1, 2, 5, 3},
		},
		{
			name: "multiple peaks left side",
			nums: []int{1, 3, 2, 5, 4},
		},
		{
			name: "multiple peaks right side",
			nums: []int{1, 3, 2, 4, 6},
		},
		{
			name: "valley then peak",
			nums: []int{5, 1, 2, 3, 4},
		},
		{
			name: "peak then valley",
			nums: []int{1, 5, 4, 3, 2},
		},
		{
			name: "zigzag pattern",
			nums: []int{1, 3, 2, 4, 3, 5, 4},
		},
		{
			name: "larger ascending array",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "larger descending array",
			nums: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "peak in middle of large array",
			nums: []int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6},
		},
		{
			name: "early peak in large array",
			nums: []int{1, 8, 7, 6, 5, 4, 3, 2},
		},
		{
			name: "late peak in large array",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 15, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindPeakArray(tt.nums)

			// Validate that result is within bounds
			if result < 0 || result >= len(tt.nums) {
				t.Errorf("FindPeakArray(%v) = %d; index out of bounds", tt.nums, result)
				return
			}

			// Validate that the returned index is actually a peak
			isPeak := true

			// Check left neighbor
			if result > 0 && tt.nums[result] <= tt.nums[result-1] {
				isPeak = false
			}

			// Check right neighbor
			if result < len(tt.nums)-1 && tt.nums[result] <= tt.nums[result+1] {
				isPeak = false
			}

			if !isPeak {
				t.Errorf("FindPeakArray(%v) = %d (value=%d); not a peak element",
					tt.nums, result, tt.nums[result])
			}
		})
	}
}
