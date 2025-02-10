package solutions

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	for i, j := 0, 0; j < len(t); j++ {
		// 如果相等，则i移动，否则移动j，直到s都匹配完
		if s[i] == t[j] {
			i++
		}
		if i == len(s) {
			return true
		}
	}
	return false
}
