package solutions

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for node := head; node != nil; {
		newNode := &Node{Val: node.Val, Next: node.Next}
		node.Next = newNode
		node = newNode.Next
	}

	for node := head; node != nil; {
		if node.Random != nil {
			// 拷贝节点的 Random 指针 指向当前节点的 Random 指针的下一个节点
			// 因为当前节点的下一个节点就是当前节点的拷贝节点
			node.Next.Random = node.Random.Next
		}
		// 移动到下一个节点非拷贝节点
		node = node.Next.Next
	}

	// 断开拷贝节点和原节点的连接
	copyHead := head.Next
	for node := head; node != nil; node = node.Next {
		copyNode := node.Next
		node.Next = node.Next.Next
		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}
	}

	return copyHead
}
