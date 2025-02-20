package solutions

// 先用dfs把跟边缘相关的都变成A，然后再遍历一遍把O变成X，把A变成O
func Slove(board [][]byte) {
	var m, n int

	m = len(board)
	n = len(board[0])

	// 深度优先，把跟边缘相关的O都变成A, 再把剩余的O变成X，把A变回O
	var dfs func(grid [][]byte, row, col int)
	dfs = func(grid [][]byte, row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] != 'O' {
			return
		}
		grid[row][col] = 'A'
		dfs(grid, row-1, col)
		dfs(grid, row+1, col)
		dfs(grid, row, col-1)
		dfs(grid, row, col+1)
	}

	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	// 因为第1列和最后1列都已处理，所以范围是[1, n-2]
	for j := 1; j < n-1; j++ {
		dfs(board, 0, j)
		dfs(board, m-1, j)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
