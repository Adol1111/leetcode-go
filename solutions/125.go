package solutions

import "strings"

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		// 如果非数字和字母，则跳过
		if s[i] < '0' || s[i] > '9' && (s[i] < 'a' || s[i] > 'z') {
			i++
			continue
		}
		// 如果非数字和字母，则跳过
		if s[j] < '0' || s[j] > '9' && (s[j] < 'a' || s[j] > 'z') {
			j--
			continue
		}
		// 如果相等，则两边同时移动，否则返回false
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
