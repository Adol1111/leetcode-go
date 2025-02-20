package solutions

func TotalNQueens(n int) int {
	count := 0
	// 初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	isValid := func(board [][]byte, row, col int) bool {
		// 皇后是要互相不冲突，也就是说：两个皇后不能在同一行、同一列或者同一斜线上。
		// 这里不检查同一行时因为每调用一次 backtrack 只放一个皇后，而每行只放一个皇后。
		// 当这行结束递归后，会对该行进行回溯，把皇后放在下一列，所以不需要检查同一行。

		// 检查列是否有皇后互相冲突
		for i := 0; i < n; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		// 检查右上方是否有皇后互相冲突
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		// 检查左上方是否有皇后互相冲突
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	var backtrack func(board [][]byte, row int)
	backtrack = func(board [][]byte, row int) {
		// 结束条件
		if row == n {
			count++
			return
		}
		for col := 0; col < n; col++ {
			// 排除不合法选择
			if !isValid(board, row, col) {
				continue
			}
			// 做选择
			board[row][col] = 'Q'
			// 进入下一行决策
			backtrack(board, row+1)
			// 撤销选择
			board[row][col] = '.'
		}
	}
	backtrack(board, 0)
	return count
}
