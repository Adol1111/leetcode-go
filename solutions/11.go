package solutions

func MaxArea(height []int) int {
	m := 0
	for i, j := 0, len(height)-1; i < j; {
		h := min(height[i], height[j])
		m = max(m, h*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return m
}
