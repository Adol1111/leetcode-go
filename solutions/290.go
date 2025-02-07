package solutions

func WordPattern(pattern string, s string) bool {
	// 不使用split, 仅靠双指针
	m := map[byte]string{}
	n := map[string]byte{}
	j := 0
	for i := range pattern {
		start := j
		for ; j < len(s); j++ {
			if s[j] == ' ' {
				break
			}
		}
		if j > len(s) {
			return false
		}
		word := s[start:j]
		j++
		if m[pattern[i]] == "" {
			m[pattern[i]] = word
		} else if m[pattern[i]] != word {
			return false
		}
		if n[word] == 0 {
			n[word] = pattern[i]
		} else if n[word] != pattern[i] {
			return false
		}
	}
	return true
}
