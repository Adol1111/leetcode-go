package solutions

func Trap(height []int) int {
	length := len(height)
	lmax := make([]int, length)
	rmax := make([]int, length)

	// 最小值为，当前位置左右两边最大值中的较小的一个（木桶原理）- 当前高度，min(leftMax[i],rightMax[i])−height[i]
	// 先计算左边的最大值，再计算右边的最大值

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
