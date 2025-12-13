package twoPointers

func TwoSum(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[right] + nums[left]

		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}
