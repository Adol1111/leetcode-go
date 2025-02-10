package solutions

func ReverseWords(s string) string {
	length := len(s)
	ans := make([]byte, length)
	idx := 0
	// 从后往前遍历，找到一个单词，然后输出
	for i := length - 1; i >= 0; {
		if s[i] == ' ' {
			i--
			continue
		}
		// 找到一个单词的起始位置(循环结束时，j会指向空格位置或者-1，所以其实位置为j+1)
		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}
		for k := j + 1; k <= i; k++ {
			ans[idx] = s[k]
			idx++
		}
		if idx == length {
			return string(ans)
		}
		// 补充一个空格，并更新i的位置为单词的起始位置-1(即j的位置)
		ans[idx] = ' '
		idx++
		i = j
	}
	return string(ans[:idx-1])
}
