package solutions

import "math"

func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue := math.MinInt32
	// maxGain 获取该节点的最大贡献 1. 当前节点 2. 当前节点 + 右节点的最大贡献 3. 当前节点 + 左节点的最大贡献
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 获取左节点的最大贡献值，如果小于0，则舍弃
		maxLeft := max(0, maxGain(root.Left))
		// 获取右节点的最大贡献值，如果小于0，则舍弃
		maxRight := max(0, maxGain(root.Right))
		// 当前节点的最大和为 当前节点的值 + 左右节点的最大贡献值，因为已经和0比较，如果小于0，已经舍弃，直接相加即可
		sum := root.Val + maxLeft + maxRight
		// 如果当前节点的最大贡献值大于全局最大和，则更新
		maxValue = max(maxValue, sum)
		// 返回最大贡献值
		return root.Val + max(maxLeft, maxRight)
	}
	maxGain(root)
	return maxValue
}
