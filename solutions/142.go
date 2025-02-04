package solutions

func Trap(height []int) int {
	length := len(height)
	lmax := make([]int, length)
	rmax := make([]int, length)

	lmax[0] = height[0]
	for i := 1; i < length; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}

	rmax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}

	sum := 0
	for i := 0; i < length; i++ {
		sum += min(lmax[i], rmax[i]) - height[i]
	}
	return sum
}
