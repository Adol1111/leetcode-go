package solutions

import "strings"

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if s[i] < '0' || s[i] > '9' && (s[i] < 'a' || s[i] > 'z') {
			i++
			continue
		}
		if s[j] < '0' || s[j] > '9' && (s[j] < 'a' || s[j] > 'z') {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
