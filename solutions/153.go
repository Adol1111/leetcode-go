package solutions

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	for left <= right {
		mid := (left + right) / 2
		// 最小值在无序的一侧, 如果两边都有序则在nums[0]
		if nums[0] <= nums[mid] {
			left = mid + 1
		} else if nums[mid] <= nums[len(nums)-1] {
			right = mid - 1
		}
	}
	return nums[left]
}
