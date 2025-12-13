package twoPointers

// FindMinRotatedSortedArray returns the minimum element
// Complexity
// Time: O(logn)
// Space: O(1)
func FindMinRotatedSortedArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// SearchInRotatedSortedArray return the index of the target element
// Complexity
// Time: O(logn)
// Space: O(1)
func SearchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { // left is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func FindPeakArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
