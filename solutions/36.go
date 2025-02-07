package solutions

func IsValidSudoku(board [][]byte) bool {
	row := [9][9]int{}
	col := [9][9]int{}
	box := [3][3][9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			val := board[i][j] - '1'
			if row[i][val] == 0 && col[j][val] == 0 && box[i/3][j/3][val] == 0 {
				row[i][val] = 1
				col[j][val] = 1
				box[i/3][j/3][val] = 1
			} else {
				return false
			}
		}
	}
	return true
}
