package solutions

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

func QuadNodeConstruct(grid [][]int) *QuadNode {
	var dfs func([][]int, int, int) *QuadNode
	dfs = func(rows [][]int, c0, c1 int) *QuadNode {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				// 不是叶子节点
				if v != rows[0][c0] {
					rmid, cmid := len(rows)/2, (c0+c1)/2
					return &QuadNode{
						Val:         true,
						IsLeaf:      false,
						TopLeft:     dfs(rows[:rmid], c0, cmid),
						TopRight:    dfs(rows[:rmid], cmid, c1),
						BottomLeft:  dfs(rows[rmid:], c0, cmid),
						BottomRight: dfs(rows[rmid:], cmid, c1),
					}
				}
			}
		}
		return &QuadNode{
			Val:    rows[0][c0] == 1,
			IsLeaf: true,
		}
	}
	return dfs(grid, 0, len(grid))
}
