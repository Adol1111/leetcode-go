package solutions

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 如果d[i][j] 不是障碍物，则
	// d[i][j] = d[i-1][j] + d[i][j-1]
	// 否则d[i][j]=0

	// 定义起点初始值
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	n := len(obstacleGrid)
	m := len(obstacleGrid[0])

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, m)
	}

	d[0][0] = 1

	// 先初始化第一行和第一列的路径数量
	for i := 1; i < n; i++ {
		if obstacleGrid[i][0] == 1 {
			d[i][0] = 0
			continue
		}
		d[i][0] = d[i-1][0]
	}
	for j := 1; j < m; j++ {
		if obstacleGrid[0][j] == 1 {
			d[0][j] = 0
			continue
		}
		d[0][j] = d[0][j-1]
	}

	// 根据状态转移方程求最小路径和
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				d[i][j] = 0
				continue
			}
			d[i][j] = d[i-1][j] + d[i][j-1]
		}
	}
	return d[n-1][m-1]
}
