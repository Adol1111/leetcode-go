package solutions

func StrStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	// 构造 pi 数组，pi[i] 表示 needle[:i+1] 最长相等前后缀的长度
	// 比如 aba , 模式值为 001, 最后一个a的最长相等前缀为a，长度为1
	// 比如 abab, 模式值为 0012, 最后一个ab和第一个ab相等，长度为2
	// 比如 aabaaa, 模式值为 010122, b后面的aaa分别是122，表示
	//	- 第一个a和开头的a有相同前缀a，长度为2
	//	- 第一个aa和开头的aa有相同前缀aa，长度为2
	//  - 第二个aa也和开头的a有相同前缀a，长度为2
	// 因此 aabaaab 的 模式值为 0101223
	// pi[0]始终为0
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			// 如果不相等，则回退到上一个最长相等前缀的下一个字符
			// 长度和位置是等价的，表示前N个可以不比较，直接从j开始比较
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}
