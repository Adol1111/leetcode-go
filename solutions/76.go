package solutions

func MinWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}

	// 这里用数组会比 map 更快，但是要求字符的 ascii 码在 0-127 之间，如果题目是 unicode 字符，就要用 map
	sFreq, tFreq := [256]int{}, [256]int{}
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	position := make([]int, 0)

	minLen := len(s) + 1
	minStart := 0
	needCnt := 0
	i := 0

	for left, right := 0, 0; right < len(s); right++ {
		current := s[right]
		if tFreq[current-'a'] == 0 {
			continue
		}

		// 记录 t 中字符在 s 中的位置，可以快速跳过不需要的字符
		position = append(position, right)

		// 只要是t中的字母就++, 不管是否超过t中的次数，对于多处理的部分，会在下面的循环中处理
		sFreq[current-'a']++
		if sFreq[current-'a'] <= tFreq[current-'a'] {
			needCnt++
		}

		for needCnt == len(t) {
			left = position[i]
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				needCnt--
			}
			sFreq[s[left]-'a']--
			i++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
