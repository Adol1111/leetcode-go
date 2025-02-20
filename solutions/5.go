package solutions

func LongestPalindrome(s string) string {
	// dp[i][j] 表示 s[i:j] 是否是回文串
	// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	// (j-1)-(i+1) + 1 >= 2 => j - i >= 3 , 当 j - i < 3 时，不需要检查子串是否回文

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	var maxLen int
	var maxStr string

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if j-i < 3 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					maxStr = s[i : j+1]
				}
			}
		}
	}

	if maxLen == 0 {
		return string(s[0])
	}
	return maxStr
}

func LongestPalindrome2(s string) string {
	// 中心扩散法
	if s == "" {
		return ""
	}

	// 判断回文中心扩散到最远的范围
	expandAroundCenter := func(s string, left, right int) (int, int) {
		for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		}
		// 返回扩散的最远范围，因为for循环关系，所以要向内缩一格
		return left + 1, right - 1
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 奇数情况，以当前字符为中心，向两边扩散
		left1, right1 := expandAroundCenter(s, i, i)
		// 偶数情况，以当前字符和下一个字符为中心，向两边扩散
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}
