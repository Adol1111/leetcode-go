package solutions

func MinSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLen := len(nums) + 1
	for i, j, sum := 0, 0, 0; j < len(nums); j++ {
		// 先固定左指针，然后右指针向右移动，直到找到满足条件的子数组
		// 如果右指针找到了满足条件的子数组，那么左指针向右移动，同时减少sum的值（因为移除了一个左边的元素），直到不满足条件
		// 比较当前最小长度 和 j-i+1 的长度，取最小值
		// j-i+1为当前子数组的长度
		sum += nums[j]
		for sum >= s {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
