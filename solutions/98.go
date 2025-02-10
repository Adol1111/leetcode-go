package solutions

import "math"

func IsValidBST(root *TreeNode) bool {
	var isValidBST func(root *TreeNode, minValue int, maxValue int) bool
	isValidBST = func(root *TreeNode, minValue int, maxValue int) bool {
		if root == nil {
			return true
		}

		if root.Val <= minValue || root.Val >= maxValue {
			return false
		}

		return isValidBST(root.Left, minValue, root.Val) && isValidBST(root.Right, root.Val, maxValue)
	}

	return isValidBST(root, math.MinInt64, math.MaxInt64)
}
