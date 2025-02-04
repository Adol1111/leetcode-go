package solutions

import "strings"

func FullJustify(words []string, maxWidth int) []string {
	length := len(words)
	ans := []string{}
	for i := 0; i < length; {
		l := len(words[i])
		k := i
		for j := i + 1; j < length; j++ {
			if l+len(words[j])+(j-i) > maxWidth {
				break
			}
			l += len(words[j])
			k = j
		}
		s := words[i]
		if k == i {
			s += strings.Repeat(" ", maxWidth-l)
		} else {
			if k == length-1 {
				for j := i + 1; j <= k; j++ {
					s += " " + words[j]
				}
				s += strings.Repeat(" ", maxWidth-len(s))
			} else {
				var maxSpace = (maxWidth - l) / (k - i)
				space := 0
				for j := i + 1; j <= k; j++ {
					if (maxWidth-l)%(k-i) != 0 && space < (maxWidth-l)%(k-i) {
						s += strings.Repeat(" ", maxSpace+1)
						space++
					} else {
						s += strings.Repeat(" ", maxSpace)
					}
					s += words[j]
				}
			}
		}
		ans = append(ans, s)
		i = k + 1
	}
	return ans
}
