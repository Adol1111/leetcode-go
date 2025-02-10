package solutions

// 这里采用贪心算法
// 需要求∑(a[ri]-a[li])，其中a[ri]-a[li]可以拆分为多个相邻元素的差值，即
// a[ri]-a[li] =（a[ri]-a[ri-1]）+（a[ri-1]-a[ri-2]）+...+（a[li+1]-a[li]）
// 如果要使得这段时间都最大，那么只有当a[ri]-a[ri-1]>0时才加入到结果中，否则取0, 即不持有股票
// 同理，要使∑(a[ri]-a[li])最大，即每个a[ri]-a[li]都取最大值
func MaxProfit2(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		result += max(prices[i]-prices[i-1], 0)
	}
	return result
}
