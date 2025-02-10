package solutions

func Candy(ratings []int) int {
	length := len(ratings)
	ans := make([]int, length)
	ans[0] = 1
	// 1. 先正着发，最少是1，如果比前一个大，就比前一个多1，否则发最少量1
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = 1
		}
	}
	// 倒过来再修正一遍，如果ratings[i+1] < ratings[i], 则修正ans[i] = ans[i+1] + 1
	// 如果已经比ans[i+1]大了，就不用修正
	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if ans[i] <= ans[i+1] {
				ans[i] = ans[i+1] + 1
			}
		}
	}

	sum := 0
	for _, v := range ans {
		sum += v
	}
	return sum
}
