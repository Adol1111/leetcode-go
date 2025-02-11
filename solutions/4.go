package solutions

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		// 保证数组1的长度小于数组2 这样比较方便后续计算
		nums1, nums2 = nums2, nums1
	}

	m := len(nums1)
	n := len(nums2)

	// 左边元素的总个数
	totalLeft := (m + n + 1) / 2
	left := 0
	right := m

	// a b c d
	// e f g h i

	for left < right {
		i := left + (right-left+1)/2 // i为数组1 left 和 right 中间的位置
		j := totalLeft - i
		// 如果数组1中分割线右侧的元素大于数组2中分割线左侧的元素
		if nums1[i-1] > nums2[j] {
			// 下一轮搜索的区间 [left, i - 1]
			right = i - 1
		} else {
			// 下一轮搜索的区间 [i, right]
			left = i
		}
	}

	i := left
	j := totalLeft - i

	var (
		nums_im1 int
		nums_i   int
		nums_jm1 int
		nums_j   int
	)

	if i == 0 {
		nums_im1 = math.MinInt32
	} else {
		nums_im1 = nums1[i-1]
	}
	if i == m {
		nums_i = math.MaxInt32
	} else {
		nums_i = nums1[i]
	}
	if j == 0 {
		nums_jm1 = math.MinInt32
	} else {
		nums_jm1 = nums2[j-1]
	}
	if j == n {
		nums_j = math.MaxInt32
	} else {
		nums_j = nums2[j]
	}
	if (m+n)%2 == 1 {
		return math.Max(float64(nums_im1), float64(nums_jm1))
	} else {
		return (math.Max(float64(nums_im1), float64(nums_jm1)) + math.Min(float64(nums_i), float64(nums_j))) / 2
	}
}
