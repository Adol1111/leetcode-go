package solutions

func MaxProfit3(prices []int) int {
	// 一共有5种情况
	// 1. 未进行过任何操作；
	// 2. 只进行过一次买操作；
	// 3. 进行了一次买操作和一次卖操作，即完成了一笔交易；
	// 4. 在完成了一笔交易的前提下，进行了第二次买操作；
	// 5. 完成了全部两笔交易。

	// 未进行过任何操作，可以忽略，因此由以下方程
	// 1. 第i天第一次买，可以不进行任何操作(同i-1天)，或买入一次股票
	// buy1 = max(buy1', -prices[i])
	// 2. 第i天第一次买，可以不进行任何操作(同i-1天)，或在第1次买入的情况下卖出一次股票
	// sell1 = max(sell1', buy1 + prices[i])
	// 3. 第i天第2次买，可以不进行任何操作(同i-1天)，或在第1次卖的基础上买入股票
	// buy2 = max(buy2', sell1-prices[i])
	// 4. 第i天第2次卖，可以不进行任何操作(同i-1天), 或在第2次买入的情况下卖出股票
	// sell2 = max(sell2', buy2 + prices[i])

	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
