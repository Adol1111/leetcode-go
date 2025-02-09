package solutions

// 先序遍历，第一个节点是根节点，然后在中序遍历中找到根节点的位置，左边是左子树，右边是右子树，然后递归构建树
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	// 排除先序遍历中的根节点 和 中序遍历中的根节点
	root.Left = BuildTree(preorder[1:index+1], inorder[:index])
	root.Right = BuildTree(preorder[index+1:], inorder[index+1:])
	return root
}
