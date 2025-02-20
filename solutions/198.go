package solutions

func Rob(nums []int) int {
	// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	// 第i项的收益为 i-2 项的收益加上第i项的收益， 或者 i-1 项的收益（当前房间不偷） 取最大值
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
