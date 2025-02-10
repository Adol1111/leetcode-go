package solutions

import "math"

func GetMinimumDifference(root *TreeNode) int {
	var (
		prev     *TreeNode
		minValue = math.MaxInt32
	)
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		if prev != nil {
			minValue = min(root.Val-prev.Val, minValue)
		}
		prev = root
		inOrder(root.Right)
	}
	inOrder(root)
	return minValue
}
