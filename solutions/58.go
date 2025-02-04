package solutions

func LengthOfLastWord(s string) int {
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if l == 0 {
				continue
			}
			return l
		} else {
			l++
		}
	}
	return l
}
