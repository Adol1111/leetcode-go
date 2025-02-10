package solutions

func Jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxPosition := 0
	end := 0
	steps := 0

	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}
