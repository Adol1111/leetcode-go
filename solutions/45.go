package solutions

func Jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxPosition := 0
	end := 0
	steps := 0

	// 每个位置都有N种选择，即从i的位置跳1到nums[i]步，即i+nums[i]
	// 为了让次数最少，我们每次都选择能跳的最远的位置
	// 由于不知道哪个位置能跳的更远，因此要把到最远位置之间的可能性都要测试一遍，如果达到了上一跳的最远位置（end位置），
	// 说明所有尝试都已验证过，则把步骤+1，然后更新end位置为，从上一跳开始到zend为止所有尝试中的最远位置即可，step+1即可

	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}
