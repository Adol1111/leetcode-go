package solutions

func MinSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLen := len(nums) + 1
	for i, j, sum := 0, 0, 0; j < len(nums); j++ {
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
