package solutions

// 参考433，两题的思路完全相同
func LadderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}

	adjacency := make([][]int, len(wordList))
	endIndex := -1

	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			endIndex = i
		}

		for j := i + 1; j < len(wordList); j++ {
			if isAdjacent(wordList[i], wordList[j]) {
				adjacency[i] = append(adjacency[i], j)
				adjacency[j] = append(adjacency[j], i)
			}
		}
	}

	if endIndex == -1 {
		return -1
	}

	visited := make([]bool, len(wordList))
	queue := []int{}
	// 查询beginWord的所有邻接节点，作为初始节点
	for i, s := range wordList {
		if isAdjacent(beginWord, s) {
			queue = append(queue, i)
			visited[i] = true
		}
	}

	for step := 1; len(queue) > 0; step++ {
		tmp := queue
		queue = nil

		for _, cur := range tmp {
			if cur == endIndex {
				return step
			}

			for _, next := range adjacency[cur] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return -1
}
