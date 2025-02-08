package solutions

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	node := head
	for i := 1; i < left; i++ {
		prev = node
		node = node.Next
	}

	leftNode := prev
	prev = node
	newHead := node
	for i := left + 1; i <= right; i++ {
		node := prev.Next
		prev.Next = node.Next
		node.Next = newHead
		newHead = node
	}
	if leftNode != nil {
		leftNode.Next = newHead
	} else {
		head = newHead
	}

	return head
}

// a->b->c->d->e
// 2, 4
// 2 = b
// 4 = d
// leftNode = a
// 1. prev = a, node = b
