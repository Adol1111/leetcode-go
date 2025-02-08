package solutions

func SetZeroes(matrix [][]int) {
	rowSize, colSize := len(matrix), len(matrix[0])

	colFlag := false

	// 第一轮遍历，仅把第一列和对应列的第一行标记为0, 这样不会检测到被标记的0，而只会捕捉到原始的0
	// 但第一列如果存在原始的0，需要有单独的标记，否则会被误认为是被标记的0
	// 第二轮遍历，根据第一行和第一列的标记，将对应的行列置为0，为了避免被第一行手动标记的0干扰，需要从最后一行开始遍历
	// 最后第0列需要根据colFlag来判断是否置为0

	// 原始            // 第一轮标记
	// 1  2  3  4     // 1  0  3  4
	// 5  0  7  8     // 0  0  7  8
	// 9  10 11 12    // 9  10 11 12
	// 13 14 15 16    // 13 14 15 16

	for i := 0; i < rowSize; i++ {
		if matrix[i][0] == 0 {
			colFlag = true
		}
		for j := 1; j < colSize; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := rowSize - 1; i >= 0; i-- {
		for j := 1; j < colSize; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if colFlag {
			matrix[i][0] = 0
		}
	}
}
