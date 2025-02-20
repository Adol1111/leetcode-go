package solutions

func IsInterleave(s1 string, s2 string, s3 string) bool {
	// f(i,j) 表示 s1的前i个字符和s2的前j个字符是否能交错组成s3的前i+j个字符
	// 前i-1个字符和s2的前j个字符交错，则s1的第i个字符和s3的第i+j个字符相等 (第i个字符，在s1中的序号是i-1)
	// 同理，前i个字符和s2的前j-1个字符交错，则s2的第j个字符和s3的第i+j个字符相等 (第j个字符，在s2中的序号是j-1)
	// 因此，f(i,j) = f(i-1,j) && s1[i-1] == s3[i+j-1] || f(i,j-1) && s2[j-1] == s3[i+j-1]

	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[len(s1)][len(s2)]
}
