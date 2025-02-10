package solutions

func HIndex(citations []int) int {
	n := len(citations)
	count := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			count[n]++
		} else {
			count[c]++
		}
	}

	total := 0
	for i := n; i >= 0; i-- {
		total += count[i]
		if total >= i {
			return i
		}
	}

	return 0
}
