package solutions

func id2rc(id, n int) (row, col int) {
	row, col = (id-1)/n, (id-1)%n // 二维数组中的行号和列号
	if row%2 == 1 {
		// 因为是蛇行，如果行号是奇数，那么列号是从右往左数的，所以需要将列号反转
		col = n - 1 - col
	}
	row = n - 1 - row // 是从下往上数的，所以需要将行号反转
	return
}

func SnakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n+1)
	// step表示到达当前位置需要的步数
	type pair struct{ id, step int }
	q := []pair{{1, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 1; i <= 6; i++ { // 按题目要求每次会走1-6步，所以可以遍历6次，搜索所有的可能性
			next := p.id + i
			if next > n*n { // 超出边界，后续的次数都不会到达终点，所以直接break
				break
			}
			row, col := id2rc(next, n) // 得到下一步的行列
			if board[row][col] > 0 {   // 存在蛇或梯子
				next = board[row][col]
			}
			if next == n*n { // 到达终点
				return p.step + 1
			}
			if !visited[next] {
				visited[next] = true
				// 记录新到达的位置，步数+1，每个步骤都会尝试6次，如果已经遍历过，就不需要再遍历了
				q = append(q, pair{next, p.step + 1})
			}
		}
	}
	return -1
}
