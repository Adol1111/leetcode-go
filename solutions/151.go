package solutions

func ReverseWords(s string) string {
	length := len(s)
	ans := make([]byte, length)
	idx := 0
	for i := length - 1; i >= 0; {
		if s[i] == ' ' {
			i--
			continue
		}
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
		ans[idx] = ' '
		idx++
		i = j
	}
	return string(ans[:idx-1])
}
