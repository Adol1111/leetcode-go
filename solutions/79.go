package solutions

func Exist(board [][]byte, word string) bool {
	var dirs = []struct{ row, col int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var backtrack func(board [][]byte, word string, index int, row, col int) bool
	backtrack = func(board [][]byte, word string, index int, row, col int) bool {
		// 越界或非目标字符，直接返回
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != word[index] {
			return false
		}
		// 最后一个字符，且相等，返回true
		if index == len(word)-1 {
			return true
		}

		ch := board[row][col]
		// 设置为0，表示已经遍历过，避免重复遍历
		board[row][col] = '0'

		for _, dir := range dirs {
			nrow, ncol := row+dir.row, col+dir.col
			if backtrack(board, word, index+1, nrow, ncol) {
				return true
			}
		}
		// 回溯
		board[row][col] = ch

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}

	return false
}
