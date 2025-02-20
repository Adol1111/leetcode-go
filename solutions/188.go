package solutions

func MaxProfit4(k int, prices []int) int {
	// 参考123题
	buy := make([]int, k)
	sell := make([]int, k)

	for i := 0; i < k; i++ {
		buy[i], sell[i] = -prices[0], 0
	}

	for i := 1; i < len(prices); i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])

		for j := 1; j < k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k]
}
