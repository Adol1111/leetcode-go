package solutions

func Candy(ratings []int) int {
	length := len(ratings)
	ans := make([]int, length)
	ans[0] = 1
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = 1
		}
	}
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
