package solutions

func MinPathSum(grid [][]int) int {
	// d[i][j] = min(d[i-1][j], d[i][j-1]) + grid[i][j]
	n := len(grid)

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, len(grid[i]))
	}

	// 定义起点初始值
	d[0][0] = grid[0][0]

	// 先求第一行和第一列的最小路径和
	for i := 1; i < n; i++ {
		d[i][0] = d[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		d[0][j] = d[0][j-1] + grid[0][j]
	}

	// 根据状态转移方程求最小路径和
	for i := 1; i < n; i++ {
		for j := 1; j < len(grid[i]); j++ {
			d[i][j] = min(d[i-1][j], d[i][j-1]) + grid[i][j]
		}
	}
	return d[n-1][len(grid[n-1])-1]
}
