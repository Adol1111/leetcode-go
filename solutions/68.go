package solutions

import "strings"

func FullJustify(words []string, maxWidth int) []string {
	length := len(words)
	ans := []string{}
	for i := 0; i < length; {
		l := len(words[i])
		k := i
		for j := i + 1; j < length; j++ {
			// l 为累计的单词长度, 单词的个数，直接通过j-i获取
			// 累计单词长度 + 当前单词长度 + 单词之间的空格数
			if l+len(words[j])+(j-i) > maxWidth {
				break
			}
			l += len(words[j])
			k = j
		}
		s := words[i]
		if k == i {
			// 只有一个单词，后面补充空格
			s += strings.Repeat(" ", maxWidth-l)
		} else {
			if k == length-1 {
				// 最后一个单词，前面只有一个空格，最后一个单词之后补满空格
				for j := i + 1; j <= k; j++ {
					s += " " + words[j]
				}
				s += strings.Repeat(" ", maxWidth-len(s))
			} else {
				// maxWidth - l 为空格数， k-i 为单词数量，maxSpace为每个单词之间的空格数
				// 按照余数分配空格，每次最多比maxSpace多一个空格, 直到余数分配完
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
		i = k + 1 // 更新i的位置
	}
	return ans
}
