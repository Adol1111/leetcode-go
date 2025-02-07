package solutions

func FindSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordsLen := len(words)
	total := wordLen * wordsLen

	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	checkMap := func(map1 map[string]int, map2 map[string]int) bool {
		for k, v := range map1 {
			if map2[k] != v {
				return false
			}
		}
		return true
	}

	// 方案1: 每次发现不匹配就移动left重新开始，这样会有很多重复的计算
	// 方案2: 发现不匹配不终止一直查询下去，直到最后，这样只需要查询 word 的长度次
	//		- 如果有不匹配单次，则移动left到right+wordLen，同时清空整个统计从头开始
	//		- 如果匹配单词次数等于words中的次数，则检查是否全部匹配，如果全部匹配则记录
	//		- 如果right-left == total，说明数量多余需要的数量，那么移动left，同时减掉left的单词次数
	ans := make([]int, 0)
	for i := 0; i < wordLen; i++ {
		right, left := i, i

		wordCount := make(map[string]int)
		for left+total <= len(s) {
			word := s[right : right+wordLen]
			if _, ok := wordsMap[word]; !ok {
				// 如果不是words中的单词，那么就重新开始, 但不break继续下一个单词，一直查到最后
				wordCount = make(map[string]int)
				left = right + wordLen
				right = left
				continue
			}
			right += wordLen
			wordCount[word]++
			if wordCount[word] == wordsMap[word] {
				if checkMap(wordsMap, wordCount) {
					ans = append(ans, left)
				}
			}
			// 如果超过了words的长度，那么就要移动left，同时要把left的单词从wordCount中减掉
			if right-left == total {
				wordCount[s[left:left+wordLen]]--
				left += wordLen
			}
		}
	}

	return ans
}
