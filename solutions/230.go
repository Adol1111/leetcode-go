package solutions

func KthSmallest(root *TreeNode, k int) int {
	var (
		stack []*TreeNode
	)

	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		stack = append(stack, root)
		inOrder(root.Right)
	}
	inOrder(root)

	return stack[k-1].Val
}
