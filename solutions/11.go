package solutions

func MaxArea(height []int) int {
	m := 0
	// 容积 = 高度 * 宽度
	// 根据木桶原理，高度取决于最短的那块木板，所以从两边开始，每次移动最短的那块木板
	// 最大容积为当前值和 新的容积 比较，取最大值
	for i, j := 0, len(height)-1; i < j; {
		h := min(height[i], height[j])
		m = max(m, h*(j-i))
		// 移动最短的那块木板，因为如果移动长的那块，容积一定会减小
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return m
}
