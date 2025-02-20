package solutions

func NumIslands(grid [][]byte) int {
	count := 0

	// 深度优先搜索，如果遇到1，就把它周围的1都置为0
	var dfs func(grid [][]byte, row, col int)
	dfs = func(grid [][]byte, row, col int) {
		// 越界或非1，直接返回
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
			return
		}
		// 设置为0，表示已经遍历过，避免重复遍历
		grid[row][col] = '0'
		dfs(grid, row-1, col)
		dfs(grid, row+1, col)
		dfs(grid, row, col-1)
		dfs(grid, row, col+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				// 广度优先遍历完成后，此岛屿周边所有1都被置为0，再寻找下一个岛屿
				dfs(grid, i, j)
			}
		}
	}
	return count
}
