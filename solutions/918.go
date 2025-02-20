package solutions

func MaxSubarraySumCircular(nums []int) int {
	// 要使得环形和最大，可以反过来考虑，即要使得环形和最大，那么中间的部分必然是最小的
	// 即 f(i)=min{f(i−1)+nums[i],nums[i]}
	// 因此可以参考53题，求中间最大和 和 中间最小和，然后 max(中间最大和，总和-中间最小和)即可
	// 存在一个例外，即所有数字都是负数，此时最小和 = 全部数字之和，如果取总和-中间最小和，会得到0

	totalSum := nums[0]
	maxValue := nums[0]
	minValue := nums[0]
	preMax := nums[0]
	preMin := nums[0]

	for i := 1; i < len(nums); i++ {
		totalSum += nums[i]
		preMax = max(preMax+nums[i], nums[i])
		preMin = min(nums[i]+preMin, nums[i])
		minValue = min(minValue, preMin)
		maxValue = max(maxValue, preMax)
	}

	// 如果最大值和最小值相等，说明数组中全部都是负数，此时直接取最大值即可
	if totalSum == minValue {
		return maxValue
	}

	return max(totalSum-minValue, maxValue)
}

func MaxSubarraySumCircular2(nums []int) int {
	// leftMax[i]=max(leftMax[i−1],sum(nums[0:i+1])
	// 先求左边每个 f(i) 的最大值，然后由于是环形数组，所以可以从右到左求右边每个 f(i) 的最大值
	// 然后maxValue = max{maxValue, right[i：n]+leftMax(i-1)} 即可得到最大值
	// rightSum(i) = rightSum(i+1) + nums[i], 因此可以倒过来求
	// 左边的最大值可以参考 53 题，即 f(i)=max{f(i−1)+nums[i],nums[i]}，不过需要使用额外的存储存，避免改变nums的值

	n := len(nums)
	leftMax := make([]int, n)

	leftMax[0] = nums[0]
	leftSum, pre, maxValue := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		// 类似53题，只是要用pre保存结果，避免影响nums[i-1]的值
		pre = max(pre+nums[i], nums[i])
		// 此处的maxValue是 53中的情况，可以从任意位置开始，取最大值，此时结果不是环形的，而是连续排列的
		// 后续会和环形的结果比较，到底是环形大还是，直接53题的情况比较大
		maxValue = max(maxValue, pre)
		// 由于是环形数组，所以左侧必然是从0-i连续分布，故leftSum需要保持累加
		leftSum += nums[i]
		leftMax[i] = max(leftMax[i-1], leftSum)
	}

	// 从右到左枚举后缀，固定后缀，选择最大前缀
	rightSum := 0
	for i := n - 1; i > 0; i-- {
		rightSum += nums[i]
		maxValue = max(maxValue, rightSum+leftMax[i-1])
	}
	return maxValue
}
