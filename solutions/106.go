package solutions

func BuildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == root.Val {
			break
		}
	}
	// 排除后序遍历中的根节点 和 中序遍历中的根节点
	root.Left = BuildTree2(inorder[:index], postorder[:index])
	root.Right = BuildTree2(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
