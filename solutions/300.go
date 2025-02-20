package solutions

func LengthOfLIS(nums []int) int {
	// dp[i] = max(dp[j]) + 1, 其中 0 ≤ j < i 且 num[j] < num[i]
	dp := make([]int, len(nums))
	dp[0] = 1
	maxValue := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxValue = max(maxValue, dp[i])
	}
	return maxValue
}

func LengthOfLIS2(nums []int) int {
	// 贪心+二分查找
	// 原理，要让上升子序列尽可能的长，就要让上升子序列的末尾元素尽可能的小
	// 所以我们维护一个子数组d，d[i] 表示长度为 i 的最长上升子序列的末尾元素的最小值
	// 当遇到nums[i] 时，如果nums[i] > d[len-1]，那么直接加入到 d 的末尾即可，否则对数组的第j位进行更新，d[j] = nums[i]
	// 需要满足 d[j-1] < d[j] < d[j+1], 这里可以通过二分查找来实现(单调性可见下面的证明)
	//
	// 由定义可知，一定有 d[j] <= d[i] j < i，所以 d 是单调递增的
	// 证明（反证法）假设 d[j]>=d[i]，因为 j<i，所以从d[i]删除末尾i-j个元素，正好是d[j] (d[i]的定义是长度为i的上升子序列)
	// 所以d[j]<d[i], 与前提条件 d[j]>=d[i]矛盾
	//
	// 根据以上描述，我们可以这样做
	// 1. 如果 nums[i] 比 d[len-1] 大，那么直接加入到 d 的末尾即可
	// 2. 否则，用二分查找找出第一个比nums[i]小的d[j]，然后更新d[j+1] = nums[i]
	// 3. 最后返回 d 的长度即可
	//
	// 注意 d 并非实际的最长子序列，但是长度是一样的, 因为某些位置被后面的更小的元素替换了

	d := make([]int, len(nums))
	d[0] = nums[0]
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[length-1] {
			d[length] = nums[i]
			length++
			continue
		}

		// 二分查找
		left, right := 0, length-1
		pos := right
		for left <= right {
			mid := (left + right) / 2
			if d[mid] >= nums[i] {
				pos = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		d[pos] = nums[i]
	}
	return length
}
