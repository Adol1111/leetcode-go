package solutions

func MinDistance(word1 string, word2 string) int {
	// dp[i][j] 表示 word1[:i] 和 word2[:j] 的最小编辑距离
	// 如果 word1[i] == word2[j]，那么不需要编辑，则 dp[i][j] = dp[i-1][j-1]
	// 否则，有3种编辑方法
	// 1. 删除或插入 word1[i]，则 dp[i][j] = dp[i-1][j] + 1
	// 2. 删除或插入 word2[j]，则 dp[i][j] = dp[i][j-1] + 1
	// 3. 替换 word1[i] 或者 word2[j]，则 dp[i][j] = dp[i-1][j-1] + 1
	// dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)

	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	dp[0][0] = 0
	for i := 1; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for j := 1; j < len(word2)+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
