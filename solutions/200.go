package solutions

func NumIslands(grid [][]byte) int {
	count := 0

	var dfs func(grid [][]byte, row, col int)
	dfs = func(grid [][]byte, row, col int) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
			return
		}
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
				dfs(grid, i, j)
			}
		}
	}
	return count
}
