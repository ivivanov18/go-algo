package twoPointers

func MoveZeroes(nums []int) {
	nextNonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nextNonZero], nums[i] = nums[i], nums[nextNonZero]
			nextNonZero++
		}
	}
}
