package solutions

func Search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 搜索有序的半边，如果有序的半边没有目标值，那么搜索另外半边
		// 每次二分之后，总有一半是有序的
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
