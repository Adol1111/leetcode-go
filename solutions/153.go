package solutions

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	// 如果数组有序则直接返回
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	for left <= right {
		mid := (left + right) / 2
		// 观察可知，最小值在无序的一侧，左侧有序则移动左侧，右侧无序则移动右侧
		if nums[0] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}
