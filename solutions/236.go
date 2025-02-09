package solutions

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	// 如果只有右侧则返回右侧，如果只有左侧则返回左侧，如果左右都有则公共祖先就是自己，返回root
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
