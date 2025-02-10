package solutions

func CanJump(nums []int) bool {
	n := len(nums)
	maxPos := 0
	for i := 0; i < n; i++ {
		if i > maxPos {
			return false
		}
		maxPos = max(maxPos, i+nums[i])
	}
	return true
}
