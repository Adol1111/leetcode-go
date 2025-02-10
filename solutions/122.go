package solutions

func MaxProfit2(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		result += max(prices[i]-prices[i-1], 0)
	}
	return result
}
