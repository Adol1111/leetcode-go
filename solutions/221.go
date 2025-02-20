package solutions

func MaximalSquare(matrix [][]byte) int {
	// dp[i][j] 以i,j为右下角，且i,j为1的正方形的最大边长
	// 1. 如果i,j和(i-1,j-1)构成正方形，则 他的左、上、左上都是1，此时可以根据3者的dp取最小值确定
	// dp[i][j]=min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
	// 2. 不构成则
	// dp[i][j]=0

	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	maxValue := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				}
				maxValue = max(maxValue, dp[i][j])
			}
		}
	}
	return maxValue * maxValue
}
