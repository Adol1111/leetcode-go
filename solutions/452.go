package solutions

import "sort"

func FindMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	// 由于已经按照右边界从小到大排完序, 所以每当一个区间的左边界大于最大的右边界，则说明该区间和已有区间没有重合，因此需要新的箭

	maxRight := points[0][1]
	cnt := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > maxRight {
			maxRight = points[i][1]
			cnt++
		}
	}
	return cnt
}
