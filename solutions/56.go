package solutions

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int = [][]int{
		intervals[0],
	}

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		for i := len(result) - 1; i >= 0; i-- {
			if interval[0] <= result[i][1] && interval[0] >= result[i][0] {
				result[i][1] = max(interval[1], result[i][1])
				break
			} else if interval[1] <= result[i][1] && interval[1] >= result[i][0] {
				result[i][0] = min(interval[0], result[i][0])
				break
			} else if interval[0] <= result[i][0] && interval[1] >= result[i][1] {
				result[i][0] = interval[0]
				result[i][1] = interval[1]
				break
			} else if interval[0] > result[i][1] || interval[1] < result[i][0] {
				result = append(result, interval)
				break
			}
		}
	}

	return result
}
