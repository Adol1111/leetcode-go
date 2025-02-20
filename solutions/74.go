package solutions

func SearchMatrix(matrix [][]int, target int) bool {
	up := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	// 先找第一列，有4种可能性
	// 1. 第一列存在目标 => 返回true
	// 2. 中点比目标大 => 目标在up~mid-1之间
	// 3. 中点比目标小 => 目标在mid+1~down之间
	// 4. 上下界重合 => 目标不存在, 此时锁定了目标所在的行，需要在该行进行二分查找
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

	// 上下重合，但是down已经越界，说明目标不存在
	if down == -1 {
		return false
	}

	// down和up重合，取down或up都可以
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
