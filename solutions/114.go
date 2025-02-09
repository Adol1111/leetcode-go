package solutions

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	Flatten(root.Left)
	Flatten(root.Right)
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	// 1. 左节点和右节点已经递归展开成右节点
	// 2. 左节点目前是现在的右节点
	// 3. 把右节点挂载在左节点的最后一个节点上
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
