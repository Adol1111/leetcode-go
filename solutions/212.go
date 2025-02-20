package solutions

func FindWords(board [][]byte, words []string) []string {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}

	m, n := len(board), len(board[0])
	seen := map[string]bool{}

	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		// 如果越界 或 当前节点为# 或 不存在，直接返回
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		ch := board[x][y]
		if ch == '#' {
			return
		}
		node = node.children[ch-'a']
		if node == nil {
			return
		}

		// 如果当前节点正好是单词的结尾，则说明单词存在，记录到seen中
		if node.word != "" {
			seen[node.word] = true
		}

		// 避免深度遍历时重复访问
		board[x][y] = '#'

		// 遍历当前位置的四周，如果周围还未访问，则进行深度遍历
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			dfs(node, nx, ny)
		}
		// 遍历完成后恢复，给下一次遍历使用
		board[x][y] = ch
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(t, i, j)
		}
	}

	ans := make([]string, 0, len(seen))
	for k := range seen {
		ans = append(ans, k)
	}
	return ans
}
