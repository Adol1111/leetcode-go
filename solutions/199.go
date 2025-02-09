package solutions

func RightSideView(root *TreeNode) []int {
	var result []int
	var rightSideView func(root *TreeNode, level int)

	rightSideView = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 如果当前层数等于结果数组的长度，说明当前节点是当前层数的最右侧节点
		// 否则已插入，则不再插入
		if level == len(result) {
			result = append(result, root.Val)
		}
		// 先找右边，再找左边，确保右边的节点一定会先插入结果数组，从而保证结果数组中的节点是最右侧的节点
		rightSideView(root.Right, level+1)
		rightSideView(root.Left, level+1)
	}

	rightSideView(root, 0)
	return result
}
