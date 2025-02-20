package solutions

func MaxSubArray(nums []int) int {
	// f(i)=max{f(i−1)+nums[i],nums[i]}
	// 如果f(i-1)+nums[i]的和 < nums[i] 则丢弃前面的值，否则取f(i−1)+nums[i]
	// 把累计值存到nums[i]中
	if len(nums) == 0 {
		return 0
	}

	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i]+nums[i-1], nums[i])
		maxValue = max(maxValue, nums[i])
	}
	return maxValue
}
