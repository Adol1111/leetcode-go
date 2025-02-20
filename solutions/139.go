package solutions

func WordBreak(s string, wordDict []string) bool {
	// dp[i] 表示 s[:i] 是否可以被 wordDict 中的单词组成
	// dp[i] = dp[j] && s[j:i] in wordDict

	contains := func(word string, wordDict []string) bool {
		for _, w := range wordDict {
			if w == word {
				return true
			}
		}
		return false
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && contains(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
