package solutions

func SearchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid
			right = mid
			for left >= 0 && nums[left] == target {
				left--
			}
			for right < len(nums) && nums[right] == target {
				right++
			}
			return []int{left + 1, right - 1}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
