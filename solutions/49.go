package solutions

func GroupAnagrams(strs []string) [][]string {
	m := make(map[[26]int]int)
	var res [][]string

	for _, str := range strs {
		var arr [26]int
		for _, ch := range str {
			arr[ch-'a']++
		}
		if idx, ok := m[arr]; ok {
			res[idx] = append(res[idx], str)
		} else {
			m[arr] = len(res)
			res = append(res, []string{str})
		}
	}
	return res
}
