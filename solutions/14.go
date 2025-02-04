package solutions

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
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
