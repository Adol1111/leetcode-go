package solutions

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	// 取第一个字符串，然后逐个字符比较
	// 如果每个字符串的第i位都相等，则再取一个字符，否则返回
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := 1; j < length; j++ {
			if i >= len(strs[j]) || strs[j][i] != ch {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
