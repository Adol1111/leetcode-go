package solutions

func MaxProfit(prices []int) int {
	maximumProfit, prevMin := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		maximumProfit = max(maximumProfit, prices[i]-prevMin)
		prevMin = min(prevMin, prices[i])
	}
	return maximumProfit
}
