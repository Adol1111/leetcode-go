package solutions

func ZigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	direction := 1
	for len(queue) > 0 {
		level := make([]int, 0)
		next := make([]*TreeNode, 0)

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			// 倒序插入，这样可以不需要再最后做翻转
			// 方向不同时，左右子节点的插入顺序不同，正向是从左到右，反向是从右到左
			if direction == 1 {
				if node.Left != nil {
					next = append([]*TreeNode{node.Left}, next...)
				}
				if node.Right != nil {
					next = append([]*TreeNode{node.Right}, next...)
				}
			} else {
				if node.Right != nil {
					next = append([]*TreeNode{node.Right}, next...)
				}
				if node.Left != nil {
					next = append([]*TreeNode{node.Left}, next...)
				}
			}
		}
		result = append(result, level)
		queue = next
		direction *= -1
	}
	return result
}
