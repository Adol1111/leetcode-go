package solutions

func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	// answer[i-1] = nums[0] * nums[1] * ... * nums[i-2]
	// 先把左边的乘积算出来，然后再乘以右边的乘积，这样就可以不用除法
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// 倒过来，从右往左乘
	// R[i] = 1 * nums[len-1] * ... * nums[i+1]
	// answer[i] = answer[i](左边的乘积) * R[i]
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}
	return answer
}
