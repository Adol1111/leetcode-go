package solutions

func Convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	res := make([][]byte, numRows)
	step := -1
	j := 0
	// j为新单词所在的行号，忽略空格，只存储字符，j在0-numRows-1之间循环变化
	for i := 0; i < len(s); i++ {
		res[j] = append(res[j], s[i])
		// 如果到达边界，就反向
		if j == 0 || j == numRows-1 {
			step = -step
		}
		j += step
	}
	ans := []byte{}
	for i := 0; i < numRows; i++ {
		ans = append(ans, res[i]...)
	}
	return string(ans)
}
