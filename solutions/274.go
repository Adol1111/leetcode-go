package solutions

func HIndex(citations []int) int {
	n := len(citations)
	count := make([]int, n+1) // count表示引用次数为i的文章数量
	for _, c := range citations {
		if c >= n {
			count[n]++
		} else {
			count[c]++
		}
	}

	// 题目要求是找到最大的h，使得至少h篇文章引用次数大于等于h
	// 因此从后往前遍历count数组，找到最大的h
	// 如果count[i]<i，不符合要求，把total加上count[i]
	total := 0
	for i := n; i >= 0; i-- {
		total += count[i]
		if total >= i {
			return i
		}
	}

	return 0
}
