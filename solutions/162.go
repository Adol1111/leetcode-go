package solutions

import "math"

func FindPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		// 越界返回负无穷
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		// 中间的数比两边的数都大, 则找到了
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		// 中间的数比右边的数小, 则右边一定有峰值，否则左边一定有峰值
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
