package solutions

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 这里用数组不用map可以更快，但是要求字符的ascii码在0-127之间，如果题目是unicode字符，就要用map
	m := make([]int, 127)
	maxLen := 0
	for i, j := 0, 0; j < len(s); j++ {
		// 如果重复，则把左边界移动到重复字符的下一个位置，如果已经在左边界的右边，则不移动
		// 如果是map, 直接 idx,ok := m[s[j]]; ok 来判断即可
		if idx := m[s[j]]; idx != 0 || s[0] == s[j] && j != 0 {
			// 这个操作可以实现跳步，比如abba，第一个b和第二个b重复，那么i直接跳到第一个b的下一个位置，这样比传统的滑动窗口更快
			i = max(i, idx+1)
		}
		m[s[j]] = j
		maxLen = max(maxLen, j-i+1)
	}
	return maxLen
}
