package solutions

type BSTIterator struct {
	arr   []int
	index int
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.inorder(root)
	return it
}

func (it *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.inorder(node.Left)
	it.arr = append(it.arr, node.Val)
	it.inorder(node.Right)
}

func (it *BSTIterator) Next() int {
	if len(it.arr) == 0 {
		return -1
	}
	res := it.arr[it.index]
	it.index++
	return res
}

func (it *BSTIterator) HasNext() bool {
	return it.index < len(it.arr)
}
