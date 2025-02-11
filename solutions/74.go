package solutions

func SearchMatrix(matrix [][]int, target int) bool {
	up := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for up <= down {
		mid := (up + down) / 2
		if matrix[mid][0] == target {
			return true
		} else if target < matrix[mid][0] {
			down = mid - 1
		} else {
			up = mid + 1
		}
	}

	if down == -1 {
		return false
	}

	for left <= right {
		mid := (left + right) / 2
		if matrix[down][mid] == target {
			return true
		} else if target < matrix[down][mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
