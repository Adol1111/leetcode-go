package solutions

func GameOfLife(board [][]int) {
	// -1 表示将要死亡，原本是活的
	// 0 表示原本死亡
	// 1 表示原本活着
	// 2 表示将要复活，原本是死的
	// 统计时之关心原本是活的，所以只需要判断 1 和 -1

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			liveCount := 0
			// 遍历周围一圈，统计活着的数量
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					if k < 0 || k >= len(board) || l < 0 || l >= len(board[0]) || (k == i && l == j) {
						continue
					}
					if board[k][l] == 1 || board[k][l] == -1 {
						liveCount++
					}
				}
			}
			if liveCount == 3 && board[i][j] == 0 {
				board[i][j] = 2
			}
			if liveCount > 3 || liveCount < 2 {
				if board[i][j] == 1 {
					board[i][j] = -1
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == -1 {
				board[i][j] = 0
			}
		}
	}
}
