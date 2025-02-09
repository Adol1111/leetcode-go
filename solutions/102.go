package solutions

func LevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	// 广度优先遍历
	// 1. 每层循环时，把下层节点补充到队列中，为了避免遍历到下层节点，需要控制循环次数，次数为循环前的队列长度、
	// 2. 每个子节点循环后，需要把自身从队列中删除
	for len(queue) > 0 {
		level := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
