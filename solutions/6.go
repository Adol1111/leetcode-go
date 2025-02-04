package solutions

func Convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	res := make([][]byte, numRows)
	step := -1
	j := 0
	for i := 0; i < len(s); i++ {
		res[j] = append(res[j], s[i])
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
