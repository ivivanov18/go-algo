package twoPointers

func MaxArea(nums []int) int {
	left, right := 0, len(nums)-1
	currentMax := 0

	for left < right {
		width := right - left
		height := min(left, right)
		currentArea := height * width

		if currentMax < currentArea {
			currentMax = currentArea
		}

		if nums[left] < nums[right] {
			left++
		} else {
			right++
		}
	}

	return currentMax
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
