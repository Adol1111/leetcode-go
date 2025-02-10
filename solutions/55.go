package solutions

func CanJump(nums []int) bool {
	// i+nums[i] 表示从当前位置最远能跳到的位置
	// maxPos 表示当前能跳到的最远位置
	// 如果 i > maxPos，说明当前位置已经超过最远位置，否则更新 maxPos
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
