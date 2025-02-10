package solutions

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rowSize, colSize := len(matrix), len(matrix[0])
	total := rowSize * colSize
	result := make([]int, total)

	top, bottom, left, right := 0, rowSize-1, 0, colSize-1

	for n := 0; n < total; {
		// 从左到右
		for i := left; i <= right && n < total; i++ {
			result[n] = matrix[top][i]
			n++
		}
		top++

		// 从上到下
		for i := top; i <= bottom && n < total; i++ {
			result[n] = matrix[i][right]
			n++
		}
		right--

		// 从右到左
		for i := right; i >= left && n < total; i-- {
			result[n] = matrix[bottom][i]
			n++
		}
		bottom--

		// 从下到上
		for i := bottom; i >= top && n < total; i-- {
			result[n] = matrix[i][left]
			n++
		}
		left++
	}
	return result
}
