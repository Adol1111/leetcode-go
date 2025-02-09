package solutions

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return CountNodes(root.Left) + CountNodes(root.Right) + 1
}
