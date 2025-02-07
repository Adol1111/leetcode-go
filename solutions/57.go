package solutions

func Insert(intervals [][]int, newInterval []int) [][]int {
	// 插入后使用56题方法
	// intervals = append(intervals, newInterval)
	// return Merge(intervals)

	// 不使用56题方法
	merged := false
	left := newInterval[0]
	right := newInterval[1]

	var result [][]int = [][]int{}

	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				result = append(result, []int{left, right})
				merged = true
			}
			result = append(result, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			result = append(result, interval)
		} else {
			// 有交集, 先合并区间, 但不插入
			left = min(interval[0], left)
			right = max(interval[1], right)
		}
	}
	if !merged {
		result = append(result, []int{left, right})
	}
	return result
}
