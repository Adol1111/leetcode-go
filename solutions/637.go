package solutions

func AverageOfLevels(root *TreeNode) []float64 {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([]float64, 0)

	// 广度优先遍历, 参考102
	for len(queue) > 0 {
		sum := 0
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(length))
	}
	return result
}
